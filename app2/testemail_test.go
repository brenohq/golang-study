package app2

import (
	"testing"
)

func TestIsEmail(t *testing.T) {
	_, err := IsEmail("hello")
	if err == nil {
		t.Error("hello is not an email")
	}

	_, err = IsEmail("brenosc2@hotmail.com")
	if err != nil {
		t.Error("brenosc2@hotmail.com is an email")
	}
}
