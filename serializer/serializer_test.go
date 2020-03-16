package serializer

import (
	"mandocaesar/messagebus/message"
	"testing"

	"github.com/satori/uuid"
	"gotest.tools/assert"
)

func MockFuncSerializeHeader(serializer Serializer) ([]byte, error) {
	header := &message.MessageHeader{
		MessageId:     uuid.NewV4().String(),
		CorrelationId: uuid.NewV4().String(),
		MessageFlags:  1,
		MessageType:   1,
		ReturnAddress: "amqp.libgen.com",
	}

	return serializer.Encode(header, "kata.MessageHeader")
}

func TestSerializerAsAvro(t *testing.T) {
	avro := NewAvroSerializer()
	avro.GetAllSchema("../example/schemas/")

	result, err := MockFuncSerializeHeader(avro)
	headerDecode, err := avro.GetHeader(result)

	assert.Assert(t, err == nil)
	assert.Assert(t, headerDecode.MessageId != "")
}
