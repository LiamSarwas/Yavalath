package main

const (
	Player1Win  = 1
	Player2Win  = 2
	Draw        = 0
	GameNotOver = -1
)

type Game struct {
  board GameState
  p1 Player
  p2 Player
}

func (game *Game) Initialize(p1, p2 Player) {
  game.board = GameState{}
  game.board.Initialize()
  game.p1 = p1
  game.p2 = p2
}

func (game *Game) Play() int {
  // start the game and loop infinitely, a win/loss will break the loop
  for {
    move := Hex{}
    if game.board.currentPlayer {
      move = game.p1.Move(game.board)
    } else {
      move = game.p2.Move(game.board)
    }
    gameResult := game.board.MakeMove(move)
    if gameResult != GameNotOver {
      game.board.ToString()
      return gameResult
    }
    game.board.ToString()
	}
}
