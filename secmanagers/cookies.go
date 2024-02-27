package secmanagers

import (
	"log"
	"net/http"

	"github.com/gorilla/securecookie"
)

var hashKey = []byte("PX[UI^2bI1PMyKÂÂdyK9fO6RRq0?U6xN1g]3Fd$J28ShLN&jJlJ$[nVrG$Nfr!")
var secureSource = securecookie.New(hashKey, nil)

// Create a secure cookie using securecookie api from gorillamux
// Uses KEY to create an encoded version of the cookie
// Return the cookie to be set in the handler.
func CreateSecCk(name, value string) http.Cookie {
	enc, err := secureSource.Encode(name, value)
	if err != nil {
		log.Fatal(err)
	}
	cookie := http.Cookie{
		Name:  name,
		Value: enc,
	}
	return cookie
}

// Decodes the given cookie.
// Returns the cookie value as string
func DecodeSecCk(cookie http.Cookie) (out string) {
	err := secureSource.Decode(cookie.Name, cookie.Value, &out)
	if err != nil {
		log.Fatal(err)
	}
	return
}
