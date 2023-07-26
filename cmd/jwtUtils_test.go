package cmd

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestGenerateSymmetric(t *testing.T) {
	symmetric, token := GenerateSymmetric("sekret", map[string]string{}, jwt.SigningMethodHS256)
	assert.NotNil(t, symmetric)
	token, err := jwt.Parse(symmetric, func(token *jwt.Token) (interface{}, error) {
		return []byte("sekret"), nil
	})

	assert.Nil(t, err)
	assert.NotNil(t, token)
}

func TestGenerateSymetricWithCustomClaims(t *testing.T) {
	var claims = make(map[string]string)
	claims["firstName"] = "David"

	symmetric, token := GenerateSymmetric("sekret", claims, jwt.SigningMethodHS256)
	fmt.Println(symmetric)
	fmt.Println(token)

	parse := Parse(symmetric)
	fmt.Println(parse)
	fmt.Println(reflect.TypeOf(parse.Claims))
	c := parse.Claims.(jwt.MapClaims)
	MapClaimsToString(c)
	i := c["CustomClaims"]
	m := i.(map[string]interface{})
	fmt.Println(m["firstName"])
}
