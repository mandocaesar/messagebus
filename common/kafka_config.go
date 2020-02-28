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
	case "brokerUrls":
		k.BrokerURLs = value.([]string)
	case "timeout":
		k.TimeOut = value.(time.Duration)
	default:
		err = errors.New("config not found on kafka config struct")
	}
	return result, err
}

//Get return specific kafka config by key
func (k *KafkaConfig) Get(key string) (interface{}, error) {
	switch key {
	case "brokerUrls":
		return k.BrokerURLs, nil
	case "timeout":
		return k.TimeOut, nil
	default:
		return nil, errors.New("config not found")
	}
}
