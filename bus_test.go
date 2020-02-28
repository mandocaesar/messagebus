package main

import (
	driver "mandocaesar/messagebus/driver/rabbitMQ"
	"testing"

	"github.com/sirupsen/logrus"
	"gotest.tools/assert"
)

func TestNewMessageBus(t *testing.T) {
	driver := &driver.DriverRMQ{}
	a, err := NewMessageBus(driver, "schemas")

	assert.Assert(t, err == nil)
	assert.Assert(t, a != nil)
}

func TestSchemasFolderShouldNotEmpty(t *testing.T) {
	driver := &driver.DriverRMQ{}
	a, err := NewMessageBus(driver, "schemas")

	assert.Assert(t, err == nil)
	assert.Assert(t, len(a.Schemas) > 0)
}

func TestPublishFromBus(t *testing.T) {
	driver := &driver.DriverRMQ{}
	a, err := NewMessageBus(driver, "schemas")

	result, err := a.Publish("a")
	assert.Assert(t, err == nil)
	assert.Assert(t, result)
}

func register(data interface{}) (interface{}, error) {
	logrus.Info(data)
	return data, nil
}
func TestGetHandlers(t *testing.T) {
	driver := &driver.DriverRMQ{}
	a, _ := NewMessageBus(driver, "schemas")

	a.RegisterHandler("event.user.created", register)
	result, err := a.Handle("event.user.created", driver)
	assert.Assert(t, err == nil)
	assert.Assert(t, result)
}
