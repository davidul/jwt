package cmd

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

var pre = []byte("=== Generating Simple Token ===")

func TestGenJwt(t *testing.T) {
	b := new(bytes.Buffer)
	genJwtCmd.SetOut(b)
	genJwtCmd.SetErr(b)
	genJwtCmd.Run(genJwtCmd, []string{})
	assert.Equal(t, pre, b.Bytes()[0:len(pre)])
}
