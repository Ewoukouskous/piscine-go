package main

import "fmt"

func main() {
	var board [3][3]string
	var a int
	var b int
	run := true

	printBoard(board)

	for run {
		fmt.Println("Enter x: ")
		fmt.Scan(&a)

		fmt.Println("Enter y: ")
		fmt.Scan(&b)

		if a == 9 || b == 9 {
			run = false
		} else if board[a][b] == "0" {
			fmt.Println("La case est déjà occupee")
			run = false
		}

		board[a][b] = "0"

		printBoard(board)
	}
}

func printBoard(board [3][3]string) {
	for i := 0; i < len(board); i++ {
		fmt.Println("-------")
		for j := 0; j < len(board[i]); j++ {
			fmt.Print("| " + string(board[i][j]))
		}
		fmt.Println("|")
	}
	fmt.Println("-------")
}
