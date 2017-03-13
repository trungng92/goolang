// package game is the middle layer of this code.
// Overall the code is separated into three layers.
// The lowest level layer does simple functionality; all the basic actions you can do on the board.
// This package is the middle layer. It represents what actions the player can do (and who can play).
// The highest level is what the player can see. It mostly consists of any in game ui and audio.
package game

import (
	"github.com/trungng92/goolang/board"
	"github.com/trungng92/goolang/util"
)

type Game struct {
	gameBoard         board.Board
	tempGameBoard     board.Board // acts as a buffer so we can see what the board looks like without modifying the real board
	players           []board.Player
	currentPlayerTurn int          // the index of players array. Determines whose turn it is
	TurnRules         []RuleFilter // filters that are checked on each turn
	StartRules        []RuleFilter // filters that happen at the beginning of the game
}

func (g Game) CurrentPlayer() board.Player {
	return g.players[g.currentPlayerTurn]
}

func (g *Game) EndTurn() {
	// for now just copy over the state of the temp board to the real board
	// but we could do some fancy things like find which squares changed
	g.gameBoard.Copy(g.tempGameBoard)

	g.currentPlayerTurn = (g.currentPlayerTurn + 1) % len(g.players)
}

func (g *Game) FillSquare(player board.Player, coor util.Vector2) {
	g.tempGameBoard.GetSquare(coor).FillWith(player)
}

func (g *Game) CheckRules(rules []RuleFilter) {

}

// RuleFilter will determine what rules happen in play
// This could be things like:
// how many squares you can play on your turn,
// What squares you can play in,
// etc.
type RuleFilter interface{}
