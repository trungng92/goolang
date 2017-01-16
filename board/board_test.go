package board

import "testing"
import assert "gopkg.in/go-playground/assert.v1"

func TestFillSquare(t *testing.T) {
	var square = Square{PlayerNeutral}
	assert.Equal(t, square.OwnedBy(PlayerNeutral), true)
	var playerTest Player
	square.FillWith(playerTest)
	assert.Equal(t, square.OwnedBy(playerTest), true)
}

func TestFillBoard(t *testing.T) {
	var dimX, dimY = 2, 2
	var gameBoard = NewBoard(dimX, dimY)
	for y := 0; y < dimY; y++ {
		for x := 0; x < dimX; x++ {
			var square = gameBoard.GetSquare(x, y)
			assert.Equal(t, square.OwnedBy(PlayerNeutral), true)
		}
	}

	var playerTest Player
	gameBoard.FillSquare(playerTest, 0, 0)
	for y := 0; y < dimY; y++ {
		for x := 0; x < dimX; x++ {
			gameBoard.FillSquare(playerTest, x, y)
			var square = gameBoard.GetSquare(x, y)
			assert.Equal(t, square.OwnedBy(playerTest), true)
			assert.Equal(t, square.OwnedBy(PlayerNeutral), true)
		}
	}
}
