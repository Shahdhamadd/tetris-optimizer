# tetris-optimizer

Reads a text file full of tetrominoes and packs them into the smallest
square possible. Each piece is printed with a letter (`A` for the first,
`B` for the second, ...) and empty cells are printed as `.`.

## Usage

```bash
go run . examples/sample.txt
```

Output:

```
ABBBB.
ACCCEE
AFFCEE
A.FFGG
HHHDDG
.HDD.G
```

If the file is missing, badly formatted, or contains an invalid tetromino,
the program prints `ERROR`.

## Input format

- Each tetromino is a 4x4 block of `#` (filled) and `.` (empty).
- Blocks are separated by exactly one empty line.
- Every block must contain exactly 4 `#` cells that touch each other.
- At most 26 pieces (one per uppercase letter).

## How it works

1. `parser.go` reads the file, validates every block and trims each shape
   to its top-left corner.
2. `solver.go` starts with the smallest square that could hold all the
   cells (`ceil(sqrt(4 * pieces))`) and tries to place the pieces one by
   one using backtracking. If they do not fit, the square grows by one.
3. `main.go` prints the board.

## Tests

```bash
go test ./...
```

The `examples/` folder has good inputs, bad inputs, and a hard case used
by the tests.
