package v1

import (
	elastic "gopkg.in/olivere/elastic.v2"

	version "github.com/hashicorp/go-version"
	"transporter/adaptor/elasticsearch/clients"
	"transporter/client"
	"transporter/log"
	"transporter/message"
	"transporter/message/ops"
)

var (
	_ client.Writer = &Writer{}
)

// Writer implements client.Writer and client.Session for sending requests to an elasticsearch
// cluster in individual requests.
type Writer struct {
	index    string
	esClient *elastic.Client
	logger   log.Logger
}

func init() {
	constraint, _ := version.NewConstraint(">= 1.4, < 2.0")
	clients.Add("v1", constraint, func(opts *clients.ClientOptions) (client.Writer, error) {
		esOptions := []elastic.ClientOptionFunc{
			elastic.SetURL(opts.URLs...),
			elastic.SetSniff(false),
			elastic.SetHttpClient(opts.HTTPClient),
			elastic.SetMaxRetries(2),
		}
		if opts.UserInfo != nil {
			if pwd, ok := opts.UserInfo.Password(); ok {
				esOptions = append(esOptions, elastic.SetBasicAuth(opts.UserInfo.Username(), pwd))
			}
		}
		esClient, err := elastic.NewClient(esOptions...)
		if err != nil {
			return nil, err
		}
		w := &Writer{
			index:    opts.Index,
			esClient: esClient,
			logger:   log.With("writer", "elasticsearch").With("version", 1),
		}
		return w, nil
	})
}

func (w *Writer) Write(msg message.Msg) func(client.Session) (message.Msg, error) {
	return func(s client.Session) (message.Msg, error) {
		indexType := msg.Namespace()
		var id string
		if _, ok := msg.Data()["_id"]; ok {
			id = msg.ID()
		}

		var err error
		switch msg.OP() {
		case ops.Delete:
			_, err = w.esClient.Delete().Index(w.index).Type(indexType).Id(id).Do()
		case ops.Insert:
			_, err = w.esClient.Index().Index(w.index).Type(indexType).Id(id).BodyJson(msg.Data()).Do()
		case ops.Update:
			_, err = w.esClient.Index().Index(w.index).Type(indexType).BodyJson(msg.Data()).Id(id).Do()
		}
		if msg.Confirms() != nil && err == nil {
			msg.Confirms() <- struct{}{}
		}
		return msg, err
	}
}
