package util

import (
	"math/rand"
	"time"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var r = rand.NewSource(time.Now().Unix())

func RandInt() int {
	return int(r.Int63())
}

func RandString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[int(r.Int63())%(len(letters))]
	}
	return string(b)
}
