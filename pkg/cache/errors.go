package cache

import (
	"fmt"
	"strings"
)

type CacheError struct {
	code    int
	details []string
}

func NewCacheError(code int, ds ...string) CacheError {
	c := CacheError{code: code}
	return c.WithDetails(ds...)
}

func (c CacheError) Error() string {
	var msg string
	if c.details != nil {
		msg = strings.Join(c.details, ", ")
	} else {
		msg = ""
	}
	return fmt.Sprintf("Cache init error:( %s", msg)
}

func (c CacheError) WithDetails(ds ...string) CacheError {
	c.details = append(c.details, ds...)
	return c
}

func (c CacheError) Code() int {
	return c.code
}

func (c CacheError) Equal(c2 error) bool {
	if err, ok := c2.(CacheError); ok {
		return c.code == err.code
	}
	return false
}
