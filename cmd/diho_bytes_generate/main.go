package main

import (
	"bytes_generated/generate"
	"bytes_generated/parse"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("Usage: %s <filename>", os.Args[0])
	}
	filename := os.Args[1]
	packageName := os.Args[2]

	packetDescr := parse.ParseBinaryFile(filename)
	generate.GenerateEncodeForStruct(filename, packageName, packetDescr)

	fmt.Println(packetDescr)
}
