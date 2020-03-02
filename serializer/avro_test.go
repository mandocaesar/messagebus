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

func TestAvroGetSchemaZeroLength(t *testing.T) {
	avro := NewAvroSerializer()
	schemas := avro.GetAllSchema("").(map[string]avrolib.Schema)
	assert.Assert(t, len(schemas) != 0)
}

func TestAvroGetSchema(t *testing.T) {
	avro := NewAvroSerializer()
	avro.GetAllSchema("../schemas/")
	result := avro.GetSchema("kata.MessageHeader")
	assert.Assert(t, result != nil)
}

func TestAvroGetNilSchema(t *testing.T) {
	avro := NewAvroSerializer()
	avro.GetAllSchema("../schemas/")
	result := avro.GetSchema("")
	assert.Assert(t, result == nil)
}

func TestAvroGetNilSchemas(t *testing.T) {
	avro := NewAvroSerializer()
	avro.GetAllSchema("")
	result := avro.GetSchema("kata.MessageHeader")
	assert.Assert(t, result == nil)
}

func TestAvroParse(t *testing.T) {
	testSchema := `{ "type": "record",
			"namespace": "example.avro",
			"name": "System",
			"fields": [
				{"name": "user", "type": "string"}
			]
		}`
	avro := NewAvroSerializer()
	result, err := avro.ParseSchema(testSchema)
	assert.Assert(t, err == nil)
	assert.Assert(t, result != nil)
}
