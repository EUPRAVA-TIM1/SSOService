package service

import (
	"EuprvaSsoService/data"
	"errors"
	"fmt"
)

type GradjaniService interface {
	Whoami(token string) (*data.Gradjanin, error)
	GetByJMBG(jmbg string) (*data.Gradjanin, error)
}

type gradjaniService struct {
	repo       data.GradjaniRepo
	secretRepo data.SecretRepo
}

func NewGradjaniService(repo data.GradjaniRepo, secretRepo data.SecretRepo) GradjaniService {
	return gradjaniService{repo: repo, secretRepo: secretRepo}
}

/* Whoami returns user from db that has JMBG that is provided as JWT subject*/
func (g gradjaniService) Whoami(token string) (*data.Gradjanin, error) {
	secret, err := g.secretRepo.GetSecret()
	if err != nil {
		return nil, errors.New("Problem while reading secret from db")
	}
	err = ValidateJwt(token, secret.Secret)
	if err == nil {
		jmbg, err := GetPrincipal(token, secret.Secret)
		if err != nil {
			fmt.Println(err.Error())
			return nil, errors.New("Problem while geting pricipals from token")
		}
		return g.repo.GetByJmbg(jmbg)
	}
	return nil, errors.New("Invalid JWT token")
}

func (g gradjaniService) GetByJMBG(jmbg string) (*data.Gradjanin, error) {
	return g.repo.GetByJmbg(jmbg)
}
