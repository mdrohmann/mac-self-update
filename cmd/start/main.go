package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) >= 2 && os.Args[1] == "installer" {
		fmt.Println("Running installer")
		cmd := exec.Command("./installer")
		cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
		if err := cmd.Run(); err != nil {
			log.Fatalf("Failed to run installer command: %v", err)
		}
	} else {
		if err := os.Remove("./start"); err != nil {
			log.Fatalf("Failed to remove start: %v", err)
		}
		if err := copy("./stop", "start"); err != nil {
			log.Fatalf("Failed to copy stop to start: %v", err)
		}
	}

	fmt.Println("Re-running ./start command")
	cmd := exec.Command("./start")
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to run updated command: %v", err)
	}
}

func copy(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}

	_, err = io.Copy(out, in)
	if err != nil {
		out.Close()
		return err
	}

	if err := out.Close(); err != nil {
		return err
	}

	return os.Chmod(dst, 0755)
}
