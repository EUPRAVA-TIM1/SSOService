package config

import (
	"os"
)

type Config struct {
	Port          string
	Host          string
	RedisPort     string
	RedisHost     string
	MysqlPort     string
	MySqlHost     string
	MySqlRootPass string
}

const (
	redisHost            = "SSO_REDIS_HOST"
	redisPort            = "SSO_REDIS_PORT"
	defaultRedisHost     = "sso_redis"
	defaultRedisPort     = "6379"
	portKey              = "SSO_SERVICE_PORT"
	hostKey              = "SSO_SERVICE_HOST"
	defaultHost          = "sso"
	defaultPort          = "8000"
	mySqlPort            = "SSO_MYSQL_PORT"
	defaultMySqlPort     = "3306"
	mySqlHost            = "SSO_MYSQL_HOST"
	defaultMySqlHost     = "sso_mysql"
	mySqlRootPass        = "SSO_MYSQL_ROOT_PASSWORD"
	defaultMySqlRootPass = "root"
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

	if port, set := os.LookupEnv(mySqlPort); set && port != "" {
		c.MysqlPort = port
	} else {
		c.MysqlPort = defaultMySqlPort
	}

	if host, set := os.LookupEnv(mySqlHost); set && host != "" {
		c.MySqlHost = host
	} else {
		c.MySqlHost = defaultMySqlHost
	}

	if pass, set := os.LookupEnv(mySqlRootPass); set && pass != "" {
		c.MySqlRootPass = pass
	} else {
		c.MySqlRootPass = defaultMySqlRootPass
	}
	return
}
