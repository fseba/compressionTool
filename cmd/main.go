package main

import (
	"compressionTool/internal/file"
	utils "compressionTool/pkg"
	"fmt"
	"os"
)

// TODO: accept file name as input & return error if file is valid
// TODO: read file
// TODO: determine frequency of each character

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

	for char, count := range charCounts {
		fmt.Printf("%c: %d\n", char, count)
	}
}
