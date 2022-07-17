package main

import (
	"flag"

	"github.com/BurntSushi/toml"
	"github.com/MaximumTroubles/go-rest-api-app/internal/app/apiserver"
	"github.com/sirupsen/logrus"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	//TODO: uncomment this part when needed and delete this comment
	// logrus.SetFormatter(&logrus.JSONFormatter{})

	flag.Parse()

	config := apiserver.NewConfig()

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		logrus.Fatal(err)
	}

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		logrus.Fatal(err)
	}

}
