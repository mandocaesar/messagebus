package main

import (
	"os"

	"path/filepath"

	"github.com/mandocaesar/messagebus/driver"

	"github.com/sirupsen/logrus"
)

//MessageBus abstraction struct for message bus
type MessageBus struct {
	service       *driver.Driver
	Channel       string
	Schemas       []string
	Handlers      map[string]func(data interface{})
	Subscriptions []string
	Publisher     []string
}

//NewMessageBus intantiate new MessageBus instance with specific driver
func NewMessageBus(service *driver.Driver, avroSchemaFolder string) (*MessageBus, error) {
	var files []string
	err := filepath.Walk(avroSchemaFolder, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	return &MessageBus{service: service, Schemas: files}, nil
}

//Publish a message to message queue
func (m *MessageBus) Publish(model string) (interface{}, error) {
	result, err := m.service.Publish(model)
	if err != nil {
		logrus.Error(err)
	}
	return result, err
}

//RegisterHandler function to register handler function when an event occured
func (m *MessageBus) RegisterHandler(key string, function func(data interface{})) {
	m.Handlers[key] = function
}
