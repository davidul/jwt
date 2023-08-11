package cmd

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncodeNoToken(t *testing.T) {
	b := new(bytes.Buffer)
	encodeCmd.SetOut(b)
	encodeCmd.SetErr(b)
	encodeCmd.Run(encodeCmd, []string{})
	assert.Equal(t, "Error: No token provided\n", b.String())
}

func TestEncode(t *testing.T) {
	b := new(bytes.Buffer)
	encodeCmd.SetOut(b)
	encodeCmd.SetErr(b)
	encodeCmd.Run(encodeCmd, []string{"{\"test\":\"test\"}"})
	assert.Equal(t, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0ZXN0IjoidGVzdCJ9.b5o0a23g5hGLIIcaqpfX3-oOuJVVRalv7xwkXRuHPG8\n", b.String())
}
