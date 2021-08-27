package main

import (
	"io"
	"log"
	"os"
)

func main() {
	if err := os.Remove("./start"); err != nil {
		log.Fatalf("Failed to remove start: %v", err)
	}
	if err := copy("./stop", "start"); err != nil {
		log.Fatalf("Failed to copy stop to start: %v", err)
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
