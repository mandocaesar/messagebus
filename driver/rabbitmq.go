package driver

import common "mandocaesar/messagebus/common"

type RMQ struct {
	config common.Config
}

//SetConfig set config configuration
func (d *RMQ) SetConfig(key string, value interface{}) common.Config {
	return nil
}

//Connect connect to rmq message broker
func (d *RMQ) Connect() error { return nil }

//SendReply reply to request reply pattern
func (d *RMQ) SendReply(topic string) error { return nil }

//PublishTo publish to a topic
func (d *RMQ) PublishTo(topic string) error { return nil }

//Publish to an exchange or queue
func (d *RMQ) Publish(model interface{}) error { return nil }

//Subscribe to a queue or exchange
func (d *RMQ) Subscribe(model interface{}) interface{} { return nil }
