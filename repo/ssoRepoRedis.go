package repo

import (
	data "EuprvaSsoService/data"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

type SecretRepoRedis struct {
	cli *redis.Client
}

const (
	secretKey = "secret"
	issuerKey = "issuer:%s"
)

func NewSecretRepoRedis(client *redis.Client) data.SecretRepo {
	return &SecretRepoRedis{
		cli: client,
	}
}

func (r SecretRepoRedis) GetSecret() (*data.Secret, error) {

	value := r.cli.Get(secretKey).Val()
	if value == "" {
		return nil, errors.New("request doesnt exists or has expired")
	}
	data := &data.Secret{}
	err := json.Unmarshal([]byte(value), data)
	if err != nil {
		return nil, errors.New("problem with reading data from DB")
	}
	return data, nil
}

func (r SecretRepoRedis) GetIssuer(issuer string) bool {
	value := r.cli.Get(fmt.Sprintf(issuerKey, issuer)).Val()
	if value == "" {
		panic("GRESKA PRI CITANJU")
		return false
	}
	return true
}

/*
Save
takes in secret value and ttl (in milliseconds) and stores it for that amount of time
*/
func (r SecretRepoRedis) Save(value data.Secret, ttl int) error {

	data, _ := json.Marshal(value)
	err := r.cli.Set(secretKey, data, time.Duration(ttl)*time.Millisecond).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r SecretRepoRedis) Remove(id string) error {
	err := r.cli.Del(secretKey).Err()
	if err != nil {
		return err
	}
	return nil
}
