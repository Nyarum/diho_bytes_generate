package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Nyarum/diho_bytes_generate/generate"
	"github.com/Nyarum/diho_bytes_generate/parse"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("Usage: %s <filename>", os.Args[0])
	}
	filename := os.Args[1]
	packageName := os.Args[2]

	packetDescr := parse.ParseBinaryFile(filename)
	generate.GenerateEncodeForStruct(filename, packageName, packetDescr)
	generate.GenerateDecodeForStruct(filename, packageName, packetDescr)

	fmt.Println(packetDescr)
}
