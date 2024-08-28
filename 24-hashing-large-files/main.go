package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
)

func printUsage() {
	fmt.Println("Usage:", os.Args[0], " <filename>")
	fmt.Println("Example:", os.Args[1], " diskimage.iso")
}

func checkArgs() string {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
	return os.Args[1]
}

func main() {
	filename := checkArgs()

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	hasher := md5.New()

	_, err = io.Copy(hasher, file)
	if err != nil {
		log.Fatal(err)
	}

	checksum := hasher.Sum(nil)

	fmt.Printf("MD5 checksums: %x\n", checksum)

}
