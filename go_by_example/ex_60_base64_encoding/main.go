package main

import (
	b64 "encoding/base64"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Base64 Encoding")

	data := "abc123|?$*&()'-=@~"

	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	sDec, err := b64.StdEncoding.DecodeString(sEnc)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(sDec))
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, err := b64.URLEncoding.DecodeString(uEnc)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(uDec))
}
