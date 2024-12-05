package utils

func IsKeyInMap(m map[interface{}]interface{}, key interface{}) bool {
	_, ok := m[key]
	return ok
}
