package board

// For now have 2 players, although this number shouldn't need to change
var player1 = Player{}
var player2 = Player{}
var PlayerNeutral = Player{}

// Board contains all the squares/pieces,
// and it's what the player plays on.
type Board struct {
	Field [][]Square
}

func NewBoard(dimX, dimY int) Board {
	var field = make([][]Square, dimY)
	for y := 0; y < dimY; y++ {
		field[y] = make([]Square, dimX)
		for x := 0; x < dimX; x++ {
			field[y][x] = Square{PlayerNeutral}
		}
	}
	return Board{
		Field: field,
	}
}

func (b *Board) GetSquare(dimX, dimY int) Square {
	return b.Field[dimY][dimX]
}

func (b *Board) FillSquare(player Player, dimX, dimY int) {
	var square = b.Field[dimY][dimX]
	square.FillWith(player)
	// TODO: log that a board square was filled
	// and fire off an event
	// possibly redundant with square being filled event
}

// fillTrappedSquares takes the squares that are trapped (unreachable by other players)
// and fills them in with the player who trapped the squares
func (b *Board) FillTrappedSquares(players []Player) {
	// TODO: Implement
}

type Square struct {
	owner Player
}

func (s Square) OwnedBy(player Player) bool {
	return s.owner == player
}

func (s *Square) FillWith(player Player) {
	s.SetOwner(player)
	// TODO: log that a square was filled
	// and fire off an event
}

func (s *Square) SetOwner(player Player) {
	s.owner = player
}

type Player struct {
	name string
}
