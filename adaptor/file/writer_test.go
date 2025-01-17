package file

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"transporter/adaptor"
	"transporter/message"
	"transporter/message/ops"
)

func TestWrite(t *testing.T) {
	tmpD, err := ioutil.TempDir("", "write_test")
	if err != nil {
		t.Fatalf("unable to create tmp dir, %s", err)
	}
	defer os.RemoveAll(tmpD)
	f, err := os.Create(filepath.Join(tmpD, "data.json"))
	if err != nil {
		t.Fatalf("unable to create file, %s", err)
	}
	defer f.Close()
	tmpSession := &Session{file: f}
	confirms, cleanup := adaptor.MockConfirmWrites()
	defer adaptor.VerifyWriteConfirmed(cleanup, t)
	w := newWriter()
	for i := 0; i < 2; i++ {
		msg := message.From(ops.Insert, "test", map[string]interface{}{"_id": "546656989330a846dc7ce327", "test": "hello world"})
		if i == 0 {
			msg = message.WithConfirms(confirms, msg)
		}
		if _, err := w.Write(msg)(tmpSession); err != nil {
			t.Errorf("unexpected Write error, %s\n", err)
		}
	}
	golden := filepath.Join("testdata", "write_test.golden")
	expected, _ := ioutil.ReadFile(golden)
	actual, _ := ioutil.ReadFile(filepath.Join(tmpD, "data.json"))

	if !bytes.Equal(actual, expected) {
		t.Errorf("mismatched data in file, expected %s, got %s", string(expected), string(actual))
	}
}
