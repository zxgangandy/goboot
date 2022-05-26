package utils

import (
	"encoding/json"
	"math/rand"
	"time"
	"unsafe"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var (
	chars    = "0123456789abcdefghijklmnopqrstuvwxyz"
	charsLen = len(chars)
	rng      = rand.NewSource(time.Now().UnixNano())
	mask     = int64(1<<6 - 1)
)

func RandomString(n int) string {
	buf := make([]byte, n)
	for idx, cache, remain := n, rng.Int63(), 10; idx > 0; {
		if remain == 0 {
			cache, remain = rng.Int63(), 10
		}
		buf[idx-1] = chars[int(cache&mask)%charsLen]
		cache >>= 6
		remain--
		idx--
	}
	return *(*string)(unsafe.Pointer(&buf))
}

// Convert json string to map
func JsonToMap(jsonStr string) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// Convert map json string
func MapToJson(m map[string]interface{}) (string, error) {
	jsonByte, err := json.Marshal(m)
	if err != nil {
		return "", err
	}

	return string(jsonByte), nil
}
