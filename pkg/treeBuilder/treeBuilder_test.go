package treebuilder_test

import (
	tb "compressionTool/pkg/treeBuilder"
	"testing"
)

var testData = map[rune]int{
	'C': 32,
	'D': 42,
	'E': 120,
	'K': 7,
	'L': 42,
	'M': 24,
	'U': 37,
	'Z': 2,
}

func TestBuildHuffmanTree(t *testing.T) {
	root := tb.BuildHuffmanTree(testData)

	if root == nil {
		t.Fatalf("Expected a valid tree but got nil")
	}

	expectedTotal := 0
	for _, frequency := range testData {
		expectedTotal += frequency
	}
	if root.Frequency != expectedTotal {
		t.Errorf("Expected root frequency to be %d, but got %d", expectedTotal, root.Frequency)
	}
}

func TestHuffmanEncoding(t *testing.T) {
	root := tb.BuildHuffmanTree(testData)
	encoding := tb.GenerateHuffmanCodes(root)

	if len(encoding) != len(testData) {
		t.Errorf("Expected %d encoding entries but got %d", len(testData), len(encoding))
	}
}
