package utils

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
)

const SALTSIZE = 16

func GenerateRandomSalt(saltSize int) ([]byte, error) {
	var salt = make([]byte, SALTSIZE)

	_, err := rand.Read(salt[:])
	if err != nil {
		return nil, err
	}
	return salt, nil
}

func HashPassword(password string, salt []byte) (string, error) {
	var err error

	passwordBytes := []byte(password)
	sha512Hasher := sha512.New()
	passwordBytes = append(passwordBytes, salt...)

	_, err = sha512Hasher.Write(passwordBytes)
	if err != nil {
		return "", err
	}

	hashedPasswordBytes := sha512Hasher.Sum(nil)

	return hex.EncodeToString(hashedPasswordBytes), nil
}

func DoPasswordMatch(hashedPassword, currPassword string, salt []byte) (bool, error) {
	currPasswordHash, err := HashPassword(currPassword, salt)
	if err != nil {
		return false, err
	}

	return hashedPassword == currPasswordHash, nil
}
