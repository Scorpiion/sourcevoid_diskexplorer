package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Generate secret path with 144 bits of entropy
	var entropy [18]byte
	fmt.Println("Generating secret path...")
	if _, err := rand.Read(entropy[:]); err != nil {
		log.Fatalln("Could not generate path:", err.Error())
	}
	path := "/" + base64.URLEncoding.EncodeToString(entropy[:]) + "/"
	fmt.Printf("The path is: xxx.sourcevoid.net%s", path)
	fmt.Println("")
	fmt.Println("\n")
	
	// Serve filesystem at secret path
	http.Handle(path, http.StripPrefix(path, http.FileServer(http.Dir("/"))))
	log.Fatalln(http.ListenAndServe("localhost:8080", nil))
}
