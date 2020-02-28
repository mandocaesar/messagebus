package main

import (
	"os"

	"path/filepath"

	driver "mandocaesar/messagebus/driver"

	"github.com/sirupsen/logrus"
)

//MessageBus abstraction struct for message bus
type MessageBus struct {
	Service       driver.Driver
	Channel       string
	Schemas       []string
	Handlers      map[string]func(data interface{})
	Subscriptions []string
	Publisher     []string
}

//NewMessageBus intantiate new MessageBus instance with specific driver
func NewMessageBus(service driver.Driver, avroSchemaFolder string) (*MessageBus, error) {
	var files []string
	err := filepath.Walk(avroSchemaFolder, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})

	if err != nil {
		panic(err)
	}
	return &MessageBus{Service: service, Schemas: files}, nil
}

//Publish a message to message queue
func (m *MessageBus) Publish(model interface{}) (interface{}, error) {
	result, err := m.Service.Publish(model)
	if err != nil {
		logrus.Error(err)
	}
	return result, err
}

//RegisterHandler function to register handler function when an event occured
func (m *MessageBus) RegisterHandler(key string, function func(data interface{})) {
	m.Handlers[key] = function
}
