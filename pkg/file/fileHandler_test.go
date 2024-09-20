package file_test

import (
	"bufio"
	"compressionTool/pkg/file"
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

func TestHeaderCreator(t *testing.T) {
	fs := fstest.MapFS{
		"test1.txt": {Data: []byte("xx")},
	}

	file, err := fs.Open("test1.txt")
	if err != nil {
		t.Errorf("could not open file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	t.Log(scanner.Text())
}

func TestWriteHeader(t *testing.T) {
	m := map[rune]string{'A': "1", 'B': "10", 'C': "101"}

	got := file.WriteHeader(m)

	want := "<compressionToolHeader>A:1; B:10; C:101</compressionToolHeader>"

	if got != want {
		t.Errorf("got %q; want %q", got, want)
	}
}
