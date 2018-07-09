package main

import (
  "time"
  "math/rand"
  "fmt"
)

type mctsAI struct{
  root *Node
}

func (m *mctsAI) Move(g GameState) Hex {
  if m.root == nil {
    m.root = &Node{}
    m.root.Initialize(g, Hex{0,0}, nil)
  }

  startTime := time.Now()
  routineCount := 0
  for loopDuration := time.Now().Sub(startTime); loopDuration < SearchDuration; {
    state := g.Clone()
    currNode := m.root
    gameResult := GameNotOver

    // Select
    for {
      if len(currNode.availableMoves) == 0 && len(currNode.children) > 0 {
        currNode = currNode.GetBestChild()
        gameResult = state.MakeMove(currNode.move)
      } else {
        break
      }
    }

    // Expand
    if moves := currNode.availableMoves; len(moves) > 0 {
      moveChoice := moves[rand.Intn(len(moves))]
      gameResult = state.MakeMove(moveChoice)
      currNode = currNode.AddChild(moveChoice, state)
    }

    // Rollout
    for {
      moves := state.GetAvailableMoves()
      if gameResult == GameNotOver && len(moves) > 0 {
        moveChoice := moves[rand.Intn(len(moves))]
        gameResult = state.MakeMove(moveChoice)
      } else {
        break
      }
    }

    // Backpropagate
    for {
      if currNode != nil {
        if gameResult == Player1Win && currNode.playerJustMoved {
          currNode.Update(1)
        } else if gameResult == Player1Win && !currNode.playerJustMoved {
          currNode.Update(0)
        } else if gameResult == Player2Win && currNode.playerJustMoved {
          currNode.Update(0)
        } else if gameResult == Player2Win && !currNode.playerJustMoved {
          currNode.Update(1)
        } else if gameResult == Draw {
          currNode.Update(0)
        }
        currNode = currNode.parent
      } else {
        break
      }
    }
    routineCount++
    loopDuration = time.Now().Sub(startTime)
  }

  mostVisitedChild := &Node{}
  mostVisits := 0
  for _, child := range m.root.children {
    if numVisits := child.visits; numVisits > mostVisits {
      mostVisitedChild = child
      mostVisits = numVisits
    }
  }
  m.root = mostVisitedChild
  return mostVisitedChild.move
}

func (m mctsAI) getOriginalRoot() *Node {
  currNode := m.root
  for {
    if currNode.parent != nil {
      currNode = currNode.parent
    } else {
      break
    }
  }
  return currNode
}

func (m mctsAI) PlayerToString() {
  fmt.Println(m.getOriginalRoot().NodeToString())
}
