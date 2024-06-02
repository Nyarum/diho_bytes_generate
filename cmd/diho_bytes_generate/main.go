package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Nyarum/diho_bytes_generate/generate"
	"github.com/Nyarum/diho_bytes_generate/parse"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <filename>", os.Args[0])
	}
	fmt.Println(os.Args)
	filename := os.Args[1]

	pkgName, packetDescrs := parse.ParseBinaryFile(filename)
	generate.GenerateEncodeForStruct(filename, pkgName, packetDescrs)
	generate.GenerateDecodeForStruct(filename, pkgName, packetDescrs)

	fmt.Println(packetDescrs)
}
