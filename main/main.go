package main

import (
	"fmt"

)

func main() {
  g := Game{}
  p1 := RandomAI{}
  p2 := RandomAI{}
  
  g.Initialize(p1, p2)

  gameStatus := g.Play()
  if gameStatus == 1 {
    fmt.Println("Player 1 is victorious!")
  } else if gameStatus == -1 {
    fmt.Println("Player 2 is victorious!")
  } else {
    fmt.Println("It's a draw, you're both winners!")
  }
}
