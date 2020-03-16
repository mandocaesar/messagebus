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
func NewAvroSerializer() Serializer {
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
func (a *AvroSerializer) GetHeader(data []byte) (*message.MessageHeader, error) {
	return message.DeserializeMessageHeader(bytes.NewReader(data))
}

//Decode parse byte to struct
func (a *AvroSerializer) Decode(data []byte, schemaName string) (interface{}, error) {

	reader := avro.NewSpecificDatumReader()
	reader.SetSchema(a.schemas[schemaName])

	buffer := new(bytes.Buffer)
	header := new(message.MessageHeader)

	decoder := avro.NewBinaryDecoder(buffer.Bytes())

	err := reader.Read(header, decoder)

	return header, err
}

//ParseSchema parse string schema into a avro format
func (a *AvroSerializer) ParseSchema(schema string) (interface{}, error) {
	return avro.ParseSchema(schema)
}

//Encode a schema into avro binary
func (a *AvroSerializer) Encode(data interface{}, schemaName string) ([]byte, error) {
	writer := avro.NewSpecificDatumWriter()
	writer.SetSchema(a.schemas[schemaName])

	var buf bytes.Buffer
	err := writer.Write(data, avro.NewBinaryEncoder(&buf))

	logrus.Error(err)
	return buf.Bytes(), err
}
