package secmanagers

import (
	"crypto/sha256"
	"encoding/base64"
)

// Function used to hash and sault the given string
func HashNSault(text string) string {
	text = "vTT3TMXDpJPAJ38A9Jmn" + text + "Ifjnmw9OkkIJ597sSLKEf1U="
	gen := sha256.New()
	_, err := gen.Write([]byte(text))
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(gen.Sum(nil))
}
