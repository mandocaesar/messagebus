package common

//Config interface for driver
type Config interface {
	Instantiate(path string)
	Set(key string, value interface{}) (bool, error)
	Get(key string) string
}
