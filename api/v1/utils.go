package v1

import (
	"crypto/sha256"
	"encoding/hex"
)

// hashString hashes a string using SHA-256
func hashString(text string) string {
	// Convert the email string to bytes
	emailBytes := []byte(text)

	// Create a new SHA-256 hasher
	hasher := sha256.New()

	// Write the email bytes to the hasher
	hasher.Write(emailBytes)

	// Get the hashed result as a byte slice
	hashedBytes := hasher.Sum(nil)

	// Convert the hashed byte slice to a hexadecimal string
	hashedEmail := hex.EncodeToString(hashedBytes)

	return hashedEmail
}
