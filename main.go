package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		fmt.Println("File does not exist!!")
	}
}

func main() {

	fmt.Println("HashGen v1.0\n")

	if len(os.Args) < 2 {
		fmt.Println("Error: File path is not specified in argument.")
		fmt.Println("\nExample:\n\thashgen ~/Downloads/ExampleFile.iso\n\thashgen passwords.txt")
		return
	}

	fileName := os.Args[1]

	file, err := os.Open(fileName)
	check(err)

	defer file.Close()

	md := md5.New()
	sha := sha1.New()
	sha2 := sha256.New()

	go io.Copy(md, file)
	go io.Copy(sha, file)
	go io.Copy(sha2, file)

	fmt.Println("Verify these Hashes with checksum.")
	fmt.Printf("md5:\t%x\n", md.Sum(nil))
	fmt.Printf("sha1:\t%x\n", sha.Sum(nil))
	fmt.Printf("sha256:\t%x\n", sha2.Sum(nil))
}
