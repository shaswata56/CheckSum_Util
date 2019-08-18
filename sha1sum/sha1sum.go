package sha1sum

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		fmt.Println("md5sum: No such file or directory")
		os.Exit(0)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: C:\\Users\\Default> md5sum path/to/file.txt")
		return
	}
	filename := os.Args[1]
	file, err := os.Open(filename)
	check(err)
	md := md5.New()
	_, _ = io.Copy(md, file)
	fmt.Printf("%x	%s\n", md.Sum(nil), filepath.Base(filename))
	_ = file.Close()
}