package utilities

import (
	"math/rand"
	"time"
)

// Generate a random 10 chars code from keys
func GenerateCode() string {
	keys := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	out := ""

	src := rand.New(rand.NewSource(time.Now().UnixNano()))

	for n := 0; n <= 9; n++ {
		out = out + string(keys[src.Intn(len(keys)-1)])
	}
	return out
}
