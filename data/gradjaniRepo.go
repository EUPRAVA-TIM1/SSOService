package data

type GradjaniRepo interface {
	GetByJmbg(jmbg string) (*Gradjanin, error)
}
