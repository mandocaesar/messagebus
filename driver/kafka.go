package driver

import (
	"context"
	"mandocaesar/messagebus/common"
	"mandocaesar/messagebus/message"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/snappy"
	_ "github.com/segmentio/kafka-go/snappy"
	"github.com/sirupsen/logrus"
)

//DriverKafka kafka driver struct
type DriverKafka struct {
	kafkaConfig    common.KafkaConfig
	ConsumerConfig kafka.ReaderConfig
	Consumer       *kafka.Reader

	ProducerConfig kafka.WriterConfig
	Producer       *kafka.Writer
}

func (d *DriverKafka) initiateProducer() {
	d.ProducerConfig = kafka.WriterConfig{
		Brokers:          d.kafkaConfig.Get("brokers").([]string),
		Balancer:         &kafka.LeastBytes{},
		CompressionCodec: snappy.NewCompressionCodec(),
		// BatchTimeout:     10 * time.Millisecond,
		// WriteTimeout:     c.TimeOut,
		// ReadTimeout:      c.TimeOut,
	}

}

//NewDriverKafka intantiate new kafka driver
func NewDriverKafka(config common.Config) (*DriverKafka, error) {
	kafkaConfig := config.(*common.KafkaConfig)

	driver := &DriverKafka{kafkaConfig: *kafkaConfig}
	driver.initiateProducer()
	return driver, nil
}

//SetConfig set config configuration
func (d *DriverKafka) SetConfig(key string, value interface{}) common.Config {
	return nil
}

//Connect connect to rmq message broker
func (d *DriverKafka) Connect() error { return nil }

//SendReply reply to request reply pattern
func (d *DriverKafka) SendReply(topic string) error { return nil }

//PublishTo publish to a topic
func (d *DriverKafka) PublishTo(topic string) error { return nil }

//Publish to an exchange or queue
func (d *DriverKafka) Publish(model interface{}) error {
	data := model.(message.ProducerMessage)

	d.ProducerConfig.Topic = data.Topic
	d.Producer = kafka.NewWriter(d.ProducerConfig)
	defer d.Producer.Close()

	ctx := context.Background()
	msg := kafka.Message{
		Key:   data.Key,
		Value: data.Message,
	}
	if err := d.Producer.WriteMessages(ctx, msg); err != nil {
		logrus.Error(err)
		return err
	}

	logrus.Info("Message published")
	logrus.Info(
		"Topic : "+data.Topic,
		" Key : "+string(data.Key),
		" Value : "+string(data.Message),
	)
	return nil
}

//Subscribe to a queue or exchange
func (d *DriverKafka) Subscribe(model interface{}) interface{} { return nil }
