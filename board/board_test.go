package board

import "testing"
import assert "gopkg.in/go-playground/assert.v1"

import "github.com/trungng92/goolang/util"

func TestFillSquare(t *testing.T) {
	var square = Square{PlayerNeutral}
	assert.Equal(t, square.OwnedBy(PlayerNeutral), true)
	var playerTest Player
	square.FillWith(playerTest)
	assert.Equal(t, square.OwnedBy(playerTest), true)
}

// fill the whole board one square at a time
func TestFillBoardBySquare(t *testing.T) {
	var dimX, dimY = 2, 2
	var size = util.Vector2{dimX, dimY}
	var gameBoard = NewBoard(size)
	for i := 0; i < len(gameBoard.Field); i++ {
		var square = gameBoard.GetSquare(util.IndexToVec(i, dimX))
		assert.Equal(t, square.OwnedBy(PlayerNeutral), true)
	}

	var playerTest Player
	for i := 0; i < len(gameBoard.Field); i++ {
		var coor = util.IndexToVec(i, dimX)
		gameBoard.FillSquare(playerTest, coor)
		var square = gameBoard.GetSquare(coor)
		assert.Equal(t, square.OwnedBy(playerTest), true)
		assert.Equal(t, square.OwnedBy(PlayerNeutral), true)
	}
}
