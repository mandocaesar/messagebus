package message

//ProducerMessage kafka producer payload
type ProducerMessage struct {
	Topic   string
	Key     []byte
	Message []byte
}
