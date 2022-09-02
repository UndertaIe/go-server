package utils

import (
	"math/rand"
	"time"
)

type Type string

const (
	// 随机字符串包含: 0-9，a-z,A-Z
	CHARS Type = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// 字母随机字符串包含: a-z,A-Z
	LETTERS Type = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// 数字随机字符串包含: 0-9
	NUMS Type = "0123456789"
)

func GetRandomString(str Type, n int) string {
	result := make([]byte, n)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	l := len(str)
	for i := 0; i < n; i++ {
		result = append(result, str[r.Intn(l)])
	}
	return string(result)
}
