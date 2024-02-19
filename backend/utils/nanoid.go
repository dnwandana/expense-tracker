package utils

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func GenerateNanoID(length int) string {
	id, err := gonanoid.Generate("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz", length)
	if err != nil {
		panic(err)
	}

	return id
}
