package main

import (
	"math/rand"
)

type RandomAI struct{}

func (r RandomAI) Move(g GameState) Hex {
	moves := g.GetAvailableMoves()
	return moves[rand.Intn(len(moves))]
}
