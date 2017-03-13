package game

import "testing"
import assert "gopkg.in/go-playground/assert.v1"

import "github.com/trungng92/goolang/util"
import "github.com/trungng92/goolang/board"

func TestEndingTurns(t *testing.T) {
	players := []board.Player{
		board.Player{"0"},
		board.Player{"1"},
		board.Player{"2"},
		board.Player{"3"},
	}

	// take a slice of each player length and try ending turns with that amount
	for i := 1; i <= len(players)-1; i++ {
		playersSlice := players[:i+1]
		game := Game{players: playersSlice}

		for j := 0; j <= len(players); j++ {
			game.EndTurn()
			playerIndex := (j + 1) % len(playersSlice)
			assert.Equal(t, game.CurrentPlayer(), playersSlice[playerIndex])
		}
	}
}

func TestEndTurnFinalizesBoard(t *testing.T) {
	player := board.Player{"1"}
	players := []board.Player{
		player,
	}

	boardSize := util.Vector2{4, 4}
	game := Game{
		players:       players,
		gameBoard:     board.NewBoard(boardSize),
		tempGameBoard: board.NewBoard(boardSize),
	}
	coor := util.IndexToVec(0, game.tempGameBoard.GetSize().X)
	assert.NotEqual(t, game.tempGameBoard.GetSquare(coor).Owner, player)
	assert.Equal(t, game.gameBoard.GetSquare(coor), game.tempGameBoard.GetSquare(coor))
	game.tempGameBoard.FillSquare(player, coor)
	assert.Equal(t, game.tempGameBoard.GetSquare(coor).Owner, player)
	game.EndTurn()
	assert.Equal(t, game.gameBoard.GetSquare(coor), game.tempGameBoard.GetSquare(coor))
}
