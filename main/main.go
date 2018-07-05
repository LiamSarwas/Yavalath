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
	for {
		fmt.Print("Input a pair of coordinates: ")
		input := pkg.GetHumanMove()
		fmt.Println(input)
		game.MakeMove(input)
    for i := 0; i < 6; i++ {
      if neighbor, ok := pkg.GetNeighbor(input, i); ok {
        game.MakeMove(neighbor)
      }
    }
		game.ToString()
	}
}
