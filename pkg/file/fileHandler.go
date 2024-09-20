package file

import (
	"fmt"
	"io"
	"io/fs"
	"strings"
	"unicode/utf8"
)

const chunkSize = 4096

func GetCharacterCounts(fileSystem fs.FS, path string) (map[rune]int, error) {
	file, err := fileSystem.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	charCounts := make(map[rune]int)

	buffer := make([]byte, chunkSize)
	var carry []byte

	for {
		n, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			return nil, fmt.Errorf("failed to read file: %w", err)
		}
		if err == io.EOF || n == 0 {
			break
		}

		chunk := append(carry, buffer[:n]...)
		carry = nil

		for i := 0; i < len(chunk); {
			r, size := utf8.DecodeRune(chunk[i:])
			if r == utf8.RuneError && size == 1 {
				carry = append(carry, chunk[i:]...)
				break
			}
			charCounts[r]++
			i += size
		}
	}

	return charCounts, nil
}

func WriteHeader(content map[rune]string) string {
	// header format:
	// <header> k:v k:v </header>
	var builder strings.Builder

	builder.WriteString("<compressionToolHeader>")

	first := true
	for k, v := range content {
		if !first {
			builder.WriteString("; ")
		}
		builder.WriteString(string(k))
		builder.WriteString(":")
		builder.WriteString(v)
		first = false
	}

	builder.WriteString("</compressionToolHeader>")

	return builder.String()
}

// Create file named according input
// Add header slice as first line
// Add decoded content from original file
