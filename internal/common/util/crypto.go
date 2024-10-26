package util

import (
	"crypto/sha256"
	"encoding/base64"
)

// EncryptWithSalt encrypts the given plaintext using salt.
func EncryptWithSalt(plaintext, salt string) string {
	// Concatenate the plaintext and salt
	data := append([]byte(plaintext), []byte(salt)...)

	// Calculate SHA256 hash
	hash := sha256.Sum256(data)

	// Encode the hash to base64 for easy representation
	encrypted := base64.StdEncoding.EncodeToString(hash[:])

	return encrypted
}
