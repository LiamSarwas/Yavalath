package pkg

import (
	"fmt"
)

type Hex struct {
	x int
	y int
}

type Chain struct {
	len int
	val int
}

type GameState struct {
	hexList   map[Hex]int
	chainList map[Hex][6]Chain
}

func (g GameState) hexGridToStringSlice() [9]string {
	hexStrings := [9]string{}
	k := 0
	for j := -4; j <= 0; j++ {
		hexString := ""
		for i := k; i <= 4; i++ {
			hexString = hexString + fmt.Sprintf(" %d", g.hexList[Hex{i, j}])
		}
		hexStrings[j+4] = hexString
		k = k - 1
	}
	k = 3
	for j := 1; j <= 4; j++ {
		hexString := ""
		for i := -4; i <= k; i++ {
			hexString = hexString + fmt.Sprintf(" %d", g.hexList[Hex{i, j}])
		}
		hexStrings[j+4] = hexString
		k = k - 1
	}
	return hexStrings
}

func (g GameState) ToString() {
	hexStrings := g.hexGridToStringSlice()
	for i := -4; i <= 4; i++ {
		k := 0
		if i >= 0 {
			k = i
		} else {
			k = -i
		}
		for j := 0; j < k; j++ {
			fmt.Print(" ")
		}
		fmt.Print(hexStrings[i+4])
		for j := 0; j < k; j++ {
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func (g *GameState) SetUp() {
	g.hexList = make(map[Hex]int)
	for i := -4; i <= 4; i++ {
		for j := -4; j <= 4; j++ {
			g.hexList[Hex{i, j}] = 0
		}
	}
}

func (g *GameState) MakeMove(coord Hex) {
	// update board
	g.hexList[coord] = 1
	// update chainList
}
