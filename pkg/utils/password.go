package utils

import "crypto"

func GetPassword(origin string, salts ...string) (string, string) {
	var salt string
	switch len(salts) {
	case 0:
		salt = GetRandomString(CHARS, 4)
	case 1:
		salt = salts[0]
	default:
		panic("salts length <= 1")
	}
	md5 := NewHasher(crypto.MD5)
	return md5.Hash(md5.Hash(md5.Hash(origin + salt))), salt
}

func EqualPassword(origin, salt, pwd string) bool {
	v, _ := GetPassword(origin, salt)
	return v == pwd
}
