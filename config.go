package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Path string
	Cmds []*Cmd
}

type Cmd struct {
	Name string `yaml:"name"`
	Exec string `yaml:"exec"`
}

var (
	ConfigFileName = ".roman.yml"
	RootDir        = string(os.PathSeparator) // for testing
)

func LoadConfig() (*Config, error) {
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

	cmds, err := ParseConfig(data)
	if err != nil {
		return nil, err
	}

	return &Config{
		Path: file,
		Cmds: cmds,
	}, nil
}

func ParseConfig(data []byte) ([]*Cmd, error) {
	var cmds []*Cmd
	if err := yaml.Unmarshal(data, &cmds); err != nil {
		return nil, err
	}

	return cmds, nil
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
