package all

import (
	// Initialize all adapters by importing this package
	_ "transporter/adaptor/elasticsearch"
	_ "transporter/adaptor/file"
	_ "transporter/adaptor/mongodb"
	_ "transporter/adaptor/postgres"
	_ "transporter/adaptor/rabbitmq"
	_ "transporter/adaptor/rethinkdb"
)
