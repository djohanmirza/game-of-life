package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

//rules:
//	- Any live cell with fewer than two live neighbours dies, as if by underpopulation.
//	- Any live cell with two or three live neighbours lives on to the next generation.
//	- Any live cell with more than three live neighbours dies, as if by overpopulation.
//	- Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.

func updateState(state *[][]int) [][]int {
	updateState := make([][]int, len(*state))
	for i, row := range *state {
		updateState[i] = make([]int, len(row))
	}

	for i, row := range *state {
		for j := range row {
			count := getCount(state, i, j)

			if count < 2 {
				updateState[i][j] = 0
			} else if count == 2 {
				updateState[i][j] = (*state)[i][j]
			} else if count == 3 {
				updateState[i][j] = 1
			} else {
				updateState[i][j] = 0
			}
		}
	}
	return updateState
}

func getCount(state *[][]int, i int, j int) int {
	count := 0

	dirs := [][]int{
		{1, -1}, {1, 0}, {1, 1},
		{0, -1}, {0, 1},
		{-1, -1}, {-1, 0}, {-1, 1},
	}

	n := len(*state)
	m := len((*state)[0])
	for _, dir := range dirs {
		x := (n + i - dir[0]) % n
		y := (m + j - dir[1]) % m

		count += (*state)[x][y]
	}

	return count
}

func stateConvertor(state *[][]int) string {
	output := ""
	for _, row := range *state {
		for _, val := range row {
			if val == 1 {
				output += "# "
			} else {
				output += "_ "
			}
		}
		output += "\n"
	}

	return output
}

func main() {
	state := make([][]int, 5)
	for i := range state {
		state[i] = make([]int, 5)
	}

	state[0][0] = 1
	state[0][2] = 1
	state[1][1] = 1
	state[2][2] = 1

	//state[0][1] = 1
	//state[1][0] = 1
	//state[2][1] = 1

	for {
		clear := exec.Command("clear")
		clear.Stdout = os.Stdout
		clear.Run()

		fmt.Println(stateConvertor(&state))
		state = updateState(&state)

		time.Sleep(1 * time.Second)
	}

}
