package bs

import (
	"errors"
	"fmt"
	"sort"
)

type Position struct {
	R int
	C int
}

type Line struct {
	From      Position
	Direction Position
	Len       int
}

type Lines []Line

func (l Lines) Len() int {
	return len(l)
}

func (l Lines) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

type LinesByLen struct {
	Lines
}

func (s LinesByLen) Less(i, j int) bool {
	return s.Lines[i].Len < s.Lines[j].Len
}

func valid(board [][]int) bool {
	for _, row := range board {
		if len(row) != len(board) {
			return false
		}
	}
	return true
}

func cloneBoard(board [][]int) (cloned [][]int) {
	cloned = make([][]int, len(board))
	for i, row := range board {
		cloned[i] = make([]int, len(row))
		copy(cloned[i], row)
	}
	return
}

func sign(n int) (result int) {
	switch {
	case n > 0:
		result = 1
	case n == 0:
		result = 0
	case n < 0:
		result = -1
	}
	return
}

func makeLine(s, e Position) (line *Line, err error) {
	if !validLine(s, e) {
		return nil, errors.New("Not a valid line.")
	}
	line = &Line{
		From:      s,
		Direction: Position{sign(e.R - s.R), sign(e.C - s.C)},
	}
	if line.Direction.R != 0 {
		line.Len = (e.R-s.R)/line.Direction.R + 1
	} else {
		line.Len = (e.C-s.C)/line.Direction.C + 1
	}
	return
}

func validLine(s, e Position) bool {
	return s.R == e.R || s.C == e.C || (s.R-e.R) == (s.C-e.C) || (s.R-e.R) == -(s.C-e.C)
}

func solve(board [][]int) (solved bool, solution []Line) {
	size := len(board)
	used := make([][]bool, size)
	for i := range used {
		used[i] = make([]bool, size)
	}
	usableCount := size * size
	restore := func(positions []Position) {
		for _, p := range positions {
			used[p.R][p.C] = false
		}
	}

	solution = make([]Line, 0)

	useLine := func(line Line) (result []Position, err error) {
		cur := line.From
		to := Position{line.From.R + (line.Len-1)*line.Direction.R, line.From.C + (line.Len-1)*line.Direction.C}
		if used[line.From.R][line.From.C] || used[to.R][to.C] {
			err = errors.New("blabla")
			return
		}
		for i := 0; i < line.Len; i += 1 {
			if !used[cur.R][cur.C] {
				if board[cur.R][cur.C] != board[line.From.R][line.From.C] {
					restore(result)
					result = nil
					err = errors.New("blabla")
					return
				}
				used[cur.R][cur.C] = true
				result = append(result, cur)
			}
			cur.R += line.Direction.R
			cur.C += line.Direction.C
		}
		return
	}
	availableLines := func() (result []Line) {
		result = make([]Line, 0)
		lb := len(board)
		num := len(board) * len(board)
		for n1 := 0; n1 < num; n1 += 1 {
			for n2 := 0; n2 < n1; n2 += 1 {
				p1 := Position{n1 / lb, n1 % lb}
				p2 := Position{n2 / lb, n2 % lb}
				if board[p1.R][p1.C] == board[p2.R][p2.C] {
					if line, err := makeLine(p1, p2); err == nil && line.Len > 1 {
						result = append(result, *line)
					}
				}
			}
		}
		sort.Sort(sort.Reverse(LinesByLen{result}))
		return
	}

	var solver func()
	solver = func() {
		for _, line := range availableLines() {
			setList, err := useLine(line)
			if err != nil {
				continue
			}
			fmt.Printf("Using line (%d, %d), direction (%d, %d) len %d\n", line.From.R, line.From.C, line.Direction.R, line.Direction.C, line.Len)
			solution = append(solution, line)
			usableCount -= len(setList)
			if usableCount == 0 {
				fmt.Printf("Solved!!!\n")
				fmt.Printf("S size %d\n", len(solution))
				solved = true
				return
			}
			solver()
			if solved {
				fmt.Printf("Solved...\n")
				return
			}
			solution = solution[:len(solution)-1]
			fmt.Printf("Not use line (%d, %d), direction (%d, %d) len %d\n", line.From.R, line.From.C, line.Direction.R, line.Direction.C, line.Len)
			usableCount += len(setList)
			restore(setList)
		}
		return
	}
	solver()
	fmt.Printf("S size %d\n", len(solution))
	return
}

func Solve(board [][]int) (solution []Line, err error) {
	if !valid(board) {
		err = errors.New("Board is not a square.")
		return
	}
	fmt.Println("Solving...")
	_, solution = solve(cloneBoard(board))
	return solution, nil
}
