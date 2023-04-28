package main

import (
	"EuprvaSsoService/config"
	"EuprvaSsoService/startup"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(&config)
	server.Start()
}
