package secmanagers

import (
	"crypto/sha256"
	"encoding/base64"
	"math/rand"
	"time"
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

// Function to check if the two filled passwords fields are equal
func EqualString(textone, texttwo string) bool {
	for t := range textone {
		if textone[t] != texttwo[t] {
			return false
		}
	}
	return true
}

// Function to generate a long authcookie
// The code gets hashed and salted and returned
func AuthCookieGen() string {
	src := rand.New(rand.NewSource(time.Now().UnixNano()))
	var alfabeth string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890!?^£$%&()/_-:.;,§ç@"
	var out string = ""
	for n := 0; n < 64; n++ {
		out = string(alfabeth[src.Intn(len(alfabeth))]) + out
	}
	return HashNSault(out)
}
