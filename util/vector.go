package util

type Vector2 struct {
	X int
	Y int
}

func (v *Vector2) Add(other Vector2) {
	v.X += other.X
	v.Y += other.Y
}

// VecToIndex converts a 2d vector and width so that it's usable by a 1d array
// Side note: I like to use 1d arrays to make passing in and initializing objects easier
func VecToIndex(vec Vector2, width int) int {
	return vec.Y*width + vec.X
}

func IndexToVec(index, width int) Vector2 {
	return Vector2{index % width, index / width}
}

// Simple Queue class for Vector2 because go doesn't have generics :(
type QueueVector2 []Vector2

func (q *QueueVector2) Push(n Vector2) {
	*q = append(*q, n)
}

func (q *QueueVector2) Pop() Vector2 {
	var n = (*q)[0]
	*q = (*q)[1:]
	return n
}

func (q QueueVector2) Len() int {
	return len(q)
}
