package config

import (
	"log"
	"os"
	"path"

	"github.com/caarlos0/env/v10"
	"github.com/goccy/go-yaml"
	"github.com/pkg/errors"
)

type Configuration struct {
	Env Env
	App Application
}

type Env struct {
	ConfigFolder string `env:"GOCARDS_CONFIG_FOLDER" envDefault:"~/.gocards"`
}

type Application struct {
	Language string `yaml:"language"`
}

func LoadFromFile() (*Configuration, error) {
	envConf := Env{}
	err := env.Parse(&envConf)
	if err != nil {
		return nil, err
	}

	yamlPath := path.Join(envConf.ConfigFolder, "config.yml")
	yamlFile, err := os.ReadFile(yamlPath)
	if err != nil {
		return nil, errors.Wrap(err, "could not read config file")
	}

	res := Application{}
	err = yaml.Unmarshal(yamlFile, &res)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return &Configuration{
		Env: envConf,
		App: res,
	}, nil
}

// TODO add a SaveToFile function
