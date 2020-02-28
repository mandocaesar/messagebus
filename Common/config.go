package common

//Config interface for driver
type Config interface {
	Set(key string, value interface{})
	Get(key string) string
	List() map[string]interface{}
}
