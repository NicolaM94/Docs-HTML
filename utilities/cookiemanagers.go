package utilities

import (
	"net/http"

	"github.com/gorilla/securecookie"
)

// Secure cookie instance used in latter functions
var hashKey = []byte(securecookie.GenerateRandomKey(64))
var blockKey = []byte(securecookie.GenerateRandomKey(32))
var s = securecookie.New(hashKey, blockKey)

// Generates a secure cookie, returning the cookie
func GenerateSecureCookie(name, value string) *http.Cookie {
	data := map[string]string{name: value}
	cookie := &http.Cookie{}
	if encoded, err := s.Encode(name, data); err == nil {
		cookie.Name = name
		cookie.Value = encoded
		cookie.Secure = true
		cookie.HttpOnly = true
	} else {
		panic(err)
	}
	return cookie
}

// Decodes a secure cookie, outs a name-value map and nil err if correct
func DecodeSecureCookie(cookieName string, r *http.Request) (out map[string]string, err error) {
	if cookie, err := r.Cookie(cookieName); err == nil {
		err = s.Decode(cookieName, cookie.Value, &out)
		if err == nil {
			return out, nil
		}
	}
	return nil, err
}
