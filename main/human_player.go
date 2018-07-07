package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type HumanPlayer struct{}

func Move(g GameState) Hex {
	fmt.Print("Input a pair of coordinates: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading standard input:", err)
	}

	read_line := strings.TrimSuffix(input, "\n")
	input_strings := strings.Split(read_line, ",")
	ints := []int{0, 0}
	for i, val := range input_strings {
		if s, err := strconv.Atoi(val); err == nil {
			ints[i] = s
		} else {
			fmt.Fprintln(os.Stderr, "Error converting to int:", err)
		}
	}
	return Hex{ints[0], ints[1]}
}
