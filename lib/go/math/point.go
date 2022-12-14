package math

import "image"

var (
	ManhattanMoves = []image.Point{
		// u, d, r, l
		{0, 1}, {0, -1}, {1, 0}, {-1, 0},
	}
	ChebyshevMoves = []image.Point{
		// u, d, r, l
		{0, 1}, {0, -1}, {1, 0}, {-1, 0},
		// ur, dr, ul, dl
		{1, 1}, {1, -1}, {-1, 1}, {-1, -1},
	}
)

func ClonePoint(pt image.Point) image.Point {
	return image.Point{pt.X, pt.Y}
}

func NormalizePoint(pt image.Point) image.Point {
	return image.Point{Sgn(pt.X), Sgn(pt.Y)}
}

// think 'moving' pt1 -> pt2 requires pt1 + this offset
func ChebyshevClosestOffset(pt1, pt2 image.Point) image.Point {
	return NormalizePoint(pt2.Sub(pt1))
}

func ChebyshevMove(pt1, pt2 image.Point) image.Point {
	return pt1.Add(ChebyshevClosestOffset(pt1, pt2))
}

func GetNeighbors(moves []image.Point, p image.Point) []image.Point {
	neighbors := []image.Point{}
	for _, move := range moves {
		neighbor := p.Add(move)
		neighbors = append(neighbors, neighbor)
	}

	return neighbors
}
