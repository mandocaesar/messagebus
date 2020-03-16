package driver

import (
	common "mandocaesar/messagebus/common"
	"mandocaesar/messagebus/serializer"
)

//Driver interface
type Driver interface {
	SetConfig(key string, value interface{}) common.Config
	Connect() interface{}
	SendReply(topic string) error
	PublishTo(topic string) error
	Publish(model interface{}) error
	Subscribe(model interface{}, serializer serializer.Serializer, fn func(key string, data interface{}) (interface{}, error))
}
