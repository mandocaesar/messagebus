package driver

import (
	common "mandocaesar/messagebus/common"

	"github.com/segmentio/kafka-go"
	_ "github.com/segmentio/kafka-go/snappy"
)

//DriverKafka kafka driver struct
type DriverKafka struct {
	config       common.Config
	ReaderConfig kafka.ReaderConfig
	Reader       *kafka.Reader

	Config    kafka.WriterConfig
	Publisher *kafka.Writer
}

//NewDriverKafka intantiate new kafka driver
func NewDriverKafka(config *common.Config) (*DriverKafka, error) {
	return nil, nil
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
