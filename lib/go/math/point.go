package math

import "image"

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
