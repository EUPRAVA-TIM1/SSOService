package data

type SecretRepo interface {
	GetSecret() (*Secret, error)
	GetIssuer(issuer string) bool
	Save(value Secret, ttl int) error
	Remove(id string) error
}
