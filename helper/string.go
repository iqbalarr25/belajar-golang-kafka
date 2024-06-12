package helper

import (
	"math/rand"
	"strings"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func GenerateRandomString(length int) string {
	random := make([]byte, length)
	for i := range random {
		random[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(random)
}

func RemoveSpace(str string) string {
	return strings.ReplaceAll(str, " ", "")
}
