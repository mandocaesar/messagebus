package message

//ProducerMessage kafka producer payload
type ProducerMessage struct {
	Topic   string
	Key     []byte
	Message []byte
}

//SubscribeMessage kafka
type SubscribeMessage struct {
	Topic string
	Group string
	Fn    func(interface{})
}
