package service

import (
	"EuprvaSsoService/data"
	"crypto/rand"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"math/big"
	"time"
)

var JWTError = errors.New("error generating JWT")
var InvalidJwtClaims = errors.New("Invalid claims")
var InvalidSigningMethod = errors.New("Invalid signing method, expected HS512")

const oneHrInMs = 3600000

/*
SSOclaims is a struct of JWT claims used for authorization
in E_uprava App.
*/
type SSOclaims struct {
	jwt.RegisteredClaims
}

/*
GenerateJWT generates a json web token (JWT) signed with HS512 method.
Subject is JMBG of logged-in user.
returns the signed token if successful.
JWTError is returned if token couldn't be created
*/
func GenerateJWT(jmbg string, key string) (string, error) {
	claims := SSOclaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(oneHrInMs * time.Millisecond)),
			Subject:   jmbg,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return "", JWTError
	}
	return tokenString, nil
}

func GenerateSecretCode() (*data.Secret, error) {
	ret := make([]byte, secretLength)
	for i := 0; i < secretLength; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(secretLetters))))
		if err != nil {
			return nil, errors.New("problem while generating secret")
		}
		ret[i] = secretLetters[num.Int64()]
	}
	return &data.Secret{Secret: string(ret), ExpiresAt: time.Now().Add(oneHrInMs * time.Millisecond)}, nil
}