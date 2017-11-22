//HR - Data Structures - 2D Array - Hour Glass

package main

import (
	"fmt"
	"strconv"
)

var input = [6][6]int{}

func main() {
	cmds, err := getInput()
	if err != nil {
		fmt.Println(err)
		return
	}

	i, j := 0, 0
	for _, c := range cmds {
		if j == 6 {
			j = 0
			i++
		}
		n, err := strconv.Atoi(c)
		if err != nil {
			fmt.Println(err)
			return
		}
		input[i][j] = n
		j++
	}

	maxSum := 0
	for i = 0; i < 6; i++ {
		for j = 0; j < 6; j++ {
			if isHourGlass(i, j) {
				sum := sumHourGlass(i, j)
				if sum > maxSum {
					maxSum = sum
				}
			}
		}
	}
	fmt.Println(maxSum)
}

func isHourGlass(i, j int) bool {
	if i+2 <= 5 && j+2 <= 5 {
		return true
	}
	return false
}

func sumHourGlass(i, j int) int {
	return input[i][j] + input[i][j+1] + input[i][j+2] + input[i+1][j+1] + input[i+2][j] + input[i+2][j+1] + input[i+2][j+2]
}

func getInput() (cmds []string, err error) {
	input := ""

	for {
		if _, err = fmt.Scanf("%s", &input); err != nil {
			if err.Error() == "EOF" {
				err = nil
				break
			} else {
				return nil, err
			}
		}
		cmds = append(cmds, input)
	}
	return
}
