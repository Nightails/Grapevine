package auth

import (
	"testing"
)

func TestHashAndCheckPassword(t *testing.T) {
	tests := []struct {
		name          string
		password      string
		inputPassword string // Password to check against the hash
		wantMatch     bool
	}{
		{
			name:          "correct password",
			password:      "secure-password-123",
			inputPassword: "secure-password-123",
			wantMatch:     true,
		},
		{
			name:          "incorrect password",
			password:      "secure-password-123",
			inputPassword: "wrong-password",
			wantMatch:     false,
		},
		{
			name:          "empty password",
			password:      "",
			inputPassword: "",
			wantMatch:     true,
		},
		{
			name:          "special characters",
			password:      "p@$$w0rd!_with_spaces ",
			inputPassword: "p@$$w0rd!_with_spaces ",
			wantMatch:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash, err := HashPassword(tt.password)
			if err != nil {
				t.Fatalf("HashPassword() error = %v", err)
			}

			if hash == "" {
				t.Error("HashPassword() returned an empty string")
			}

			match := CheckPassword(tt.inputPassword, hash)
			if match != tt.wantMatch {
				t.Errorf("CheckPassword() match = %v, want %v", match, tt.wantMatch)
			}
		})
	}
}

func TestCheckPassword_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		password string
		hash     string
		want     bool
	}{
		{
			name:     "invalid hash format",
			password: "password",
			hash:     "not-a-valid-hash",
			want:     false,
		},
		{
			name:     "empty hash",
			password: "password",
			hash:     "",
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckPassword(tt.password, tt.hash); got != tt.want {
				t.Errorf("CheckPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
