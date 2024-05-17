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
}

func TestDecodeBase64String(t *testing.T) {
	decoded, err := DecodeBase64String("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9")
	if err != nil {
		t.Errorf("Error decoding base64 string: %s", err)
	}
	assert.Equal(t, decoded, []byte(`{"alg":"HS256","typ":"JWT"}`))
}
