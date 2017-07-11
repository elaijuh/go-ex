package main

import (
	"encoding/ascii85"
	"encoding/base64"
	"fmt"
)

func main() {

	// Demo for ascii85
	fmt.Println("--- Demo for ascii85")
	src := []byte("魑魅")
	buffer := make([]byte, ascii85.MaxEncodedLen(len(src)))
	encodedbytes := ascii85.Encode(buffer, src)

	fmt.Printf("source bytes length %v\n", len(src))
	fmt.Printf("encoded bytes length %v\n", encodedbytes)
	fmt.Printf("src %s is encoded to dest %s\n", src, buffer)

	ascii85.Decode(src, buffer, true)
	fmt.Printf("dst %s is decoded to src %s\n", buffer, src)

	// Demo for base64
	fmt.Println("--- Demo for base64")
	src64 := []byte("魑魅")
	encoded64bytes := base64.StdEncoding.EncodeToString(src64)
	fmt.Printf("src %s is encoded to dest %s\n", src64, encoded64bytes)

	src64, _ = base64.StdEncoding.DecodeString(encoded64bytes)
	fmt.Printf("dst %s is decoded to src %s\n", encoded64bytes, src64)
}
