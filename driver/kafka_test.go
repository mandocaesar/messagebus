package driver

import (
	"bytes"
	"mandocaesar/messagebus/common"
	"mandocaesar/messagebus/message"
	"mandocaesar/messagebus/serializer"
	"testing"

	_ "github.com/joho/godotenv/autoload"
	"gotest.tools/assert"
)

type KafkaTest struct {
	Config    *common.KafkaConfig
	Driver    Driver
	Functions map[int32]func(data interface{}) (interface{}, error)
}

func NewKafkaTest() *KafkaTest {
	config := &common.KafkaConfig{}
	config.Instantiate("../.env.example")

	driver, err := NewKafka(config)
	functions := make(map[int32]func(data interface{}) (interface{}, error))
	if err != nil {
		return nil
	}
	return &KafkaTest{Config: config, Driver: driver, Functions: functions}
}

func TestKafkaDriverInstantiate(t *testing.T) {
	instance := NewKafkaTest()
	broker := instance.Config.Get("brokers")

	assert.Assert(t, broker != "")
	assert.Assert(t, len(broker) != 0)
}

func TestKafkaPublish(t *testing.T) {
	instance := NewKafkaTest()
	header := &message.MessageHeader{
		CorrelationId: "111",
		MessageFlags:  1,
		MessageId:     "lalala",
		MessageType:   1,
		ReturnAddress: "aaaa"}

	avroSerializer := serializer.NewAvroSerializer()
	avroSerializer.GetAllSchema("../example/schemas/")
	serializeMessage, err := avroSerializer.Encode(header, "kata.MessageHeader")

	assert.Assert(t, err == nil)

	message := &message.ProducerMessage{
		Topic:   "test",
		Key:     []byte("test"),
		Message: serializeMessage}

	err = instance.Driver.Publish(message)
	assert.Assert(t, err == nil)
}

func Handle(model interface{}) (interface{}, error) {
	return message.DeserializeMessageHeader(bytes.NewReader(model.([]byte)))
}

func TestKafkaConsume(t *testing.T) {
	instance := NewKafkaTest()
	avroSerializer := serializer.NewAvroSerializer()
	avroSerializer.GetAllSchema("../example/schemas/")
	subscription := &message.SubscribeMessage{Group: "", Topic: "test"}

	instance.Functions[0] = Handle
	instance.Functions[1] = Handle

	result, err := instance.Driver.Subscribe(subscription, avroSerializer, instance.Functions)
	assert.Assert(t, result != nil)
	assert.Assert(t, err == nil)
}
