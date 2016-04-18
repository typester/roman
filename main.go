package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	configs, err := LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load %s: %s\n", ConfigFileName, err.Error())
	}

	if len(os.Args) < 2 {
		log.Fatalf("Required command argument to run\n")
	}

	cmd := os.Args[1]

	for _, c := range configs {
		if c.Name == cmd {
			if err := ExecCommand(c, os.Args[2:]); err != nil {
				if err, ok := err.(*exec.ExitError); ok {
					if st, ok := err.ProcessState.Sys().(syscall.WaitStatus); ok {
						os.Exit(st.ExitStatus())
					}
				}
				log.Fatalf("command exec failed: %s\n", err.Error())
			}
			return
		}
	}

	log.Fatalf("no config found for command '%s'\n", cmd)
}
