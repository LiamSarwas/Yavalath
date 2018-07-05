package main

import (
	"fmt"
  "github.com/LiamSarwas/Yavalath/pkg"
)

func main() {
	game := pkg.GameState{}
	game.SetUp()
	game.ToString()

	// start the game and loop infinitely, a win/loss will break the loop
  i := 0
  isGameOver := false
  gameStatus := 0
  for {
    if i%2 == 0 {
      fmt.Print("Input a pair of coordinates: ")
      input := pkg.GetHumanMove()
      fmt.Println(input)
      gameStatus, isGameOver = game.MakeMove(input, 1)
      game.ToString()
    } else {
      gameStatus, isGameOver = game.MakeMove(pkg.GetAIMove(game), 2)
    }
    if isGameOver {
      if gameStatus == 1 and i%2 == 0 {
        fmt.Println("Player 1 is victorious!")
      }
      if gameStatus == 1 and i%2 == 1 {
        fmt.Println("Player 2 is victorious!")
      }
      if gameStatus == -1 and i%2 == 0 {
        fmt.Println("Player 2 is victorious!")
      }
      if gameStatus == -1 and i%2 == 1 {
        fmt.Println("Player 1 is victorious!")
      }
      if gameStatus == 0 {
        fmt.Println("It's a draw, you're both winners!")
      }
    }
    i += 1
	}
}
