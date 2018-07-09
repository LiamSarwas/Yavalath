package main

import (
	"fmt"
)

var axialDirections = []Hex{Hex{-1, 0}, Hex{0, -1}, Hex{+1, -1},
	Hex{+1, 0}, Hex{0, +1}, Hex{-1, +1}}

type Hex struct {
	x int
	y int
}

type GameState struct {
	hexList           map[Hex]int
	availableMoves    map[Hex]bool
	numAvailableMoves int
	currentPlayer     bool
}

func hexAdd(a, b Hex) Hex {
	return Hex{a.x + b.x, a.y + b.y}
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
	fmt.Println()
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

func (g GameState) getChainAlongAxis(coord Hex, axis int) int {
	maxChainLength := 1
	// look along one of the three axes 0,1,2 to see how many adjacent
	// tiles have been played by the current player
	currHex := coord
	chainNotBroken := true
	for chainNotBroken {
		if nextNeighbor, isValidNeighbor := GetNeighbor(currHex, axis); isValidNeighbor {
			if g.hexList[nextNeighbor] == 1 && !g.currentPlayer {
				maxChainLength++
				currHex = nextNeighbor
			} else if g.hexList[nextNeighbor] == 2 && g.currentPlayer {
				maxChainLength++
				currHex = nextNeighbor
			} else {
				chainNotBroken = false
			}
		} else {
			chainNotBroken = false
		}
	}
	currHex = coord
	chainNotBroken = true
	for chainNotBroken {
		if nextNeighbor, isValidNeighbor := GetNeighbor(currHex, axis+3); isValidNeighbor {
			if g.hexList[nextNeighbor] == 1 && !g.currentPlayer {
				maxChainLength++
				currHex = nextNeighbor
			} else if g.hexList[nextNeighbor] == 2 && g.currentPlayer {
				maxChainLength++
				currHex = nextNeighbor
			} else {
				chainNotBroken = false
			}
		} else {
			chainNotBroken = false
		}
	}
	return maxChainLength
}

func (g GameState) getMaxChain(coord Hex) int {
	maxChain := 0
	for i := 0; i <= 2; i++ {
		if axisChain := g.getChainAlongAxis(coord, i); axisChain > maxChain {
			maxChain = axisChain
		}
	}
	return maxChain
}

func (g *GameState) Initialize() {
	g.hexList = make(map[Hex]int)
	g.availableMoves = make(map[Hex]bool)
	k := 0
	for j := -4; j <= 0; j++ {
		for i := k; i <= 4; i++ {
			g.hexList[Hex{i, j}] = 0
			g.availableMoves[Hex{i, j}] = true
			g.numAvailableMoves++
		}
		k = k - 1
	}
	k = 3
	for j := 1; j <= 4; j++ {
		for i := -4; i <= k; i++ {
			g.hexList[Hex{i, j}] = 0
			g.availableMoves[Hex{i, j}] = true
			g.numAvailableMoves++
		}
		k = k - 1
	}
}

// move validation is unnecessary because the engine will
// only pick from valid availableMoves
func (g *GameState) MakeMove(coord Hex) int {
	// update board
	if !g.currentPlayer {
		g.hexList[coord] = 1
	} else {
		g.hexList[coord] = 2
	}
	g.availableMoves[coord] = false
	g.numAvailableMoves--

	// check win loss draw conditions
	maxChain := g.getMaxChain(coord)
	if maxChain >= 4 && !g.currentPlayer {
		return Player1Win
	} else if maxChain >= 4 && g.currentPlayer {
		return Player2Win
	} else if maxChain == 3 && !g.currentPlayer {
		return Player2Win
	} else if maxChain == 3 && g.currentPlayer {
		return Player1Win
	}
	if g.numAvailableMoves == 0 {
		return Draw
	}

	// swap current player
	g.currentPlayer = !g.currentPlayer
	return GameNotOver
}

func (g *GameState) AvailableMovesCopy() map[Hex]bool {
	newAvailableMoves := make(map[Hex]bool)
	for k, v := range g.availableMoves {
		newAvailableMoves[k] = v
	}
	return newAvailableMoves
}

func (g *GameState) Clone() GameState {
	newBoard := GameState{}
	newBoard.Initialize()
	newBoard.currentPlayer = g.currentPlayer
	newBoard.hexList = make(map[Hex]int)

	for k, v := range g.hexList {
		newBoard.hexList[k] = v
	}
	for k, v := range g.availableMoves {
		newBoard.availableMoves[k] = v
	}
	return newBoard
}
