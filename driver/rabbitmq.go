package driver

import common "mandocaesar/messagebus/common"

type DriverRMQ struct {
	config common.Config
}

//SetConfig set config configuration
func (d *DriverRMQ) SetConfig(key string, value interface{}) common.Config {
	return nil
}

//Connect connect to rmq message broker
func (d *DriverRMQ) Connect() error { return nil }

//SendReply reply to request reply pattern
func (d *DriverRMQ) SendReply(topic string) error { return nil }

//PublishTo publish to a topic
func (d *DriverRMQ) PublishTo(topic string) error { return nil }

//Publish to an exchange or queue
func (d *DriverRMQ) Publish(model interface{}) error { return nil }

//Subscribe to a queue or exchange
func (d *DriverRMQ) Subscribe(model interface{}) interface{} { return nil }
