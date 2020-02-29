package common

//Config interface for driver
type Config interface {
	Instantiate()
	Set(key string, value interface{}) (bool, error)
	Get(key string) interface{}
}
