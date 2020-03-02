package serializer

import (
	"testing"

	avrolib "gopkg.in/avro.v0"
	"gotest.tools/assert"
)

func TestAvroGetSchemas(t *testing.T) {
	avro := &AvroSerializer{}
	schemas := avro.GetAllSchema("../schemas/").(map[string]avrolib.Schema)
	assert.Assert(t, len(schemas) > 0)
}

func TestAvroGetSchema(t *testing.T) {
	avro := NewAvroSerializer()
	avro.GetAllSchema("../schemas/")
	result := avro.GetSchema("kata.MessageHeader")
	assert.Assert(t, result != nil)
}
