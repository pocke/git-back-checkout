package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	err := Main(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func Main(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("Arguments is missing")
	}
	br := args[0]
	back := "HEAD~"
	if len(args) >= 2 {
		back = args[1]
	}

	if err := exec.Command("git", "branch", br).Run(); err != nil {
		return err
	}
	if err := exec.Command("git", "reset", "--hard", back).Run(); err != nil {
		return err
	}
	if err := exec.Command("git", "checkout", br).Run(); err != nil {
		return err
	}

	return nil
}
