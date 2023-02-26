package values

const (
	White = iota
	Black
)

type Coord struct {
	X uint8
	Y uint8
}

type Move struct {
	X int
	Y int
}
