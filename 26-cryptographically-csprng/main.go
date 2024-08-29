package main

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

func main() {
	limit := int64(math.MaxInt64)
	randInt, err := rand.Int(rand.Reader, big.NewInt(limit))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Random int value:", randInt)

	var number uint32
	err = binary.Read(rand.Reader, binary.BigEndian, &number)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Random uint32 value:", number)

	numBytes := 4
	randomBytes := make([]byte, numBytes)
	rand.Read(randomBytes)
	fmt.Println("Random byte values:", randomBytes)
}
