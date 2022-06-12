package utils

import (
	"crypto/rand"
	"math/big"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GetUniqueString(n int) string {
	b := make([]rune, n)
	for i := range b {
		if next, err := rand.Int(rand.Reader, big.NewInt(int64(len(letterRunes)))); err == nil {
			b[i] = letterRunes[next.Int64()]
		}
	}

	return string(b)
}
