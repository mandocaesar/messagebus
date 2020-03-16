package bus

import (
	driver "mandocaesar/messagebus/driver"
	"testing"

	"github.com/sirupsen/logrus"
	"gotest.tools/assert"
)

var path = "../example/schemas"

func TestNewMessageBus(t *testing.T) {
	driver := &driver.Kafka{}
	a, err := NewMessageBus(driver, path)

	assert.Assert(t, err == nil)
	assert.Assert(t, a != nil)
}

func TestSchemasFolderShouldNotEmpty(t *testing.T) {
	driver := &driver.Kafka{}
	a, err := NewMessageBus(driver, path)

	assert.Assert(t, err == nil)
	assert.Assert(t, len(a.Schemas) > 0)
}

func TestPublishFromBus(t *testing.T) {
	driver := &driver.Kafka{}
	a, err := NewMessageBus(driver, path)

	err = a.Publish("a")
	assert.Assert(t, err == nil)
}

func register(data interface{}) (interface{}, error) {
	logrus.Info(data)
	return data, nil
}
func TestGetHandlers(t *testing.T) {
	driver := &driver.Kafka{}
	a, _ := NewMessageBus(driver, path)
	var handleThis = "hello world!"

	a.RegisterHandler(1, register)

	result, err := a.Handle(1, handleThis)

	assert.Assert(t, err == nil)
	assert.Assert(t, result == handleThis)
}
