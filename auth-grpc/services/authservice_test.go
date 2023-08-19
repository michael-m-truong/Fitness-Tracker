package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	assert := assert.New(t)
	password := "mysecretpassword"
	hash, err := HashPassword(password)
	assert.NoError(err, "HashPassword returned an error")

	// Check if the hash is a valid bcrypt hash
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	assert.NoError(err, "Generated hash is not valid")
}

func TestCheckPasswordHash(t *testing.T) {
	assert := assert.New(t)
	password := "mysecretpassword"
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	assert.NoError(err, "Error generating hash")

	// Test a correct password
	match := CheckPasswordHash(password, string(hash))
	assert.True(match, "CheckPasswordHash returned false for correct password")

	// Test an incorrect password
	incorrectPassword := "wrongpassword"
	match = CheckPasswordHash(incorrectPassword, string(hash))
	assert.False(match, "CheckPasswordHash returned true for incorrect password")
}

func TestCheckPasswordHashInvalidHash(t *testing.T) {
	assert := assert.New(t)
	password := "mysecretpassword"
	invalidHash := "invalidhash"

	match := CheckPasswordHash(password, invalidHash)
	assert.False(match, "CheckPasswordHash returned true for invalid hash")
}

func TestCheckPasswordHashEmptyPassword(t *testing.T) {
	assert := assert.New(t)
	password := ""
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	assert.NoError(err, "Error generating hash")

	match := CheckPasswordHash(password, string(hash))
	assert.True(match, "CheckPasswordHash returned false for empty password")
}
