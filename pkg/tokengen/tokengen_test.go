package tokengen

import (
	"testing"
)

func TestNew(t *testing.T) {
	token := New(16)
	if 16 != len(token) {
		t.Errorf("failed to generate valid token")
	}

	token = New(15)
	if 14 != len(token) {
		t.Errorf("failed to generate valid token")
	}
}
