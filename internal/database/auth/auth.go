package auth

import (
	"errors"
	"net/http"
)

// get api key from the header
func GetApiKey(headers http.Header) (string, error) {
	// authorization should be in the header -> header.authorization
	// in the tutorial it said that the authorization param should be like: authorization: apikey {the actual key}
	// but I have not implemented that yet
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no api key found")
	}
	return val, nil
}