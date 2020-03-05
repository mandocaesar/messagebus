package serializer

import "mandocaesar/messagebus/message"

//Serializer interface as abstraction for all serializer
type Serializer interface {
	GetAllSchema(path string) interface{}
	GetSchema(name string) interface{}
	GetHeader(data []byte) (*message.MessageHeader, error)
	ParseSchema(schema string) (interface{}, error)
	Decode(data []byte) (interface{}, error)
	Encode(data interface{}, schemaName string) ([]byte, error)
}
