package main

import (
	"flag"
	"fmt"
	"os"
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "for decode or encode jwt token\n")
		fmt.Fprintf(os.Stderr, "Usage : \n")
		fmt.Fprintf(os.Stderr, "cat file | parjwt\n")
		fmt.Fprintf(os.Stderr, "echo <jwt token> | parjwt\n")
		fmt.Fprintf(os.Stderr, "cat file | parjwt e      >> for encode\n")
		fmt.Fprintf(os.Stderr, "cat file | parjwt d      >> for decode\n")

	}
}

type Color string

const (
	Black  Color = "\u001b[1;30m"
	Red          = "\u001b[1;31m"
	Green        = "\u001b[1;32m"
	Yellow       = "\u001b[1;33m"
	Blue         = "\u001b[1;34m"
	Cyan         = "\u001b[1;36m"
	Reset        = "\u001b[0m"
)

func colorize(color Color, message string) {
	fmt.Println(string(color), message)
}
func main() {
	flag.Parse()
	a := flag.Arg(0)
	if a == "" {
		flag.Usage()
	} else if a == "e" {
		encode()
	} else if a == "d" {
		decode()
	}
}
