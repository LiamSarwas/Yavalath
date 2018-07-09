package main

import (
	"fmt"
  "math/rand"
  "time"
)

const (
  SearchDuration = 1*time.Second
)

func main() {
  rand.Seed(time.Now().UnixNano())
  g := Game{}
  p1 := &mctsAI{}
  p2 := RandomAI{}

  g.Initialize(p1, p2)

  gameStatus := g.Play()
  if gameStatus == 1 {
    fmt.Println("Player 1 is victorious!")
  } else if gameStatus == 2 {
    fmt.Println("Player 2 is victorious!")
  } else {
    fmt.Println("It's a draw, you're both winners!")
  }
  p1.PlayerToString()
}
