package math

import "image"

func ChebyshevDistance(pt1, pt2 image.Point) int {
	return Max(Abs(pt1.X-pt2.X), Abs(pt1.Y-pt2.Y))
}
