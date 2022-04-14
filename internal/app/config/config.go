package config

import (
	"os"

	"github.com/pkg/errors"
	"go.uber.org/multierr"
	"gopkg.in/yaml.v2"
)

type Config struct {
	App 	App 	`yaml:"app"`
	DB  	DB 		`yaml:"db"`
	API		API 	`yaml:"api"`
	Process Process `yaml:"process"`
}

type App struct {
	Port string `yaml:"port"`
}

type DB struct {
	Conn string `yaml:"conn"`
}

type API struct {
	Key string `yaml:"key"`
}

type Process struct {
	Ticker int `yaml:"ticker"`
}

func New(path string) (config Config, err error) {
	file, err := os.Open(path)
	if err != nil {
		return config, errors.Wrapf(err, "open config by path %s", path)
	}
	defer func(err error) {
		multierr.AppendInto(&err, file.Close())
	}(err)

	d := yaml.NewDecoder(file)

	err = d.Decode(&config)
	if err != nil {
		return config, errors.Wrap(err, "decode config information")
	}

	return config, nil
}