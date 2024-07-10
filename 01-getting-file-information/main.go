package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	fileInfo, err := os.Stat("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File name:", fileInfo.Name())
	fmt.Println("File in bytes:", fileInfo.Size())
	fmt.Println("Permissions:", fileInfo.Mode())
	fmt.Println("Last modified:", fileInfo.ModTime())
	fmt.Println("Is Directory:", fileInfo.IsDir())
	fmt.Println("System interface type:", fileInfo.Sys())
	json, err := json.MarshalIndent(fileInfo.Sys(), " ", " ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("System info: %s\n", json)
}
