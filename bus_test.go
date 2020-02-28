package main

import (
	driver "mandocaesar/messagebus/driver/rabbitMQ"
	"testing"

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
