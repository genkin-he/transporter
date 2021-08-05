package all

import (
	// ensures init functions get called
	_ "transporter/adaptor/elasticsearch/clients/v1"
	_ "transporter/adaptor/elasticsearch/clients/v2"
	_ "transporter/adaptor/elasticsearch/clients/v5"
	_ "transporter/adaptor/elasticsearch/clients/v7"
)
