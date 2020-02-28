package driver

import common "github.com/mandocaesar/messagebus/Common"

//Driver interface
type Driver interface {
	SetConfig() common.Config
	Connect() error
	SendReply(topic string) error
	PublishTo(topic string)
	Publish(model interface{}) (interface{}, error)
	Subscribe(model interface{}) interface{}
}
