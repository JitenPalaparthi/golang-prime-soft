package models_test

import (
	"http-echo-demo/models"
	"testing"
)

func TestUserValidateSuccess(t *testing.T) {
	user := &models.User{Name: "Jiten", Email: "Jitenp@outlool.com"}
	err := user.Validate()
	if err != nil {
		t.Fail()
	}
}

func TestUserValidateFail(t *testing.T) {
	user := &models.User{Email: "Jitenp@outlool.com"}
	err := user.Validate()
	if err == nil {
		t.Error("error should not be nil when Name field is empty")
	}
}
