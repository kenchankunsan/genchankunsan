package kenchankunsan

import (
	"bytes"
	"math/rand"
	"time"
	"unicode/utf8"
)

const KEN = "けん"

var SUFFIXES = [3]string{
	"ちゃん",
	"くん",
	"さん",
}

func Kenchankunsan(length int) string {
	var ken bytes.Buffer

	rand.Seed(time.Now().UnixNano())
	suffixesLength := len(SUFFIXES)

	ken.WriteString(KEN)

	for utf8.RuneCount(ken.Bytes()) < length {
		i := rand.Intn(suffixesLength)
		ken.WriteString(SUFFIXES[i])
	}

	ken.WriteString("…")

	return ken.String()
}
