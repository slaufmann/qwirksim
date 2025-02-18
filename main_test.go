package main

import "testing"

const (
	TestTileCount    = int(30)
	TestTileQuantity = int(1)
)

func assertEqual(t testing.TB, got, want bool) {
	t.Helper()

	if got != want {
		t.Errorf("got %v but want %v", got, want)
	}
}

func assertEqualInt(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %v but want %v", got, want)
	}
}

func findTileInQueue(tb testing.TB, t Tile, q []Tile) []int {
	tb.Helper()

	var indices []int

	for index, tile := range q {
		if isEqualTile(tb, t, tile) {
			indices = append(indices, index)
		}
	}

	return indices
}

func isEqualTile(t testing.TB, a, b Tile) bool {
	t.Helper()

	if a.Colour == b.Colour && a.Shape == b.Shape {
		return true
	}

	return false
}

func TestRedIsColour(t *testing.T) {
	var given Colour = Red
	got := given.isColour()
	want := true

	assertEqual(t, got, want)
}

func TestUnknownColour(t *testing.T) {
	var given Colour = 10
	got := given.isColour()
	want := false

	assertEqual(t, got, want)
}

func TestCircleIsShape(t *testing.T) {
	var given Shape = Circle
	got := given.isShape()
	want := true

	assertEqual(t, got, want)
}

func TestUnknownShape(t *testing.T) {
	var given Shape = 10
	got := given.isShape()
	want := false

	assertEqual(t, got, want)
}

func TestCreateTileQueueGivesCorrectLength(t *testing.T) {
	queue := createTileQueue(TestTileCount, TestTileQuantity)

	assertEqualInt(t, len(queue), int(TestTileCount))
}

func TestCreateTileQueueGivesValidTiles(t *testing.T) {
	queue := createTileQueue(TestTileCount, TestTileQuantity)

	for _, tile := range queue {
		got := tile.Colour.isColour() && tile.Shape.isShape()
		want := true

		assertEqual(t, got, want)
	}
}

func TestCreateTileQueueGivesCorrectQuantities(t *testing.T) {
	queue := createTileQueue(TestTileCount, TestTileQuantity)

	for _, tile := range queue {
		got := len(findTileInQueue(t, tile, queue))
		want := TestTileQuantity

		assertEqualInt(t, int(got), int(want))
	}
}
