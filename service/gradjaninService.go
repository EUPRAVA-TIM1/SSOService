package service

import "EuprvaSsoService/data"

type GradjaniService interface {
	GetByJMBG(jmbg string) (*data.Gradjanin, error)
}

type gradjaniService struct {
	repo data.GradjaniRepo
}

func NewGradjaniService(repo data.GradjaniRepo) GradjaniService {
	return gradjaniService{repo: repo}
}

func (g gradjaniService) GetByJMBG(jmbg string) (*data.Gradjanin, error) {
	return g.repo.GetByJmbg(jmbg)
}
