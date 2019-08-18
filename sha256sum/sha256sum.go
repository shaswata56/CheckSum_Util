package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: C:\\Users\\Default> sha256sum path/to/file.txt")
		return
	}
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("sha256sum: No such file or directory")
		os.Exit(0)
	}
	hash := sha256.New()
	_, _ = io.Copy(hash, file)
	fmt.Printf("%x	%s\n", hash.Sum(nil), filepath.Base(filename))
	_ = file.Close()
}