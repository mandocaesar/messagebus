package serializer

//Serializer interface as abstraction for all serializer
type Serializer interface {
	GetAllSchema(path string) interface{}
	GetSchema(name string) interface{}
	ParseSchema(schema string) interface{}
	Decode(data interface{}, asStruct interface{}) interface{}
	Encode(data interface{}, schemaName string) interface{}
}
