package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		fmt.Println("File does not exist!!")
	}
}

func sha_1(filename string) {
	file, err := os.Open(filename)
	check(err)
	sha := sha1.New()
	_, _ = io.Copy(sha, file)
	fmt.Printf("sha1:\t%x\n", sha.Sum(nil))
	_ = file.Close()
}

func sha2(filename string) {
	file, err := os.Open(filename)
	check(err)
	sha := sha256.New()
	_, _ = io.Copy(sha, file)
	fmt.Printf("sha256:\t%x\n", sha.Sum(nil))
	_ = file.Close()
}

func sha5(filename string) {
	file, err := os.Open(filename)
	check(err)
	sha := sha512.New()
	_, _ = io.Copy(sha, file)
	fmt.Printf("sha512:\t%x\n", sha.Sum(nil))
	_ = file.Close()
}

func main() {

	fmt.Println("HashGen v1.0\n")

	if len(os.Args) < 2 {
		fmt.Println("Error: File path is not specified in argument.")
		fmt.Println("\nExample:\n\thashgen ~/Downloads/ExampleFile.iso sha1\n\thashgen passwords.txt sha512")
		return
	}
	if len(os.Args) < 3 {
		fmt.Println("Error: Hash type is not specified in argument.")
		fmt.Println("\nExample:\n\thashgen ~/Downloads/ExampleFile.iso sha1\n\thashgen passwords.txt sha512")
		return
	}

	fileName := os.Args[1]
	typ := os.Args[2]

	if typ == "sha1" {
		sha_1(fileName)
	}
	if typ == "sha256" {
		sha2(fileName)
	}
	if typ == "sha512" {
		sha5(fileName)
	}
}
