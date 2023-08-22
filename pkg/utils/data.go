package utils

func GetMapDefault(k, defaultValue interface{}, m map[interface{}]interface{}) interface{} {
	if value, ok := m[k]; ok {
		return value
	} else {
		return defaultValue
	}
}
