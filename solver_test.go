package main

import (
	"os"
	"strings"
	"testing"
	"time"
)

var (
	shapeI = Tetromino{{0, 0}, {1, 0}, {2, 0}, {3, 0}}
	shapeO = Tetromino{{0, 0}, {0, 1}, {1, 0}, {1, 1}}
	shapeT = Tetromino{{0, 1}, {1, 0}, {1, 1}, {1, 2}}
	shapeS = Tetromino{{0, 1}, {0, 2}, {1, 0}, {1, 1}}
	shapeL = Tetromino{{0, 0}, {1, 0}, {2, 0}, {2, 1}}
)

func TestMinSize(t *testing.T) {
	cases := map[int]int{
		1: 2,
		2: 3,
		3: 4,
		4: 4,
		5: 5,
		8: 6,
	}
	for count, want := range cases {
		if got := minSize(count); got != want {
			t.Errorf("minSize(%d) = %d, want %d", count, got, want)
		}
	}
}

func TestSolveSample(t *testing.T) {
	input := strings.Join([]string{
		"...#\n...#\n...#\n...#",
		"....\n....\n....\n####",
		".###\n...#\n....\n....",
		"....\n..##\n.##.\n....",
		"....\n.##.\n.##.\n....",
		"....\n....\n##..\n.##.",
		"##..\n.#..\n.#..\n....",
		"....\n###.\n.#..\n....",
	}, "\n\n") + "\n"

	pieces, err := Parse(input)
	if err != nil {
		t.Fatalf("parse failed: %v", err)
	}

	got := Render(Solve(pieces))
	want := "ABBBB.\nACCCEE\nAFFCEE\nA.FFGG\nHHHDDG\n.HDD.G\n"
	if got != want {
		t.Errorf("wrong solution\ngot:\n%s\nwant:\n%s", got, want)
	}
}

func TestSolveSinglePiece(t *testing.T) {
	got := Render(Solve([]Tetromino{shapeO}))
	if got != "AA\nAA\n" {
		t.Errorf("single O piece: got\n%s", got)
	}
}

func TestSolvePerfectSquare(t *testing.T) {
	board := Solve([]Tetromino{shapeO, shapeO, shapeO, shapeO})

	if len(board) != 4 {
		t.Fatalf("expected 4x4 board, got %dx%d", len(board), len(board))
	}
	if dots := strings.Count(Render(board), "."); dots != 0 {
		t.Errorf("expected 0 empty cells, got %d", dots)
	}
}

func TestSolveNeedsBiggerSquare(t *testing.T) {
	board := Solve([]Tetromino{shapeS, shapeS, shapeS, shapeS})

	if len(board) != 5 {
		t.Fatalf("expected 5x5 board, got %dx%d", len(board), len(board))
	}
	if dots := strings.Count(Render(board), "."); dots != 9 {
		t.Errorf("expected 9 empty cells, got %d", dots)
	}
}

func TestSolveEveryPieceUsed(t *testing.T) {
	pieces := []Tetromino{shapeI, shapeL, shapeT, shapeO, shapeS}
	out := Render(Solve(pieces))

	for i := range pieces {
		letter := string(rune('A' + i))
		if n := strings.Count(out, letter); n != 4 {
			t.Errorf("piece %s appears %d times, want 4", letter, n)
		}
	}
	if strings.Contains(out, "F") {
		t.Error("unexpected extra piece F in output")
	}
}

func TestSolveExampleFiles(t *testing.T) {
	cases := map[string]int{
		"examples/good00.txt": 0,
		"examples/good01.txt": 9,
		"examples/good02.txt": 4,
		"examples/good03.txt": 5,
		"examples/hard.txt":   1,
	}

	for file, wantDots := range cases {
		content, err := os.ReadFile(file)
		if err != nil {
			t.Fatalf("cannot read %s: %v", file, err)
		}
		pieces, err := Parse(string(content))
		if err != nil {
			t.Fatalf("%s: parse failed: %v", file, err)
		}

		start := time.Now()
		out := Render(Solve(pieces))
		elapsed := time.Since(start)

		if dots := strings.Count(out, "."); dots != wantDots {
			t.Errorf("%s: expected %d empty cells, got %d\n%s", file, wantDots, dots, out)
		}
		if elapsed > 2*time.Second {
			t.Errorf("%s: solver too slow: %v", file, elapsed)
		}
	}
}

func TestSolveHardIsFast(t *testing.T) {
	pieces := []Tetromino{
		shapeI, shapeI, shapeO, shapeT,
		{{0, 0}, {0, 1}, {1, 1}, {1, 2}},
		shapeL, shapeT,
		{{0, 0}, {0, 1}, {0, 2}, {0, 3}},
		shapeS,
		{{0, 0}, {1, 0}, {1, 1}, {2, 0}},
		{{0, 2}, {1, 0}, {1, 1}, {1, 2}},
		{{0, 1}, {1, 1}, {2, 0}, {2, 1}},
	}

	start := time.Now()
	board := Solve(pieces)
	elapsed := time.Since(start)

	if len(board) != 7 {
		t.Errorf("expected 7x7 board, got %dx%d", len(board), len(board))
	}
	if dots := strings.Count(Render(board), "."); dots != 1 {
		t.Errorf("expected 1 empty cell, got %d", dots)
	}
	if elapsed > 2*time.Second {
		t.Errorf("solver too slow: %v", elapsed)
	}
}
