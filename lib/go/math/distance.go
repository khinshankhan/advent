package math

import "image"

func ManhattanDistance(p1, p2 image.Point) int {
	return Abs(p1.X-p2.X) + Abs(p1.Y-p2.Y)
}

func ChebyshevDistance(p1, p2 image.Point) int {
	return Max(Abs(p1.X-p2.X), Abs(p1.Y-p2.Y))
}
