package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"os"
)

func printUsage() {
	fmt.Printf(os.Args[0] + `

Encrypt or decrypt a file using AES with a 256-bit key file.
This program can also generate 256-bit keys.

Usage:
  ` + os.Args[0] + ` [-h|--help]
  ` + os.Args[0] + ` [-g|--genkey]
  ` + os.Args[0] + ` <keyFile> <file> [-d|--decrypt]

Examples:
  # Generate a 32-byte (256-bit) key
  ` + os.Args[0] + ` --genkey

  # Encrypt with secret key. Output to STDOUT
  ` + os.Args[0] + ` --genkey > secret.key

  # Encrypt message using secret key. Output to ciphertext.dat
  ` + os.Args[0] + ` secret.key message.txt > ciphertext.dat

  # Decrypt message using secret key. Output to STDOUT
  ` + os.Args[0] + ` secret.key ciphertext.dat -d

  # Decrypt message using secret key. Output to message.txt
  ` + os.Args[0] + ` secret.key ciphertext.dat -d > cleartext.txt
`)
}

func checkArgs() (string, string, bool) {
	if len(os.Args) < 2 || len(os.Args) > 4 {
		printUsage()
		os.Exit(1)
	}

	if len(os.Args) == 2 {
		if os.Args[1] == "-h" || os.Args[1] == "--help" {
			printUsage()
			os.Exit(1)
		}
		if os.Args[1] == "-g" || os.Args[1] == "--genkey" {
			key := generateKey()
			fmt.Printf(string(key[:]))
			os.Exit(0)
		}
	}

	if len(os.Args) == 3 {
		return os.Args[1], os.Args[2], false
	}

	if len(os.Args) == 4 {
		if os.Args[3] != "-d" && os.Args[3] != "--decrypt" {
			fmt.Println("Error: Unknown usage.")
			printUsage()
			os.Exit(1)
		}
		return os.Args[1], os.Args[2], true
	}
	return "", "", false
}

func generateKey() []byte {
	randomBytes := make([]byte, 32)
	numBytesRead, err := rand.Read(randomBytes)
	if err != nil {
		log.Fatal("Error generating random key.", err.Error())
	}
	if numBytesRead != 32 {
		log.Fatal("Error generating 32 random bytes for key.")
	}
	return randomBytes
}

func encrypt(key, message []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	cipherText := make([]byte, aes.BlockSize+len(message))

	iv := cipherText[:aes.BlockSize]
	_, err = io.ReadFull(rand.Reader, iv)
	if err != nil {
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(cipherText[aes.BlockSize:], message)

	return cipherText, nil
}

func decrypt(key, cipherText []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(cipherText, cipherText)

	return cipherText, nil
}

func main() {
	keyFile, file, decryptFlag := checkArgs()

	keyFileData, err := os.ReadFile(keyFile)
	if err != nil {
		log.Fatal("Unable to read key fil contents.", err.Error())
	}

	fileData, err := os.ReadFile(file)
	if err != nil {
		log.Fatal("Unable to read key file contents.", err.Error())
	}

	if decryptFlag {
		message, err := decrypt(keyFileData, fileData)
		if err != nil {
			log.Fatal("Error decrypting.", err.Error())
		}
		fmt.Printf("%s", message)
	} else {
		cipherText, err := encrypt(keyFileData, fileData)
		if err != nil {
			log.Fatal("Error encrypting.", err.Error())
		}
		fmt.Printf("%s", cipherText)
	}
}
