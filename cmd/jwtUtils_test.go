package cmd

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestGenerateSymmetric(t *testing.T) {
	symmetric, _ := GenerateSymmetric("sekret", map[string]string{})
	assert.NotNil(t, symmetric)
	parse, err := jwt.Parse(symmetric, nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(parse.Method)
}

func TestGenerateSymetricWithCustomClaims(t *testing.T) {
	var claims = make(map[string]string)
	claims["firstName"] = "David"

	symmetric, token := GenerateSymmetric("sekret", claims)
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
