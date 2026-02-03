package generator

import (
	"crypto/sha256"

	"github.com/mr-tron/base58"
)

func sha256f(url string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(url))
	return algorithm.Sum(nil)
}

func base58Encoded(b []byte) string {
	return base58.Encode(b)
}

func GenerateShortLink(url, userId string) string {
	urlHashBytes := sha256f(url + userId)
	return base58Encoded(urlHashBytes[:8])
}
