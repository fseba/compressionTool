package main

import (
	"compressionTool/pkg/file"
	tb "compressionTool/pkg/treeBuilder"
	utils "compressionTool/pkg/utils"
	"fmt"
	"os"
)

func main() {
	utils.PrintBanner()
	if len(os.Args) < 2 {
		fmt.Println("Please provide at least one argument.")
		return
	}

	// os.Args[0] is the program name
	fmt.Println("Program Name:", os.Args[0])

	// os.Args[1:] are the arguments
	for i, arg := range os.Args[1:] {
		fmt.Printf("Argument %d: %s\n", i+1, arg)
	}

	fs := os.DirFS(".")

	fileName := os.Args[1]

	charCounts, err := file.GetCharacterCounts(fs, fileName)
	if err != nil {
		fmt.Printf("Error getting char count: %v", err)
		return
	}

	node := tb.BuildHuffmanTree(charCounts)
	codes := tb.GenerateHuffmanCodes(node)

	// for rune, code := range codes {
	// 	fmt.Printf("%c: %q\n", rune, code)
	// }
	fmt.Println(file.WriteHeader(codes))
}
