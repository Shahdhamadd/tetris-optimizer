package main

import (
	"errors"
	"strings"
)

type Cell struct {
	Row, Col int
}

type Tetromino []Cell

var ErrBadFormat = errors.New("bad format")

const MaxPieces = 26

func Parse(content string) ([]Tetromino, error) {
	content = strings.ReplaceAll(content, "\r\n", "\n")

	content = strings.TrimRight(content, "\n")
	if content == "" {
		return nil, ErrBadFormat
	}

	blocks := strings.Split(content, "\n\n")
	if len(blocks) > MaxPieces {
		return nil, ErrBadFormat
	}

	pieces := make([]Tetromino, 0, len(blocks))
	for _, block := range blocks {
		piece, err := parseBlock(block)
		if err != nil {
			return nil, err
		}
		pieces = append(pieces, piece)
	}
	return pieces, nil
}

func parseBlock(block string) (Tetromino, error) {
	lines := strings.Split(block, "\n")
	if len(lines) != 4 {
		return nil, ErrBadFormat
	}

	var cells Tetromino
	for row, line := range lines {
		if len(line) != 4 {
			return nil, ErrBadFormat
		}
		for col, ch := range line {
			switch ch {
			case '#':
				cells = append(cells, Cell{row, col})
			case '.':

			default:
				return nil, ErrBadFormat
			}
		}
	}

	if len(cells) != 4 || !isConnected(cells) {
		return nil, ErrBadFormat
	}
	return normalize(cells), nil
}

func isConnected(cells Tetromino) bool {
	touches := 0
	for _, a := range cells {
		for _, b := range cells {
			sameRow := a.Row == b.Row && abs(a.Col-b.Col) == 1
			sameCol := a.Col == b.Col && abs(a.Row-b.Row) == 1
			if sameRow || sameCol {
				touches++
			}
		}
	}
	return touches == 6 || touches == 8
}

func normalize(cells Tetromino) Tetromino {
	minRow, minCol := cells[0].Row, cells[0].Col
	for _, c := range cells {
		if c.Row < minRow {
			minRow = c.Row
		}
		if c.Col < minCol {
			minCol = c.Col
		}
	}

	shifted := make(Tetromino, len(cells))
	for i, c := range cells {
		shifted[i] = Cell{c.Row - minRow, c.Col - minCol}
	}
	return shifted
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
