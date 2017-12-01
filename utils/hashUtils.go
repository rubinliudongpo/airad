package utils

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"crypto/rand"
	"golang.org/x/crypto/scrypt"
)

const pwHashBytes = 64

func Sha1(input string) string {
	if input == "" {
		return "adc83b19e793491b1c6ea0fd8b46cd9f32e592fc"
	}
	return fmt.Sprintf("%x", sha1.Sum([]byte(input)))
}

func Secret2Password(username, secret string) string {
	return Sha1(Sha1(secret[:8]) + Sha1(username) + Sha1(secret[8:]))
}

func Base64(input string) string {
	return base64.StdEncoding.EncodeToString([]byte(input))
}

func GenerateSalt() (salt string, err error) {
	buf := make([]byte, pwHashBytes)
	if _, err := io.ReadFull(rand.Reader, buf); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", buf), nil
}

func GeneratePassHash(password string, salt string) (hash string, err error) {
	h, err := scrypt.Key([]byte(password), []byte(salt), 16384, 8, 1, pwHashBytes)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h), nil
}
