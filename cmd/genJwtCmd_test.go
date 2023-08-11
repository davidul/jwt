package cmd

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

var s = "=== Generating Simple Token ===\nHeader\n\ttyp : JWT \n\talg : HS256 \nStandard Claims:\n\texp : 1970-01-20T23:42:50+01:00 \n\tiat : 1970-01-20T14:52:54+01:00 \n\tiss : iss \n\tnbf : 1970-01-20T14:54:21+01:00 \n\tsub : sub \n\taud : aud \n \nSigned string: \neyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhdWQiLCJleHAiOjE3MjMzNzAxMjUsImlhdCI6MTY5MTU3NDkyNSwiaXNzIjoiaXNzIiwibmJmIjoxNjkxNjYxMzI1LCJzdWIiOiJzdWIifQ.agCBDtpH6bCpLZPQMzJMFJW2zYZe45TCYivUyG5QvaM"
var pre = []byte("=== Generating Simple Token ===")
var header = "\nHeader\n\talg : HS256 \n\ttyp : JWT \n"

func TestGenJwt(t *testing.T) {
	b := new(bytes.Buffer)
	genJwtCmd.SetOut(b)
	genJwtCmd.SetErr(b)
	genJwtCmd.Run(genJwtCmd, []string{})
	assert.Equal(t, pre, b.Bytes()[0:len(pre)])
}
