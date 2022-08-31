package utils

import (
	"math/rand"
	"time"
)

var ASCII = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GetRandomString(n int) string {
	result := make([]byte, n)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	l := len(ASCII)
	for i := 0; i < n; i++ {
		result = append(result, ASCII[r.Intn(l)])
	}
	return string(result)
}

func GetRandomNum(n int) string {
	return ""
}
