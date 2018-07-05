package pkg

import (
	"fmt"
)

var axialDirections = []Hex{Hex{-1, 0}, Hex{0, -1}, Hex{+1, -1},
	Hex{+1, 0}, Hex{0, +1}, Hex{-1, +1}}

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
	chainList map[Hex][]Chain
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

// d is the direction (0 - 5) of the six cardinal hex axialDirections
// 0 is directly to the left and proceeds clockwise
func GetNeighbor(coord Hex, d int) (Hex, bool) {
	neighbor := hexAdd(coord, axialDirections[d])
	isValidHex := true
	if neighbor.x+neighbor.y > 4 || neighbor.x+neighbor.y < -4 {
		isValidHex = false
	}
	return neighbor, isValidHex
}

func (g *GameState) MakeMove(coord Hex) {
	// update board
	g.hexList[coord] = 1
	// update chainList
}

func hexAdd(a, b Hex) Hex {
	return Hex{a.x + b.x, a.y + b.y}
}
