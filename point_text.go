package sprint

import "fmt"

type Point struct {
	X float32
	Y float32
	Text string
}

func PointText(p Point) Point {

	fmt.Println("Text at", p.X, p.Y)

}