package main

import (
	"github.com/Jeffail/benthos/v3/lib/service"
	_ "github.com/usedatabrew/pg_stream_benthos/pg_stream"
	_ "github.com/usedatabrew/pg_stream_benthos/pg_stream_schemaless"
)

func main() {
	service.Run()
}
