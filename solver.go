package main

import (
	"fmt"
	"strconv"
)

var board = [9][9]string{
	{"5", "3", ".", ".", "7", ".", ".", ".", "."},
	{"6", ".", ".", "1", "9", "5", ".", ".", "."},
	{".", "9", "8", ".", ".", ".", ".", "6", "."},
	{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
	{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
	{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
	{".", "6", ".", ".", ".", ".", "2", "8", "."},
	{".", ".", ".", "4", "1", "9", ".", ".", "5"},
	{".", ".", ".", ".", "8", ".", ".", "7", "9"},
}

func draw() {
	for i, row := range board {
		for j, char := range row {
			if (j+1)%3 == 0 && j+1 != len(row) {
				fmt.Printf("%s | ", char)
			} else {
				fmt.Printf("%s ", char)
			}
		}
		if (i+1)%3 == 0 && i+1 != len(row) {
			fmt.Printf("\n- - - - - - - - - - -\n")
		} else {
			fmt.Println()
		}
	}
	fmt.Println()
}

func alreadyInColumn(y int, value string) bool {
	for i := 0; i < 9; i++ {
		if board[i][y] == value {
			return true
		}
	}
	return false
}

func alreadyInLine(x int, value string) bool {
	for i := 0; i < 9; i++ {
		if board[x][i] == value {
			return true
		}
	}
	return false
}

func alreadyInSquare(x int, y int, value string) bool {
	square_x, square_y := int(x/3)*3, int(y/3)*3
	for browse_x := range [3]int{} {
		for browse_y := range [3]int{} {
			if board[square_x+browse_x][square_y+browse_y] == value {
				return true
			}
		}
	}
	return false
}

func canPlace(x int, y int, value string) bool {
	return !alreadyInColumn(y, value) && !alreadyInLine(x, value) && !alreadyInSquare(x, y, value)
}

func next(x int, y int) (int, int) {
	nextX, nextY := (x+1)%9, y
	if nextX == 0 {
		nextY = y + 1
	}
	return nextX, nextY
}

func solve(x int, y int) bool {
	if y == 9 {
		return true
	}

	if board[x][y] != "." {
		return solve(next(x, y))
	} else {
		for i := 0; i < 9; i++ {
			value := strconv.Itoa(i + 1)
			if canPlace(x, y, value) {
				board[x][y] = value
				if solve(next(x, y)) {
					return true
				} else {
					board[x][y] = "."
				}
			}
		}
		return false
	}
}

func main() {
	draw()

	if solve(0, 0) {
		fmt.Println("Solution Found :")
		draw()
	} else {
		fmt.Println("No solution")
	}
}
