package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSplitToken(t *testing.T) {
	p := NewParser()
	token := "a.b.c"
	s, err := p.SplitToken(token)
	if err != nil {
		t.Errorf("Error splitting token: %s", err)
	}

	if s[0] != "a" {
		t.Errorf("Expected a, got %s", s[0])
	}

	if s[1] != "b" {
		t.Errorf("Expected b, got %s", s[1])
	}

	if s[2] != "c" {
		t.Errorf("Expected c, got %s", s[0])
	}
}

var sampleToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhdWQiLCJleHAiOjQ3MDE5NzQ0MDAsImlhdCI6MTU0NjEyODAwMCwiaXNzIjoiaXNzIiwibmJmIjoxNTQ2MjE0NDAwLCJzdWIiOiJzdWIifQ.tfZF2KvIs-ZkWuBZmS1Y4vezOJ2Qs7P_nJkFugwPS1"

func TestParseWithoutVerification(t *testing.T) {
	p := NewParser()
	err := p.ParseWithoutVerification(sampleToken)
	if err != nil {
		t.Errorf("Error parsing token: %s", err)
	}

	assert.Equal(t, p.header, `{"alg":"HS256","typ":"JWT"}`)
	assert.Equal(t, p.claims, `{"aud":"aud","exp":4701974400,"iat":1546128000,"iss":"iss","nbf":1546214400,"sub":"sub"}`)
	assert.Equal(t, p.signature, "")

	assert.Equal(t, p.headerMap["alg"], "HS256")
	assert.Equal(t, p.headerMap["typ"], "JWT")
	assert.Equal(t, p.claimsMap["aud"], "aud")
	assert.Equal(t, p.claimsMap["exp"], float64(4701974400))
}

func TestDecodeBase64String(t *testing.T) {
	decoded, err := DecodeBase64String("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9")
	if err != nil {
		t.Errorf("Error decoding base64 string: %s", err)
	}
	assert.Equal(t, decoded, []byte(`{"alg":"HS256","typ":"JWT"}`))

}
