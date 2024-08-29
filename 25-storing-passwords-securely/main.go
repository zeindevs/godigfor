package main

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func printUsage() {
	fmt.Println("Usage:", os.Args[0], " <password>")
	fmt.Println("Example:", os.Args[1], " Password!")
}

func checkArgs() string {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
	return os.Args[1]
}

var secretKey = "78c8ec932e9cde597fc14687c765bdd4b3f91b7b0617a3ed9f71bdf5d64e3d9b"

func generateSalt() string {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(randomBytes)
}

func hashPassword(plainText string, salt string) string {
	hash := hmac.New(sha256.New, []byte(secretKey))
	io.WriteString(hash, plainText+salt)
	hashValue := hash.Sum(nil)
	return hex.EncodeToString(hashValue)
}

func main() {
	password := checkArgs()
	salt := generateSalt()
	hashedPassword := hashPassword(password, salt)
	fmt.Println("Password:", password)
	fmt.Println("Salt:", salt)
	fmt.Println("Hashed Password:", hashedPassword)
}
