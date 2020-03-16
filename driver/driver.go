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
	Subscribe(model interface{}, serializer serializer.Serializer, functions map[int32]func(data interface{}) (interface{}, error)) (interface{}, error)
}
