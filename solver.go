package main

import "fmt"

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

func main() {
	draw()
	fmt.Println(alreadyInColumn(4, "8"))
	fmt.Println(alreadyInLine(4, "1"))
}
