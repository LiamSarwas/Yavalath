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
		game.ToString()
    fmt.Println(game.GetAvailableMoves())
	}
}
