package main

import "strings"

type Board [][]byte

func Solve(pieces []Tetromino) Board {
	size := minSize(len(pieces))
	for {
		board := newBoard(size)
		used := make([]bool, len(pieces))

		slack := size*size - 4*len(pieces)
		if place(board, pieces, used, slack) {
			return board
		}
		size++
	}
}

func minSize(count int) int {
	side := 2
	for side*side < count*4 {
		side++
	}
	return side
}

func newBoard(size int) Board {
	board := make(Board, size)
	for i := range board {
		board[i] = []byte(strings.Repeat(".", size))
	}
	return board
}

func place(board Board, pieces []Tetromino, used []bool, slack int) bool {
	row, col, found := firstEmpty(board)
	if !found {
		return true
	}

	for i, piece := range pieces {
		if used[i] || sameShapeStillFree(pieces, used, i) {
			continue
		}

		startCol := col - piece[0].Col
		if !fits(board, piece, row, startCol) {
			continue
		}

		stamp(board, piece, row, startCol, byte('A'+i))
		used[i] = true
		if place(board, pieces, used, slack) {
			return true
		}
		used[i] = false
		stamp(board, piece, row, startCol, '.')
	}

	if slack > 0 {
		board[row][col] = ' '
		if place(board, pieces, used, slack-1) {
			return true
		}
		board[row][col] = '.'
	}
	return false
}

func sameShapeStillFree(pieces []Tetromino, used []bool, i int) bool {
	for j := 0; j < i; j++ {
		if !used[j] && sameShape(pieces[j], pieces[i]) {
			return true
		}
	}
	return false
}

func sameShape(a, b Tetromino) bool {
	for k := range a {
		if a[k] != b[k] {
			return false
		}
	}
	return true
}

func firstEmpty(board Board) (row, col int, found bool) {
	for r := range board {
		for c := range board[r] {
			if board[r][c] == '.' {
				return r, c, true
			}
		}
	}
	return 0, 0, false
}

func fits(board Board, piece Tetromino, row, col int) bool {
	size := len(board)
	for _, c := range piece {
		r, k := row+c.Row, col+c.Col
		if r < 0 || k < 0 || r >= size || k >= size || board[r][k] != '.' {
			return false
		}
	}
	return true
}

func stamp(board Board, piece Tetromino, row, col int, letter byte) {
	for _, c := range piece {
		board[row+c.Row][col+c.Col] = letter
	}
}

func Render(board Board) string {
	var sb strings.Builder
	for _, row := range board {
		sb.WriteString(strings.ReplaceAll(string(row), " ", "."))
		sb.WriteByte('\n')
	}
	return sb.String()
}
