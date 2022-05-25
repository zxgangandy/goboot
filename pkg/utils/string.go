package utils

import (
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
