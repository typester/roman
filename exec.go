package main

import (
	"os"
	"os/exec"

	"github.com/mattn/go-shellwords"
)

func ExecCommand(conf *Config, args []string) error {
	p := shellwords.NewParser()
	p.ParseEnv = true

	cmds, err := p.Parse(conf.Exec)
	if err != nil {
		return err
	}

	cmds = append(cmds, args...)

	cmd := exec.Command(cmds[0], cmds[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
