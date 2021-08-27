package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// backup old start script
	if err := os.Rename("./start", "./start.bak"); err != nil {
		log.Fatalf("Failed to backup start executable")
	}
	if err := copyRename("./stop", "start"); err != nil {
		log.Fatalf("Failed to copy stop to start: %v", err)
	}
	// remove backup script
	if err := os.Remove("./start.bak"); err != nil {
		log.Fatalf("Failed to remove backup executable")
	}
}

func copyRename(src, dst string) error {
	if err := os.Chmod(src, 0755); err != nil {
		return fmt.Errorf("failed to make file executable: %v", err)
	}
	return os.Rename(src, dst)
}
