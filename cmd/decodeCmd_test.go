package cmd

import (
	"bytes"
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
	decodeCmd.Help()
	assert.Equal(t, ErrorNoToken+h.String(), b.String())
}

func TestDecodeCmdWithToken(t *testing.T) {
	b := new(bytes.Buffer)
	decodeCmd.SetOut(b)
	decodeCmd.SetErr(b)
	decodeCmd.Run(decodeCmd, []string{"eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhdWQiLCJleHAiOjE3MjIxNzg5MjksImlhdCI6MTY5MDM4MzcyOSwiaXNzIjoiaXNzIiwibmJmIjoxNjkwNDcwMTI5LCJzdWIiOiJzdWIifQ.puww8DUW_MhVUUzEBUmJf-t7j0jnJlYcF3ftD2BmJLJINZpfnTAwdoeFf7y0n4Hd0nAO7QKql6XN0PqlIRdph8LQr-SR_WXNVUe_8trfmQA-Zxrp-M8WCLV8msgt8waDs6_uXmi1IJiOJVB2ryNs2tEZhwLztifGN1TCU8YU2sbkP9g_Yz7zOw6BFulWiv-am2eHbxMOQeE16-i3in_JpLqT-ypn6o5zNNiYKyVFGeDftKNXk5bQPnDmWg_5mwkZi1ybqGJdy6RsUGQ8PBMPGKsM7JCvrQw8DQEcDMMQ--nZLNtkqk0BHxM7VAG-Vgs7Hz2JFQLmFQKwXgHwRt_ojg"})
	assert.Equal(t, "{\n  \"alg\": \"RS512\",\n  \"typ\": \"JWT\"\n}{\n  \"aud\": \"aud\",\n  \"exp\": 1722178929,\n  \"iat\": 1690383729,\n  \"iss\": \"iss\",\n  \"nbf\": 1690470129,\n  \"sub\": \"sub\"\n}\n", b.String())
}

func TestDecodeCmdBadToken(t *testing.T) {
	b := new(bytes.Buffer)
	decodeCmd.SetOut(b)
	decodeCmd.SetErr(b)
	decodeCmd.Run(decodeCmd, []string{"test"})
	assert.Equal(t, "token contains an invalid number of segments\n", b.String())
}
