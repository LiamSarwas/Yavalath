package main

// a player can make a move (specified by a Hex coordinate)
type Player interface {
	Move(g GameState) Hex
	GetOppMove(oppMove Hex)
}
