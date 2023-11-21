package utilities

import (
	"crypto/sha256"
	"encoding/base64"
)

// Returns the hash of the give str string.
func Hasher(str string) string {
	src := sha256.New()
	src.Sum([]byte(str))
	res := src.Sum(nil)

	return base64.StdEncoding.EncodeToString([]byte(res))
}

// Returns the hash of the given string applying salts as pre-determined.
func HashNSault(str string) string {
	var saltOne string = "Gx34eO095ipwl£ewodsbgTBd3"
	var saltTwo string = "79dsDGhqkbgfdHDK84IWPBvVM"
	var saltThree string = "y80hkvgfGFHS79vgfdSDEgd88"
	res := Hasher(str)
	res = Hasher(res + saltOne)
	res = Hasher(res + saltThree)
	res = Hasher(res + saltTwo)
	return base64.StdEncoding.EncodeToString([]byte(res))
}
