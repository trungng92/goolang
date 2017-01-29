package board

import "github.com/trungng92/goolang/util"

// For now have 2 players, although this number shouldn't need to change
var player1 = Player{}
var player2 = Player{}
var PlayerNeutral = Player{"N"}

// Board contains all the squares/pieces,
// and it's what the player plays on.
type Board struct {
	Field []Square
	size  util.Vector2
}

func NewBoard(size util.Vector2) Board {
	field := make([]Square, size.X*size.Y)
	for i := 0; i < len(field); i++ {
		field[i] = Square{PlayerNeutral}
	}
	return Board{
		Field: field,
		size:  size,
	}
}

func (b Board) IsValidCoordinate(coor util.Vector2) bool {
	return coor.X >= 0 &&
		coor.X < b.size.X &&
		coor.Y >= 0 &&
		coor.Y < b.size.Y
}

func (b *Board) GetSquare(coor util.Vector2) *Square {
	return &b.Field[util.VecToIndex(coor, b.size.X)]
}

func (b *Board) FillSquare(player Player, coor util.Vector2) {
	b.GetSquare(coor).FillWith(player)
	// TODO: log that a board square was filled
	// and fire off an event
	// possibly redundant with square being filled event
}

// fillTrappedSquares takes the squares that are trapped (unreachable by other players)
// and fills them in with the player who trapped the squares
func (b *Board) FillTrappedSquares(players []Player) {
	checkedBoard := make([]bool, len(b.Field))
	for i := 0; i < len(checkedBoard); i++ {
		println("iterating board item ", i)
		vec := util.IndexToVec(i, b.size.X)
		if !checkedBoard[i] && b.GetSquare(vec).Owner == PlayerNeutral {
			println("checking unchecked item ", i)
			checkNext := util.QueueVector2{util.IndexToVec(i, b.size.X)}
			playersReached := make(map[Player]struct{})
			b.recursivelyFindTrappedSquares(&checkedBoard, &checkNext, &playersReached)
		}
	}
}

// recursivelyFindTrappeSquares starts with a square and searches surrounding squares to see if all the squares connected are trapped.
// After no more squares can be reached, if only one player can be reached,
// then we use the stack to determine all the nodes to be awarded to that player.
func (b *Board) recursivelyFindTrappedSquares(checkedBoard *[]bool, checkNext *util.QueueVector2, playersReached *map[Player]struct{}) {
	if checkNext.Len() == 0 {
		return
	}
	current := checkNext.Pop()
	println("Currently checking ", current.X, ", ", current.Y)
	(*checkedBoard)[util.VecToIndex(current, b.size.X)] = true

	// just naiively check in 8 directions to find unchecked squares
	// this array and the vectors in it are allocated on the stack, so it shouldn't be that expensive
	toCheck := [...]util.Vector2{
		util.Vector2{-1, -1}, util.Vector2{-1, 0}, util.Vector2{-1, 1},
		util.Vector2{0, -1}, util.Vector2{0, 1},
		util.Vector2{1, -1}, util.Vector2{1, 0}, util.Vector2{1, 1},
	}
	for i := 0; i < len(toCheck); i++ {
		next := current
		next.Add(toCheck[i])
		if b.IsValidCoordinate(next) {
			square := b.GetSquare(next)
			// we only need to check squares that aren't owned by a player
			if !square.OwnedBy(PlayerNeutral) {
				println("Found square ", next.X, ", ", next.Y, " is owned by ", square.Owner.name)
				(*playersReached)[square.Owner] = struct{}{}
				(*checkedBoard)[util.VecToIndex(next, b.size.X)] = true
			} else if !(*checkedBoard)[util.VecToIndex(next, b.size.X)] {
				println("Found square ", next.X, ", ", next.Y, " is neutral. adding to check.")
				checkNext.Push(next)
			}
		}
	}

	b.recursivelyFindTrappedSquares(checkedBoard, checkNext, playersReached)
	println("Players reached: ", len(*playersReached))
	if len(*playersReached) == 1 {
		for k := range *playersReached {
			b.GetSquare(current).FillWith(k)
			println("Filling ", current.X, ", ", current.Y, " with ", k.name)
		}
	}
}

func (b Board) Print() {
	for i := 0; i < b.size.Y; i++ {
		for j := 0; j < b.size.X; j++ {
			coor := util.Vector2{j, i}
			print(b.GetSquare(coor).Owner.name)
		}
		println()
	}
}

type Square struct {
	Owner Player
}

func (s Square) OwnedBy(player Player) bool {
	return s.Owner == player
}

func (s *Square) FillWith(player Player) {
	s.Owner = player
	// TODO: log that a square was filled
	// and fire off an event
}

type Player struct {
	name string
}
