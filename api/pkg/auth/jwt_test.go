package auth

import (
	"testing"

	"github.com/plutov/gitprint/api/pkg/git"
)

func TestJWTClaims(t *testing.T) {
	user := &git.User{
		Username:    "1",
		Email:       "2",
		AccessToken: "2",
	}

	jwt, err := FillJWT(user)
	if err != nil {
		t.Fatal(err)
	}

	session, err := ReadJWTClaims(jwt)
	if err != nil {
		t.Fatal(err)
	}
	if session.User == nil {
		t.Fatalf("expected user, got nil")
	}
	if session.User.Username != user.Username {
		t.Fatalf("expected %s, got %s", user.Username, session.User.Username)
	}
}
