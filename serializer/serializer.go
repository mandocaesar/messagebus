package serializer

import "mandocaesar/messagebus/message"

//Serializer interface as abstraction for all serializer
type Serializer interface {
	GetAllSchema(path string) interface{}
	GetSchema(name string) interface{}
	GetHeader(data []byte) message.Header
	ParseSchema(schema string) interface{}
	Decode(data interface{}, asStruct interface{}) interface{}
	Encode(data interface{}, schemaName string) ([]byte, error)
}
