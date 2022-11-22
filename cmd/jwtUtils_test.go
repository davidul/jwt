package cmd

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateSymmetric(t *testing.T) {
	symmetric := GenerateSymmetric("sekret")
	assert.NotNil(t, symmetric)
	parse, err := jwt.Parse(symmetric, nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(parse.Method)
}
