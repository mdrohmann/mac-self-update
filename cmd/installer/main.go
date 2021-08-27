package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if err := copyRename("./stop", "start"); err != nil {
		log.Fatalf("Failed to copy stop to start: %v", err)
	}
}

func copyRename(src, dst string) error {
	if err := os.Chmod(src, 0755); err != nil {
		return fmt.Errorf("failed to make file executable: %v", err)
	}
	return os.Rename(src, dst)
}
