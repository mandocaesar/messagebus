package driver

import (
	"context"
	"errors"
	"mandocaesar/messagebus/common"
	"mandocaesar/messagebus/message"
	"mandocaesar/messagebus/serializer"
	"strings"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/snappy"
	_ "github.com/segmentio/kafka-go/snappy"
	"github.com/sirupsen/logrus"
)

//Kafka kafka driver struct
type Kafka struct {
	kafkaConfig *common.KafkaConfig

	//Kafka Consumer
	ConsumerConfig kafka.ReaderConfig
	Consumer       *kafka.Reader
	//Kafka Producer
	ProducerConfig kafka.WriterConfig
	Producer       *kafka.Writer
}

func (d *Kafka) initiateProducer() {
	d.ProducerConfig = kafka.WriterConfig{
		Brokers:          strings.Split(d.kafkaConfig.Get("brokers"), ";"),
		Balancer:         &kafka.LeastBytes{},
		CompressionCodec: snappy.NewCompressionCodec(),
		// BatchTimeout:     10 * time.Millisecond,
		// WriteTimeout:     c.TimeOut,
		// ReadTimeout:      c.TimeOut,
	}

}

func (d *Kafka) initiateConsumer() {

	d.ConsumerConfig = kafka.ReaderConfig{
		Brokers:  strings.Split(d.kafkaConfig.Get("brokers"), ";"),
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
		// MaxWait:         30 * time.Millisecond, // Maximum amount of time to wait for new data to come when fetching batches of messages from kafka.
		// ReadLagInterval: -1,
	}
}

//NewKafka intantiate new kafka driver
func NewKafka(config common.Config) (Driver, error) {
	kafkaConfig := config.(*common.KafkaConfig)

	driver := &Kafka{kafkaConfig: kafkaConfig}
	driver.initiateProducer()
	driver.initiateConsumer()

	return driver, nil
}

//SetConfig set config configuration
func (d *Kafka) SetConfig(key string, value interface{}) common.Config {
	return nil
}

//Connect connect to rmq message broker
func (d *Kafka) Connect() interface{} {
	return d.Producer.Stats().Dials
}

//SendReply reply to request reply pattern
func (d *Kafka) SendReply(topic string) error { return errors.New("Not Implemented yet") }

//PublishTo publish to a topic
func (d *Kafka) PublishTo(topic string) error { return nil }

//Publish to an exchange or queue
func (d *Kafka) Publish(model interface{}) error {
	data := model.(*message.ProducerMessage)

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
	logrus.Infof(
		"Topic : %s, Key : %s, Value : %s",
		data.Topic, data.Key, data.Message)

	return nil
}

//Subscribe to a queue or exchange
func (d *Kafka) Subscribe(model interface{}, serializer serializer.Serializer, functions map[int32]func(data interface{}) (interface{}, error)) (interface{}, error) {
	data := model.(*message.SubscribeMessage)

	d.ConsumerConfig.Topic = data.Topic
	d.ConsumerConfig.GroupID = data.Group

	d.Consumer = kafka.NewReader(d.ConsumerConfig)
	defer d.Consumer.Close()

	ctx := context.Background()

	for {
		message, err := d.Consumer.FetchMessage(ctx)
		if err != nil {
			logrus.Error(err)
		}

		header, err := serializer.GetHeader(message.Value)

		result, err := functions[header.MessageType](message.Value)
		if err != nil {
			logrus.Error(err)
		} else {
			logrus.Info(result)
		}
		return result, err
	}
}
