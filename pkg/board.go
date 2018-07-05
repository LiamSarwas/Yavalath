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
	hexList        map[Hex]int
	availableMoves map[Hex]bool
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

func (g GameState) GetAvailableMoves() []Hex {
	moves := []Hex{}
	hexTest := Hex{0, 0}
	for i := -4; i <= 4; i++ {
		for j := -4; j <= 4; j++ {
			hexTest.x = i
			hexTest.y = j
			if g.availableMoves[hexTest] {
				moves = append(moves, hexTest)
			}
		}
	}
	return moves
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

func (g *GameState) SetUp() {
	g.hexList = make(map[Hex]int)
	g.availableMoves = make(map[Hex]bool)
	k := 0
	for j := -4; j <= 0; j++ {
		for i := k; i <= 4; i++ {
			g.hexList[Hex{i, j}] = 0
			g.availableMoves[Hex{i, j}] = true
		}
		k = k - 1
	}
	k = 3
	for j := 1; j <= 4; j++ {
		for i := -4; i <= k; i++ {
			g.hexList[Hex{i, j}] = 0
			g.availableMoves[Hex{i, j}] = true
		}
		k = k - 1
	}
}

// move validation is unnecessary because the engine will
// only pick from valid availableMoves
func (g *GameState) MakeMove(coord Hex, player int) (int, bool) {
	isGameOver := false
	gameStatus := 0
	// update board
	g.hexList[coord] = player
	g.availableMoves[coord] = false
	// check if game is won (1) or lost (-1) or draw (0) for current player
	// if game is over, gameOver = true,
	return gameStatus, isGameOver
}

func hexAdd(a, b Hex) Hex {
	return Hex{a.x + b.x, a.y + b.y}
}
