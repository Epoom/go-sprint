package sprint

type Point struct {
	X float32
	Y float32
}

func PointDiff(p1, p2 Point) Point {

	coord1 := p1.X * p1.X + p1.Y * p1.Y
	coord2 := p2.X * p2.X + p2.Y * p2.Y

	if coord1 >= coord2 {
		return p1
	}
	return p2
}



