package sprint

import "fmt"

type Point struct {
	X float32
	Y float32
	Text string
}

func PointText(p Point) Point {

	new := fmt.Sprintf("Text at", p.X, p.Y)
	return Point{X: p.X, Y: p.Y, Text: new}
}