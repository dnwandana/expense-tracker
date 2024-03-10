package utils

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
)

// GenerateNanoID generates a random string with the given length
func GenerateNanoID(length int) string {
	id, err := gonanoid.Generate("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz", length)
	PanicIfError(err)

	return id
}
