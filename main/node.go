package main

import (
  "math"
  "strconv"
)

type Node struct {
  move Hex
  parent *Node
  children []*Node
  wins int
  visits int
  availableMoves []Hex
  playerJustMoved bool
}


func (n *Node) Initialize(g GameState, move Hex, parent *Node) {
    n.move = move
    n.parent = parent
    n.children = []*Node{}
    n.wins = 0
    n.visits = 0
    n.availableMoves = g.GetAvailableMoves()
    n.playerJustMoved = !g.currentPlayer
}

func (n *Node) GetBestChild() *Node {
  bestChild := &Node{}
  bestScore := 0.0
  for _, child := range n.children {
    if childScore := UCTScore(child); childScore > bestScore {
      bestChild = child
      bestScore = childScore
    }
  }
  return bestChild
}


func (n *Node) AddChild(move Hex, g GameState) *Node {
  newNode := Node{}
  newNode.Initialize(g, move, n)
  n.removeMove(move)
  n.children = append(n.children, &newNode)
  return &newNode
}

func (n *Node) Update(result int) {
  n.visits = n.visits + 1
  n.wins = n.wins + result
}

func (n *Node) removeMove(move Hex) {
  for i, coord := range n.availableMoves {
    if coord == move {
      n.availableMoves = append(n.availableMoves[:i], n.availableMoves[i+1:]...)
    }
  }
}

func UCTScore(n *Node) float64 {
  UCTScore := 0.0
  if n.parent != nil {
    UCTScore = float64(n.wins)/float64(n.visits) + ExploreFactor*math.Sqrt(math.Log(float64(n.parent.visits))/float64(n.visits))
  }
  return UCTScore
}

func (n Node) TreeToString(indent int) string {
  treeOut := n.indentString(indent) + n.NodeToString()
  for _, child := range n.children {
    treeOut = treeOut + child.TreeToString(indent + 1)
  }
  return treeOut
}

func (n Node) NodeToString() string {
  return "[M:" + strconv.Itoa(n.move.x) + "," + strconv.Itoa(n.move.y) + " W/V:" + strconv.Itoa(n.wins) + "/" + strconv.Itoa(n.visits) + " UCT:" + strconv.FormatFloat(UCTScore(&n), 'f', -1, 64) + "]"
}

func (n Node) indentString(indent int) string {
  in := "\n"
  for i := 0; i < indent; i++ {
    in = in + "| "
  }
  return in
}
