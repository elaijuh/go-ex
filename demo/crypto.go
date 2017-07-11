package main

import (
	"crypto/aes"
	"crypto/md5"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	// Demo for md5
	fmt.Println("--- Demo for md5")
	data := []byte("tech sing song")
	fmt.Printf("%x\n", md5.Sum(data))

	// Demo for bcrypt
	fmt.Println("--- Demo for bcrypt")
	password := []byte("iamPassword")
	hashedPassword, _ := bcrypt.GenerateFromPassword(password, 4)
	fmt.Printf("bcrypt hashed password %s\n", hashedPassword)

	err := bcrypt.CompareHashAndPassword(hashedPassword, []byte("123"))
	// err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	if err != nil {
		fmt.Println(err)
	}

	// Demo for AES
	fmt.Println("--- Demo for AES")
	block, err := aes.NewCipher([]byte("iamakey12345678"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("The block size is %d\n", block.BlockSize())

	var dst = make([]byte, 16)
	var src = []byte("iamaAESdemo12345")

	block.Encrypt(dst, src)
	fmt.Printf("AES encrypted: %x\n", dst)

	// Demo for RSA

}
