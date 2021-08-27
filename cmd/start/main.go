package main

import (
	"fmt"
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
		if err := copyRename("./stop", "start"); err != nil {
			log.Fatalf("Failed to copy stop to start: %v", err)
		}
	}

	fmt.Println("Re-running ./start command")
	cmd := exec.Command("./start")
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	if err := cmd.Start(); err != nil {
		log.Fatalf("Failed to start updated command: %v", err)
	}

	if err := cmd.Wait(); err != nil {
		log.Fatalf("Failed to wait for updated command to finish: %v", err)
	}
}

func copyRename(src, dst string) error {
	if err := os.Chmod(src, 0755); err != nil {
		return fmt.Errorf("failed to make file executable: %v", err)
	}
	return os.Rename(src, dst)
}
