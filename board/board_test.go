package board

import "testing"
import assert "gopkg.in/go-playground/assert.v1"

import "github.com/trungng92/goolang/util"

func TestFillSquare(t *testing.T) {
	square := Square{PlayerNeutral}
	assert.Equal(t, square.OwnedBy(PlayerNeutral), true)
	var playerTest Player
	square.FillWith(playerTest)
	assert.Equal(t, square.OwnedBy(playerTest), true)
}

// fill the whole board one square at a time
func TestFillBoardBySquare(t *testing.T) {
	dimX, dimY := 2, 2
	size := util.Vector2{dimX, dimY}
	gameBoard := createEmptyBoard(size)

	var playerTest Player
	for i := 0; i < len(gameBoard.Field); i++ {
		coor := util.IndexToVec(i, dimX)
		gameBoard.FillSquare(playerTest, coor)
		square := gameBoard.GetSquare(coor)
		assert.Equal(t, square.OwnedBy(playerTest), true)
		assert.Equal(t, square.OwnedBy(PlayerNeutral), false)
	}
}

func TestTrappingSquares(t *testing.T) {
	player1 := Player{"1"}
	player2 := Player{"2"}
	players := []Player{player1, player2}
	dimX, dimY := 3, 3
	size := util.Vector2{dimX, dimY}
	gameBoard := createEmptyBoard(size)

	// try various boards and make sure they look correct after trapping
	/*
	  0 1 2    1 1 2
	  0 1 0 -> 1 1 0
	  0 1 0    1 1 0
	*/

	squaresP1 := []util.Vector2{
		util.Vector2{1, 0},
		util.Vector2{1, 1},
		util.Vector2{1, 2},
	}

	squaresP2 := []util.Vector2{
		util.Vector2{2, 0},
	}

	for _, square := range squaresP1 {
		gameBoard.FillSquare(player1, square)
	}

	for _, square := range squaresP2 {
		gameBoard.FillSquare(player2, square)
	}

	gameBoard.FillTrappedSquares(players)
	finalBoard := []Player{
		player1, player1, player2,
		player1, player1, PlayerNeutral,
		player1, player1, PlayerNeutral,
	}
	for i, owner := range finalBoard {
		vec := util.IndexToVec(i, dimX)
		assert.Equal(t, gameBoard.GetSquare(vec).Owner, owner)
	}
}

func createEmptyBoard(size util.Vector2) Board {
	gameBoard := NewBoard(size)
	for i := 0; i < len(gameBoard.Field); i++ {
		gameBoard.FillSquare(PlayerNeutral, util.IndexToVec(i, size.X))
	}
	return gameBoard
}
