package auth

import "testing"

func TestHashPassword(t *testing.T) {
	hash, err := HashPassword("lalala")
	if err != nil {
		t.Errorf("error hashing password: %v", err)
	}

	if hash == "" {
		t.Error("expected hash pass cannot be empty")
	}

	if hash == "password" {
		t.Error("expected hash pass to be different from 'password'")
	}
}

func TestComparePassword(t *testing.T) {
	hash, err := HashPassword("password")
	if err != nil {
		t.Errorf("error hashing password: %v", err)
	}

	if !CompareHashedPassword(hash, []byte("password")) {
		t.Errorf("expected password same to hash")
	}

	if !CompareHashedPassword(hash, []byte("notpassword")) {
		t.Errorf("expected password not same to hash")
	}
}
