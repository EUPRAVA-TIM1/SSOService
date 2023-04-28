package service

import (
	"EuprvaSsoService/data"
	"errors"
	"fmt"
	"time"
)

const secretLetters = "7*Q%!)S+X6Iw3FoLOAN8h@W#Gt2$Rb^m9cdzMiJ(s_n1jT4VC$YguxHZEqD&kf0aKpBrlPvyeU5"
const secretLength = 64

type SSOservice interface {
	Login(credentials data.Credentials) (*data.JWTReturn, error)
	GetSecret(issuer string) (*data.Secret, error)
}

type ssoService struct {
	secretRepo data.SecretRepo
	secret     data.Secret
}

func NewSSOService(repo data.SecretRepo) SSOservice {
	secret, err := repo.GetSecret()
	if err != nil && secret != nil {
		panic(err)
	}
	if secret == nil {
		secret, err = GenerateSecretCode()
		if err != nil {
			panic(err)
		}
	}
	service := ssoService{repo, *secret}
	err = service.secretRepo.Save(service.secret, oneHrInMs)
	if err != nil {
		panic(err)
	}
	time.AfterFunc(time.Until(secret.ExpiresAt), service.UpdateSecret)
	return service
}

func (s ssoService) GetSecret(issuer string) (*data.Secret, error) {
	if s.secretRepo.GetIssuer(issuer) {
		secret, err := s.secretRepo.GetSecret()
		if err != nil {
			return nil, errors.New("there was problem with reading secret key from db (key maybe expired)")
		}
		return secret, nil
	}
	return nil, errors.New("unknown issuer")
}

func (s ssoService) Login(credentials data.Credentials) (*data.JWTReturn, error) {
	//TODO after check in users db
	jwt, err := GenerateJWT(credentials.JMBG, s.secret.Secret)
	if err != nil {
		return nil, err
	}
	return &data.JWTReturn{Token: jwt}, nil
}

/*
UpdateSecret generates  secret using GenerateSecretCode and stores it in redis db for 1hr after which it calls itself (declared in time.AfterFunc) and generates new one
*/
func (s ssoService) UpdateSecret() {
	//TODO only test delete this line
	fmt.Sprint("The secret has changed TESST")
	secret, err := GenerateSecretCode()
	err = s.secretRepo.Save(*secret, oneHrInMs)
	if err != nil {
		panic(err)
	}
	s.secret = *secret
	time.AfterFunc(time.Until(secret.ExpiresAt), s.UpdateSecret)
}
