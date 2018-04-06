package rand

import (
	"testing"
)

const (
	TEST_RANDSTRING_LEN = 10
)

func TestRandString(t *testing.T) {
	if len(RandString(TEST_RANDSTRING_LEN)) != TEST_RANDSTRING_LEN {
		t.Fatal("rand string not equal!")
	}
}
