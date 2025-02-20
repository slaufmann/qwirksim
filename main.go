package main

import (
	"fmt"
	"math/rand"
)

// Colour represents the colour of a [Tile].
type Colour int

// Values that are supported for [Colour].
const (
	Red Colour = iota
	Orange
	Yellow
	Green
	Blue
	Purple
)

// Shape represents the shape of a [Tile].
type Shape int

// Values that are supported for [Shape].
const (
	Circle Shape = iota
	Fourptstar
	Diamond
	Square
	Eightptstar
	Clover
)

// Values that serve as defaults.
// TileQuantity represents how often a specific [Tile] is present in a queue of tiles.
// TileCount represents the number of tiles in a queue.
const (
	TileQuantity = 3
	TileCount    = 108
)

// Tile represents a single tile in the qwirkle game. Each tile has a colour and a shape.
type Tile struct {
	Colour Colour
	Shape  Shape
}

// isColour returns a boolean indicating whether the given [Colour] c is a valid colour value.
func (c Colour) isColour() bool {
	switch c {
	case Red:
		fallthrough
	case Orange:
		fallthrough
	case Yellow:
		fallthrough
	case Green:
		fallthrough
	case Blue:
		fallthrough
	case Purple:
		return true
	}

	return false
}

// isShape returns a boolean indicating whether the given [Shape] s is a valid shape value.
func (s Shape) isShape() bool {
	switch s {
	case Circle:
		fallthrough
	case Fourptstar:
		fallthrough
	case Diamond:
		fallthrough
	case Square:
		fallthrough
	case Eightptstar:
		fallthrough
	case Clover:
		return true
	}

	return false
}

// Tile.GoString implements the GoStringer interface for the [Tile] type. It simply returns
// the string representation of the Colour and the Shape field like "(Colour Shape)".
func (t Tile) GoString() string {
	return fmt.Sprintf("(%#v %#v)", t.Colour, t.Shape)
}

// Colour.GoString implements the GoStringer interface for the [Colour] type. It simply returns
// the name of the colour. If the given colour is none of the defined ones, "Unknown Colour" is
// returned.
func (c Colour) GoString() string {
	str := ""

	switch c {
	case Red:
		str = "Red"
	case Orange:
		str = "Orange"
	case Yellow:
		str = "Yellow"
	case Green:
		str = "Green"
	case Blue:
		str = "Blue"
	case Purple:
		str = "Purple"
	default:
		str = "Unknown Colour"
	}

	return str
}

// Shape.GoString implements the GoStringer interface for the [Shape] type. It returns the name
// of the shape. If the given shape is none of the defined ones, "Unknown Shape" is returned.
func (s Shape) GoString() string {
	str := ""

	switch s {
	case Circle:
		str = "Circle"
	case Fourptstar:
		str = "4ptStar"
	case Diamond:
		str = "Diamond"
	case Square:
		str = "Square"
	case Eightptstar:
		str = "8ptStar"
	case Clover:
		str = "Clover"
	default:
		str = "Unknown Shape"
	}

	return str
}

// swapTileInQueue swaps two tiles in the given queue by using the given indices. The altered
// queue is returned.
func swapTileInQueue(queue []Tile, a, b int) []Tile {
	temp := queue[b]

	queue[b] = queue[a]
	queue[a] = temp

	return queue
}

// createTileQueue returns a slice of [Tile] elements with a maximum size as given by the count
// argument. Each tile can be contained multiple times but not more than the quantity argument.
func createTileQueue(count, quantity int) []Tile {
	var queue []Tile

	num := 0
	for c := Red; c <= Purple; c++ {
		for s := Circle; s <= Clover; s++ {
			tile := Tile{c, s}
			for i := 0; i < quantity; i++ {
				queue = append(queue, tile)
				num++
				if num >= count {
					return queue
				}
			}
		}
	}

	return queue
}

// shuffleTileQueue shuffles a given [Tile] queue by swapping tiles in the queue randomly. The
// queue is shuffled in place therefore altered during the function execution.
func shuffleTileQueue(queue []Tile) []Tile {
	rand.Seed(1337)

	for index, _ := range queue {
		swapIndex := rand.Intn(len(queue))
		swapTileInQueue(queue, index, swapIndex)
	}

	return queue
}

func main() {
	queue := createTileQueue(TileCount, TileQuantity)
	queue = shuffleTileQueue(queue)

	fmt.Printf("This is the queue:\n%#v\n", queue)
}
