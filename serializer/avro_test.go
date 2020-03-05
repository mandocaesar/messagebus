package serializer

import (
	"mandocaesar/messagebus/message"
	"testing"

	"github.com/satori/uuid"
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
	schema := result.(avrolib.Schema)

	assert.Assert(t, err == nil)
	assert.Assert(t, schema != nil)
}

func TestAvroDecode(t *testing.T) {
	avro := NewAvroSerializer()
	avro.GetAllSchema("../schemas/")
	header := &message.MessageHeader{
		MessageId:     uuid.NewV4().String(),
		CorrelationId: uuid.NewV4().String(),
		MessageFlags:  1,
		MessageType:   1,
		ReturnAddress: "amqp.libgen.com",
	}

	result, err := avro.Encode(header, "kata.MessageHeader")

	assert.Assert(t, err == nil)
	assert.Assert(t, result != nil)
}

func TestAvroGetHeader(t *testing.T) {
	avro := NewAvroSerializer()
	avro.GetAllSchema("../schemas/")
	header := &message.MessageHeader{
		MessageId:     uuid.NewV4().String(),
		CorrelationId: uuid.NewV4().String(),
		MessageFlags:  1,
		MessageType:   1,
		ReturnAddress: "amqp.libgen.com",
	}

	result, _ := avro.Encode(header, "kata.MessageHeader")

	headerDecode, err := avro.GetHeader(result)

	assert.Assert(t, err == nil)
	assert.Assert(t, headerDecode.MessageId != "")
}
