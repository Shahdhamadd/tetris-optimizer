package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestParseGood(t *testing.T) {
	input := "...#\n...#\n...#\n...#\n\n....\n....\n..##\n..##\n"

	pieces, err := Parse(input)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(pieces) != 2 {
		t.Fatalf("expected 2 pieces, got %d", len(pieces))
	}

	wantI := Tetromino{{0, 0}, {1, 0}, {2, 0}, {3, 0}}
	if !samePiece(pieces[0], wantI) {
		t.Errorf("I piece not normalized: got %v", pieces[0])
	}

	wantO := Tetromino{{0, 0}, {0, 1}, {1, 0}, {1, 1}}
	if !samePiece(pieces[1], wantO) {
		t.Errorf("O piece not normalized: got %v", pieces[1])
	}
}

func TestParseWindowsLineEndings(t *testing.T) {
	input := "####\r\n....\r\n....\r\n....\r\n"
	if _, err := Parse(input); err != nil {
		t.Errorf("expected CRLF input to parse, got %v", err)
	}
}

func TestParseNoTrailingNewline(t *testing.T) {
	input := "####\n....\n....\n...."
	if _, err := Parse(input); err != nil {
		t.Errorf("expected input without final newline to parse, got %v", err)
	}
}

func TestParseTrailingBlankLines(t *testing.T) {
	input := "####\n....\n....\n....\n\n\n"
	pieces, err := Parse(input)
	if err != nil {
		t.Fatalf("expected trailing blank lines to be ignored, got %v", err)
	}
	if len(pieces) != 1 {
		t.Errorf("expected 1 piece, got %d", len(pieces))
	}
}

func TestParseBad(t *testing.T) {
	cases := map[string]string{
		"empty file":          "",
		"only newlines":       "\n\n",
		"five hashes":         "##..\n##..\n#...\n....\n",
		"three hashes":        "###.\n....\n....\n....\n",
		"disconnected":        "#..#\n....\n#..#\n....\n",
		"diagonal touch":      "#...\n.#..\n..#.\n...#\n",
		"wrong character":     "####\n....\n...x\n....\n",
		"line too long":       "####.\n....\n....\n....\n",
		"line too short":      "###\n....\n....\n....\n",
		"too many lines":      "####\n....\n....\n....\n....\n",
		"too few lines":       "####\n....\n....\n",
		"two blank lines":     "####\n....\n....\n....\n\n\n####\n....\n....\n....\n",
		"no blank line":       "####\n....\n....\n....\n####\n....\n....\n....\n",
		"spaces instead":      "#   \n#   \n#   \n#   \n",
		"lowercase letters":   "aaaa\n....\n....\n....\n",
		"more than 26 pieces": repeatPiece(27),
	}

	for name, input := range cases {
		if _, err := Parse(input); err == nil {
			t.Errorf("%s: expected an error, got none", name)
		}
	}
}

func TestParseExampleFiles(t *testing.T) {
	files, err := filepath.Glob("examples/*.txt")
	if err != nil || len(files) == 0 {
		t.Fatalf("could not find example files: %v", err)
	}

	for _, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			t.Fatalf("cannot read %s: %v", file, err)
		}

		_, err = Parse(string(content))
		isBad := strings.HasPrefix(filepath.Base(file), "bad")
		if isBad && err == nil {
			t.Errorf("%s: expected ERROR, got a valid parse", file)
		}
		if !isBad && err != nil {
			t.Errorf("%s: expected valid parse, got %v", file, err)
		}
	}
}

func TestIsConnected(t *testing.T) {
	connected := Tetromino{{0, 0}, {0, 1}, {1, 1}, {1, 2}}
	if !isConnected(connected) {
		t.Error("S shape should be connected")
	}

	square := Tetromino{{0, 0}, {0, 1}, {1, 0}, {1, 1}}
	if !isConnected(square) {
		t.Error("O shape should be connected")
	}

	broken := Tetromino{{0, 0}, {0, 1}, {2, 0}, {2, 1}}
	if isConnected(broken) {
		t.Error("split shape should not be connected")
	}
}

func samePiece(a, b Tetromino) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func repeatPiece(n int) string {
	block := "####\n....\n....\n....\n"
	blocks := make([]string, n)
	for i := range blocks {
		blocks[i] = block
	}
	return strings.Join(blocks, "\n")
}
