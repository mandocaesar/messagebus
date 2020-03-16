package main

import (
	"mandocaesar/messagebus/bus"
	"mandocaesar/messagebus/common"
	"mandocaesar/messagebus/driver"
	"testing"

	"gotest.tools/assert"
)

var (
	Bus           bus.MessageBus
	KafkaInstance driver.Driver
	KafkaConfig   common.Config
)

func TestConnectToKafka(t *testing.T) {

	KafkaConfig := new(common.KafkaConfig)

	KafkaConfig.Instantiate("")
	KafkaInstance, err := driver.NewKafka(KafkaConfig)

	assert.Assert(t, err != nil)

	result := KafkaInstance.Connect()

	assert.Assert(t, result != nil)
}
