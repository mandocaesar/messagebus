package serializer

import (
	"bytes"
	"mandocaesar/messagebus/message"

	"github.com/sirupsen/logrus"
	avro "gopkg.in/avro.v0"
)

//AvroSerializer struct
type AvroSerializer struct {
	schemas map[string]avro.Schema
}

//NewAvroSerializer return new avro serializer instance
func NewAvroSerializer() *AvroSerializer {
	schemas := map[string]avro.Schema{}
	return &AvroSerializer{schemas: schemas}
}

//GetAllSchema retrieve all registered schema
func (a *AvroSerializer) GetAllSchema(path string) interface{} {
	if path == "" {
		path = "../schemas/"
	}

	a.schemas = avro.LoadSchemas(path)

	return a.schemas
}

//GetSchema get schema by key
func (a *AvroSerializer) GetSchema(name string) interface{} {
	if name == "" {
		return nil
	}

	if len(a.schemas) > 0 {
		return a.schemas[name]
	}

	return nil
}

//GetHeader parse header into struct
func (a *AvroSerializer) GetHeader(data []byte) message.Header {
	//buffer := new(bytes.Buffer)
	//	schema := a.GetSchema("kata.MessageHeader")
	schema := a.schemas["kata.MessageHeader"]
	datumReader := avro.NewSpecificDatumReader()
	datumReader.SetSchema(schema)

	decoder := avro.NewBinaryDecoder(data)

	result := &message.Header{}
	if err := datumReader.Read(result, decoder); err != nil {
		logrus.Error(err)
	}

	return *result
}

//ParseSchema parse string schema into a avro format
func (a *AvroSerializer) ParseSchema(schema string) (interface{}, error) {
	return avro.ParseSchema(schema)
}

//Decode a avro binary into struct
func (a *AvroSerializer) Decode(data interface{}, asStruct interface{}) interface{} { return nil }

//Encode a schema into avro binary
func (a *AvroSerializer) Encode(data interface{}, schemaName string) ([]byte, error) {
	writer := avro.NewSpecificDatumWriter()
	writer.SetSchema(a.schemas[schemaName])

	var buf bytes.Buffer
	err := writer.Write(data, avro.NewBinaryEncoder(&buf))
	logrus.Error(err)
	return buf.Bytes(), err
}
