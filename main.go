package main

import (
    "fmt"
)

type Colour int
const (
	UnknownColour Colour = iota
	Red
	Orange
	Yellow
	Green
	Blue
	Purple
)

type Shape int
const (
    UnknownShape Shape = iota
    Circle
    Fourptstar
	Diamond
	Square
    Eightptstar
	Clover
)

const (
    TileQuantity = 3
	TileCount = 108
)

type Tile struct {
    Colour Colour
	Shape Shape
}

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

func createTileQueue(count , quantity int) []Tile {
	var queue []Tile

	num := 0
	for c:=Red; c<=Purple; c++ {
		for s:=Circle; s<=Clover; s++ {
			tile := Tile{c,s}
			for i:=0; i<quantity; i++{
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

func Hello() string {
    return "Hello world!"
}

func main() {
    fmt.Println(Hello())
}
