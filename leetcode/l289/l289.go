// LeetCode L289
// it is a simple convolution, conceptually very easy if you know what convolution is.
// use a sliding window.
// However, it is simple to say, also simple to implement if you can use libraries like NumPy.
// I now understand better why NumPy is so popular...

package main

import "fmt"

func main() {
	b1 := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0},
	}

	gameOfLife(b1)
	printBoard(b1)
}

func gameOfLife(board [][]int) {

	// padding of 1
	oldBoard := make([][]int, len(board)+2)
	for i := range oldBoard {
		oldBoard[i] = make([]int, len(board[0])+2)
		for j := range oldBoard[i] {
			oldBoard[i][j] = 2
		}
	}

	// copy
	for j := range board {
		copy(oldBoard[1 : len(oldBoard)-1][j][1:], board[j])
	}

	// convolution
	for i := 1; i < len(oldBoard)-1; i++ {
		for j := 1; j < len(oldBoard[0])-1; j++ {
			// count living neighbors
			nCnt := 0

			// copy. be careful dealing with slices
			slidingWindow := make([][]int, 3)
			copy(slidingWindow, oldBoard[i-1:i+2])
			for k := range slidingWindow {
				slidingWindow[k] = slidingWindow[k][j-1 : j+2]
			}

			// convolution
			for m := range slidingWindow {
				for n := range slidingWindow[0] {
					if (m != 1 || n != 1) && slidingWindow[m][n] == 1 {
						nCnt++
					}
				}
			}
			if nCnt < 2 {
				board[i-1][j-1] = 0
			} else if nCnt == 3 {
				board[i-1][j-1] = 1
			} else if nCnt > 3 {
				board[i-1][j-1] = 0
			}
		}
	}

}

func printBoard(b [][]int) {
	for i := range b {
		fmt.Println(b[i])
	}
}
