package main

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"net/http"
	"os"
)

func main() {
	// Generate secret path with 144 bits of entropy
	var entropy [18]byte
	log.Println("Generating secret path...")
	if _, err := rand.Read(entropy[:]); err != nil {
		log.Fatalln("Could not generate path:", err.Error())
	}
	path := "/" + base64.URLEncoding.EncodeToString(entropy[:]) + "/"
	log.Println("Done.")
	log.Println("Open your browser at: " + os.Getenv("APP_URL") + path)
	log.Println("---")
	log.Println("\n")
	
	// Serve filesystem at secret path
	http.Handle(path, http.StripPrefix(path, http.FileServer(http.Dir("/"))))
	log.Fatalln(http.ListenAndServe("localhost:8080", nil))
}
