package auth

import (
	"testing"
)

func TestLogin(t *testing.T) {
	_, e := Login()

	if e != nil {
		t.Errorf("Failed to login %s", e)
	}
}

func TestAwsUser_GetToken(t *testing.T) {
	user, _ := Login()

	token, _ := user.GetToken()

	if token.Username != "AWS" {
		t.Errorf("Bad token %v", token)
	}
}
