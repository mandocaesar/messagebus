package bus

import (
	"os"

	"path/filepath"

	driver "mandocaesar/messagebus/driver"
	"mandocaesar/messagebus/serializer"

	"github.com/sirupsen/logrus"
)

//MessageBus abstraction struct for message bus
type MessageBus struct {
	Service       driver.Driver
	Serializer    serializer.Serializer
	Channel       string
	Schemas       []string
	Handlers      map[int]func(data interface{}) (interface{}, error)
	Subscriptions []string
	Publisher     []string
}

//NewMessageBus intantiate new MessageBus instance with specific driver
func NewMessageBus(service driver.Driver, avroSchemaFolder string) (*MessageBus, error) {
	var files []string
	handlers := make(map[int]func(data interface{}) (interface{}, error))
	err := filepath.Walk(avroSchemaFolder, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	return &MessageBus{Service: service, Schemas: files, Handlers: handlers}, nil
}

//Publish a message to message queue
func (m *MessageBus) Publish(model interface{}) error {
	err := m.Service.Publish(model)
	if err != nil {
		logrus.Error(err)
	}
	return err
}

//Subscribe subscribe to a topic
func (m *MessageBus) Subscribe(model interface{}) error {
	err := m.Service.Subscribe(model, m.Serializer)
}

//RegisterHandler function to register handler function when an event occured
func (m *MessageBus) RegisterHandler(key int, function func(data interface{}) (interface{}, error)) {
	m.Handlers[key] = function
}

//Handle functions
func (m *MessageBus) Handle(key int, data interface{}) (interface{}, error) {
	return m.Handlers[key](data)
}
