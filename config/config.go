package config

import (
	"os"
)

type Config struct {
	Port      string
	Host      string
	RedisPort string
	RedisHost string
}

const (
	redisHost        = "SSO_REDIS_HOST"
	redisPort        = "SSO_REDIS_PORT"
	defaultRedisHost = "sso_redis"
	defaultRedisPort = "6379"
	portKey          = "SSO_SERVICE_PORT"
	hostKey          = "SSO_SERVICE_HOST"
	defaultHost      = "sso"
	defaultPort      = "8000"
)

func NewConfig() (c Config) {
	if port, set := os.LookupEnv(portKey); set && port != "" {
		c.Port = port
	} else {
		c.Port = defaultPort
	}

	if host, set := os.LookupEnv(hostKey); set && host != "" {
		c.Host = host
	} else {
		c.Host = defaultHost
	}

	if host, set := os.LookupEnv(redisHost); set && host != "" {
		c.RedisHost = host
	} else {
		c.RedisHost = defaultRedisHost
	}

	if port, set := os.LookupEnv(redisPort); set && port != "" {
		c.RedisPort = port
	} else {
		c.RedisPort = defaultRedisPort
	}
	return
}
