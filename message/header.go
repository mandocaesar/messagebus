package message

//Header struct for Message header
type Header struct {
	MessageType   int32  `avro:"messageType"`
	MessageFlags  int32  `avro:"messageFlags"`
	MessageID     string `avro:"messageId"`
	CorrelationID string `avro:"correlationId"`
	ReturnAddress string `avro:"returnAddress"`
}
