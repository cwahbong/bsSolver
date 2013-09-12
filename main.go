package main

import (
	"fmt"
	"github.com/cwahbong/bsSolver/bs"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	board := make([][]int, n)
	for i := 0; i < n; i += 1 {
		board[i] = make([]int, n)
		for j := 0; j < n; j += 1 {
			fmt.Scanf("%d", &board[i][j])
		}
	}
	solution, err := bs.Solve(board)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Solution size: %d\n", len(solution))
	for _, line := range solution {
		fmt.Printf("from (%d, %d), direction (%d, %d), len %d.\n", line.From.R, line.From.C, line.Direction.R, line.Direction.C, line.Len)
	}
}
