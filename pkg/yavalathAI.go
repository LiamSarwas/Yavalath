package pkg

import "math/rand"

func GetAIMove(g GameState) Hex {
  moves := g.GetAvailableMoves()
  return moves[rand.Intn(len(moves))]
}
