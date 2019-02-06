package main

import (
	"crypto/rand"
	"encoding/base64"
	"flag"
	"fmt"
	"os"
)

func main() {
	var size int
	var help string

	flag.IntVar(&size, "size", 16, "")
	flag.StringVar(&help, "help", "off", "")

	for _, arg := range os.Args {
		if arg == "help" {
			menuhelp()
			os.Exit(0)
		}
	}

	if len(os.Args) < 2 {
		menuhelp()
		os.Exit(0)
	}

	flag.Parse()
	fmt.Println(seedGen(size))
}

func menuhelp() {
	fmt.Println("Usage: \n entropy [options]:")
	fmt.Println()
	fmt.Println("	-size 42	Create a seedword with 42 bytes. Default is 16.")
	fmt.Println()
}

func seedGen(size int) (result string) {
	crypto := make([]byte, size)
	_, err := rand.Read(crypto)
	if err != nil {
		return
	}
	result = base64.StdEncoding.EncodeToString(crypto)
	return result
}
