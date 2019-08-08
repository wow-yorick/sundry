package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct{ X, Y float64 }

type ColoredPoint struct {
	*Point
	Color color.RGBA
}

type IntList struct {
	Value int
	Tail  *IntList
}

func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

func (p Point) Sub(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}

type Path []Point

func (path Path) TranslateBy(offset Point, add bool) {
	var op func(p, q Point) Point
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}
	for i := range path {
		path[i] = op(path[i], offset)
	}
}

func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main() {
	// p := Point{1, 2}
	// q := Point{4, 6}
	// fmt.Println(Distance(p, q))
	// fmt.Println(p.Distance(q))
	// r := Point{1, 2}
	// r.ScaleBy(2)
	// fmt.Println(&r)
	// fmt.Println(r)
	// m := url.Values{"lang": {"en"}}
	// m.Add("item", "1")
	// m.Add("item", "2")

	// fmt.Println(m.Get("lang"))
	// fmt.Println(m.Get("q"))
	// fmt.Println(m.Get("item"))
	// fmt.Println(m["item"])

	// m = nil
	// fmt.Println(m.Get("item"))
	// m.Add("item", "3")
	// var cp ColoredPoint
	// cp.X = 1
	// fmt.Println(cp.Point.X)
	// cp.Point.Y = 2
	// fmt.Println(cp.Y)
	// red := color.RGBA{255, 0, 0, 255}
	// blue := color.RGBA{0, 0, 255, 255}
	// var p = ColoredPoint{Point{1, 1}, red}
	// var q = ColoredPoint{Point{5, 4}, blue}
	// fmt.Println(p.Distance(q.Point))
	// p.ScaleBy(2)
	// q.ScaleBy(2)
	// fmt.Println(p.Distance(q.Point))
	// p := ColoredPoint{&Point{1, 1}, red}
	// q := ColoredPoint{&Point{5, 4}, blue}
	// fmt.Println(p.Distance(*q.Point))
	// q.Point = p.Point
	// p.ScaleBy(2)
	// fmt.Println(*p.Point, *q.Point)
	// p := Point{1, 2}
	// q := Point{4, 6}

	// distanceFromP := p.Distance
	// fmt.Println(distanceFromP(q))
	// var origin Point
	// fmt.Println(distanceFromP(origin))
	// scaleP := p.ScaleBy
	// scaleP(2)
	// scaleP(3)
	// scaleP(10)
	// fmt.Println(p)

	p := Point{1, 2}
	q := Point{4, 6}

	distance := Point.Distance
	fmt.Println(distance(p, q))
	fmt.Printf("%T\n", distance)

	scale := (*Point).ScaleBy
	scale(&p, 2)
	fmt.Println(p)
	fmt.Printf("%T\n", scale)

}
