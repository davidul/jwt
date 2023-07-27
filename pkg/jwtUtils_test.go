package pkg

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateSymmetric(t *testing.T) {
	symmetric, token := GenerateSymmetric(DEFAULT_SECRET, map[string]string{}, jwt.SigningMethodHS256)
	assert.NotNil(t, symmetric)
	assert.NotNil(t, token)
	assert.Equal(t, token.Header["alg"], "HS256")
	assert.Equal(t, token.Header["typ"], "JWT")
	token, err := jwt.Parse(symmetric, func(token *jwt.Token) (interface{}, error) {
		return []byte(DEFAULT_SECRET), nil
	})

	assert.Nil(t, err)
	assert.NotNil(t, token)
}

func TestGenerateSymmetricWithCustomClaims(t *testing.T) {
	var claims = make(map[string]string)
	claims["firstName"] = "David"

	symmetric, token := GenerateSymmetric(DEFAULT_SECRET, claims, jwt.SigningMethodHS256)
	assert.NotNil(t, symmetric)
	assert.NotNil(t, token)
	assert.Equal(t, token.Header["alg"], "HS256")
	assert.Equal(t, token.Header["typ"], "JWT")
	assert.Equal(t, token.Claims.(CustomMapClaims).CustomClaims["firstName"], "David")

	parsedToken := Parse(symmetric, DEFAULT_SECRET)
	assert.True(t, parsedToken.Valid)

}

func TestGenerateSigned(t *testing.T) {
	var claims = make(map[string]string)
	signed := GenerateSigned(claims)
	fmt.Println(signed)
}

func TestEncode(t *testing.T) {
	encode := Encode("{\"sub\":\"1234567890\",\"name\":\"John Doe\",\"admin\":true}", DEFAULT_SECRET)
	assert.NotNil(t, encode)
}
