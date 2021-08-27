package main

import (
	"fmt"
	"io"
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
	if err := copyFile(src, dst); err != nil {
		return fmt.Errorf("failed to copy file to destination: %v", err)
	}
	if err := os.Chmod(dst, 0755); err != nil {
		return fmt.Errorf("failed to make file executable: %v", err)
	}
	return os.Rename(src, dst)
}

func copyFile(src, target string) error {
	in, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("os.Open %s failed: %v", src, err)
	}
	defer in.Close()

	// Create target file
	out, err := os.Create(target)
	if err != nil {
		return fmt.Errorf("os.Create %s failed: %v", target, err)
	}
	defer out.Close()

	// Copy bytes to target file
	_, err = io.Copy(out, in)
	if err != nil {
		return fmt.Errorf("io.Copy failed: %v", err)
	}
	err = out.Close()
	if err != nil {
		return fmt.Errorf("out.Close failed: %v", err)
	}
	return nil
}
