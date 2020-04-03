package main

import (
	"crypto/rand"
	"encoding/base64"
	"flag"
	"fmt"
	"log"
)

func main() {
	var size int
	var help string

	flag.IntVar(&size, "size", 16, "entropy -size 42 \nCreate a seedword with 42 bytes")
	flag.StringVar(&help, "help", "", "Display this help\n")

	flag.Parse()
	fmt.Println(seedGen(size))
}

func seedGen(size int) (result string) {
	crypto := make([]byte, size)

	_, err := rand.Read(crypto)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	return base64.StdEncoding.EncodeToString(crypto)
}
