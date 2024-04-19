package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSplitToken(t *testing.T) {
	token := "a.b.c"
	s1, s2, s3 := SplitToken(token)
	if s1 != "a" {
		t.Errorf("Expected a, got %s", s1)
	}

	if s2 != "b" {
		t.Errorf("Expected b, got %s", s2)
	}

	if s3 != "c" {
		t.Errorf("Expected c, got %s", s3)
	}
}

var sampleToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhdWQiLCJleHAiOjQ3MDE5NzQ0MDAsImlhdCI6MTU0NjEyODAwMCwiaXNzIjoiaXNzIiwibmJmIjoxNTQ2MjE0NDAwLCJzdWIiOiJzdWIifQ.tfZF2KvIs-ZkWuBZmS1Y4vezOJ2Qs7P_nJkFugwPS1"

func TestParseWithoutVerification(t *testing.T) {
	err := ParseWithoutVerification(sampleToken)
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
