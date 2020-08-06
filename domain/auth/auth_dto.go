package auth

import (
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/flucas97/cng/cng-baguera-auth-api/utils/error_factory"
	uuid "github.com/satori/go.uuid"
)

type auth struct {
	Name string `json:"name"`
	Uuid string `json:"_uuid"`
}

func New(name string) *auth {
	au := &auth{}
	au.Uuid = uuid.NewV4().String()
	au.Name = name
	return au
}

func (au *auth) GenerateToken() (string, *error_factory.RestErr) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["account_name"] = au.Name
	claims["uuid"] = au.Uuid

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	newToken, err := token.SignedString([]byte(os.Getenv("API_SECRET")))
	if err != nil {
		return "", error_factory.NewInternalServerError("error generating token")
	}

	return newToken, nil
}

func (a *auth) VerifyToken() bool {

	return true
}
