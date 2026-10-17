package jwt_test

import (
	"shortly/pkg/jwt"
	"testing"
)

func TestJWTCreate(t *testing.T) {
	const email = "r2d2@2.com"
	jwtService := jwt.NewJWT("5ai9GzWPycJzlPRI-44VHi6pIpgJJ-Tvf5Ro_t4GCQI")
	token, err := jwtService.Create(jwt.JWTData{
		Email: email,
	})
	if err != nil {
		t.Fatal(err)
	}
	isValid, data := jwtService.Parse(token)
	if !isValid {
		t.Fatal("Token is invalid")
	}
	if data.Email != email {
		t.Fatalf("Email %s not equal %s", data.Email, email)
	}
}
