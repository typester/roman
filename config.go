package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Name string `yaml:"name"`
	Exec string `yaml:"exec"`
}

var (
	ConfigFileName = ".roman.yml"
	RootDir        = string(os.PathSeparator) // for testing
)

func LoadConfig() ([]*Config, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	file, err := SearchConfig(cwd)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return ParseConfig(data)
}

func ParseConfig(data []byte) ([]*Config, error) {
	var configs []*Config
	if err := yaml.Unmarshal(data, &configs); err != nil {
		return nil, err
	}

	return configs, nil
}

func SearchConfig(dir string) (string, error) {
	cur := dir
	for {
		stat, err := os.Stat(filepath.Join(cur, ConfigFileName))
		if err != nil {
			if os.IsNotExist(err) {
				next := filepath.Clean(filepath.Join(cur, ".."))
				if next == cur || cur == RootDir {
					return "", os.ErrNotExist
				}
				cur = next
				continue
			}
			return "", err
		}

		return filepath.Join(cur, stat.Name()), nil
	}
}
