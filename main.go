package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"os"
)

var flagFile = flag.String("file", "", "file used for input")

var input io.Reader

func main() {
	flag.Parse()
	switch {
	case *flagFile != "":
		f, err := os.Open(*flagFile)
		if err != nil {
			fmt.Println("error opening file:", *flagFile, err)
			os.Exit(123)
		}
		defer f.Close()
		input = f
	default:
		input = os.Stdin
	}
	printMD5(input, os.Stdout)

}

func printMD5(r io.Reader, w io.Writer) {
	h := md5.New()
	if _, err := io.Copy(h, r); err != nil {
		fmt.Println("error copy data to hash")
		os.Exit(1)
	}

	fmt.Fprintf(w, "%x", h.Sum(nil))
}
