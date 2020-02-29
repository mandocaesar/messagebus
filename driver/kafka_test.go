package driver

import (
	"mandocaesar/messagebus/common"
	"testing"

	"gotest.tools/assert"
)

func TestKafkaDriverInstantiate(t *testing.T) {
	config := &common.KafkaConfig{}
	config.Instantiate()

	broker := config.Get("brokers")
	t.Log(broker)

	assert.Assert(t, broker != nil)
}
