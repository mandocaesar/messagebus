package common

import (
	"errors"
	"os"
	"time"

	dotenv "github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
)

//KafkaConfig is to configure kafka connections
type KafkaConfig struct {
	BrokerURLs string
	TimeOut    time.Duration
}

//Instantiate initiatlize kafka config
func (k *KafkaConfig) Instantiate(path string) {
	// A duration string is a possibly signed sequence of
	// decimal numbers, each with optional fraction and a unit suffix,
	// such as "300ms", "-1.5h" or "2h45m".
	// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
	if len(path) > 0 {
		err := dotenv.Load(path)
		if err != nil {
			logrus.Panic(err)
		}
	}
	duration, _ := time.ParseDuration(os.Getenv("TIMEOUT"))

	k.BrokerURLs = os.Getenv("BROKERS")
	k.TimeOut = duration
}

//Set save key pair value of a kafka config
func (k *KafkaConfig) Set(key string, value interface{}) (bool, error) {
	result := false
	var err error
	switch key {
	case "brokers":
		k.BrokerURLs = value.(string)
	case "timeout":
		k.TimeOut = value.(time.Duration)
	default:
		err = errors.New("config not found on kafka config struct")
	}
	return result, err
}

//Get return specific kafka config by key
func (k *KafkaConfig) Get(key string) string {
	switch key {
	case "brokers":
		return k.BrokerURLs
	case "timeout":
		return k.TimeOut.String()
	default:
		return ""
	}
}
