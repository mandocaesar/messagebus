package message

import uuid "github.com/satori/uuid"

//Header struct for Message header
type Header struct {
	MessageType   int       `json:"message_type"`
	MessageFlag   int       `json:"message_flag"`
	MessageID     uuid.UUID `json:"message_id"`
	CorrelationID uuid.UUID `json:"correlation_id"`
	ReturnAddress string    `json:"return_address"`
}
