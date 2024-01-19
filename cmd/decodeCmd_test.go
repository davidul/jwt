package cmd

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncodeCmd(t *testing.T) {
	encodeCmd.Run(encodeCmd, []string{"test"})
}

func TestDecodeCmdNoToken(t *testing.T) {
	b := new(bytes.Buffer)
	decodeCmd.SetOut(b)
	decodeCmd.SetErr(b)
	decodeCmd.Run(decodeCmd, []string{})

	h := new(bytes.Buffer)
	decodeCmd.SetOut(h)
	decodeCmd.SetErr(h)
	err := decodeCmd.Help()
	if err != nil {
		fmt.Print(err)
		return
	}
	assert.Equal(t, ErrorNoToken+h.String(), b.String())
}

func TestDecodeCmdWithToken(t *testing.T) {
	b := new(bytes.Buffer)
	decodeCmd.SetOut(b)
	decodeCmd.SetErr(b)
	decodeCmd.Run(decodeCmd, []string{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhdWQiLCJleHAiOjE3MjMzNzg0NjcsImlhdCI6MTY5MTU4MzI2NywiaXNzIjoiaXNzIiwibmJmIjoxNjkxNjY5NjY3LCJzdWIiOiJzdWIifQ.O_hh57if9NwiY_4qeviAXwDAh7IIvQdGuV4YWZ3qsmI"})
	assert.Equal(t, "{\n  \"alg\": \"HS256\",\n  \"typ\": \"JWT\"\n}{\n  \"aud\": \"aud\",\n  \"exp\": 1723378467,\n  \"iat\": 1691583267,\n  \"iss\": \"iss\",\n  \"nbf\": 1691669667,\n  \"sub\": \"sub\"\n}\n", b.String())
}

func TestDecodeCmdBadToken(t *testing.T) {
	b := new(bytes.Buffer)
	decodeCmd.SetOut(b)
	decodeCmd.SetErr(b)
	decodeCmd.Run(decodeCmd, []string{"test"})
	assert.Equal(t, "token contains an invalid number of segments\n", b.String())
}
