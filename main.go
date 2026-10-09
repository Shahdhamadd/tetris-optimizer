package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("ERROR")
		return
	}

	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("ERROR")
		return
	}

	pieces, err := Parse(string(content))
	if err != nil {
		fmt.Println("ERROR")
		return
	}

	board := Solve(pieces)
	fmt.Print(Render(board))
	fmt.Println()
}
