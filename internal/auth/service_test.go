package auth_test

import (
	"shortly/internal/auth"
	"shortly/internal/user"
	"testing"
)

type MockUserRepository struct{}

func (repo *MockUserRepository) Create(u *user.User) (*user.User, error) {
	return &user.User{
		Email: "r2d2@.com",
	}, nil
}

func (repo *MockUserRepository) FindByEmail(email string) (*user.User, error) {
	return nil, nil
}

func TestRegisterSucces(t *testing.T) {
	const initialEmail = "r2d2@.com"
	authService := auth.NewAuthService(&MockUserRepository{})
	email, err := authService.Register(initialEmail, "1", "Yo")
	if err != nil {
		t.Fatal(err)
	}
	if email != initialEmail {
		t.Fatalf("Email %s do not match %s", email, initialEmail)
	}
}
