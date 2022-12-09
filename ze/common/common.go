package common

import "log"

// TODO: Implement
func ValidateAuthToken(auth string) {
	if len(auth) == 0 {
		log.Fatal("Auth token must be set using the -auth flag")
	}
}

// TODO: Implement
func ValidateZapiUrl(url string) {
	if len(url) == 0 {
		log.Fatal("URL must be set using the -url flag")
	}
}
