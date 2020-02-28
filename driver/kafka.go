package driver

import common "mandocaesar/messagebus/common"

type DriverKafka struct {
	config common.Config
}

//SetConfig set config configuration
func (d *DriverKafka) SetConfig() common.Config {
	return nil
}

//Connect connect to rmq message broker
func (d *DriverKafka) Connect() error { return nil }

//SendReply reply to request reply pattern
func (d *DriverKafka) SendReply(topic string) error { return nil }

//PublishTo publish to a topic
func (d *DriverKafka) PublishTo(topic string) error { return nil }

//Publish to an exchange or queue
func (d *DriverKafka) Publish(model interface{}) (interface{}, error) { return nil, nil }

//Subscribe to a queue or exchange
func (d *DriverKafka) Subscribe(model interface{}) interface{} { return nil }
