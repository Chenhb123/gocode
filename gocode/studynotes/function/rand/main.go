package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	fmt.Println(key())
}

func key() string {
	// buf := make([]byte, 16)
	buf := []byte("证书字符串")
	_, err := rand.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	// return fmt.Sprintf("%x", buf)
	return hex.EncodeToString(buf)
}
