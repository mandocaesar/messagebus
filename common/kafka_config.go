package common

import (
	"errors"
	"time"
)

//KafkaConfig is to configure kafka connections
type KafkaConfig struct {
	BrokerURLs []string
	TimeOut    time.Duration
}

//Set save key pair value of a kafka config
func (k *KafkaConfig) Set(key string, value interface{}) (bool, error) {
	result := false
	var err error
	switch key {
	case "brokers":
		k.BrokerURLs = value.([]string)
	case "timeout":
		k.TimeOut = value.(time.Duration)
	default:
		err = errors.New("config not found on kafka config struct")
	}
	return result, err
}

//Get return specific kafka config by key
func (k *KafkaConfig) Get(key string) interface{} {
	switch key {
	case "brokers":
		return k.BrokerURLs
	case "timeout":
		return k.TimeOut
	default:
		return nil
	}
}
