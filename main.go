package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

func main() {
	h := md5.New()
	if _, err := io.Copy(h, os.Stdin); err != nil {
		fmt.Println("error copy data to hash")
		os.Exit(1)
	}
	fmt.Printf("%x", h.Sum(nil))
}
