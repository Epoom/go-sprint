package sprint

type Circle struct{
	Radius float32
	Diameter float32
	Area float32
	Perimeter float32
}

func GetCircle(r float32) Circle {
    var π float32
	π = 3.14
	return Circle{Radius: r, Diameter: r*2, Area:π*(r*r), Perimeter: 2*π*r }

}