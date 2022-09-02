package utils

import (
	"math/rand"
	"time"
)

// 字符类型
type CharType string

const (
	// 随机字符串包含: 0-9，a-z,A-Z
	CHARS CharType = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// 字母随机字符串包含: a-z,A-Z
	LETTERS CharType = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// 数字随机字符串包含: 0-9
	NUMS CharType = "0123456789"
)

func GetRandomString(str CharType, n int) string {
	result := make([]byte, 0, n) // len=0, cap=n的切片
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	l := len(str)
	for i := 0; i < n; i++ {
		result = append(result, str[r.Intn(l)])
	}
	return string(result)
}
