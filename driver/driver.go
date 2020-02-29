package driver

import common "mandocaesar/messagebus/common"

//Driver interface
type Driver interface {
	SetConfig(key string, value interface{}) common.Config
	Connect() error
	SendReply(topic string) error
	PublishTo(topic string) error
	Publish(model interface{}) error
	Subscribe(model interface{}) interface{}
}
