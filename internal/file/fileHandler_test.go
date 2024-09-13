package file_test

import (
	"compressionTool/internal/file"
	"testing"
	"testing/fstest"
)

func TestGetCharacterCounts(t *testing.T) {
	fs := fstest.MapFS{
		"test1.txt": {Data: []byte("xx")},
		"test2.txt": {Data: []byte("xxxx")},
	}

	counts, _ := file.GetCharacterCounts(fs, "test1.txt")

	count := counts['x']
	want := 2
	if count != want {
		t.Errorf("got %d; want %d", count, want)
	}
}
