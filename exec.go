package main

import (
	"os"
	"os/exec"

	"github.com/mattn/go-shellwords"
)

func ExecCommand(conf *Cmd, args []string) error {
	p := shellwords.NewParser()
	p.ParseEnv = true

	shell, err := p.Parse(conf.Exec)
	if err != nil {
		return err
	}

	shell = append(shell, args...)

	cmd := exec.Command(shell[0], shell[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
