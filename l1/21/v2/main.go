package main

// практическая демонстрация на простом примере

import (
	"fmt"
	"math"
)

// группа классов,которые мы адаптируем: Rectangle, Circle

type Rectangle struct {
	a int
	b int
}

func (r *Rectangle) ShowRectangleArea() {
	fmt.Println("rectangle area:", r.a*r.b)
}

type Circle struct {
	r int
}

func (c *Circle) ShowCircleArea() {
	fmt.Println("circle area:", math.Pi*math.Pow(float64(c.r), 2))
}

// целевой интерфейс - Target
type FigureAdapter interface {
	ShowArea()
}

// адаптер для Rectangle
type RectangleAdapter struct {
	*Rectangle
}

func NewRectangleAdapter(rectangle *Rectangle) FigureAdapter {
	return &RectangleAdapter{
		rectangle,
	}
}

func (ra *RectangleAdapter) ShowArea() {
	ra.ShowRectangleArea()
}

type Square struct {
	a int
}

// уже реализует интефейс FigureAdapter -> для Square адаптер не нужен
func (s *Square) ShowArea() {
	fmt.Println("square area:", s.a*s.a)
}

// адаптер для Circle
type CircleAdapter struct {
	*Circle
}

func NewCircleAdapter(circle *Circle) FigureAdapter {
	return &CircleAdapter{
		circle,
	}
}

func (ca *CircleAdapter) ShowArea() {
	ca.ShowCircleArea()
}

func main() {
	rectangle := &Rectangle{a: 5, b: 10}
	square := &Square{a: 7}
	circle := &Circle{r: 4}
	figures := []FigureAdapter{NewRectangleAdapter(rectangle), square, NewCircleAdapter(circle)}
	for _, v := range figures {
		v.ShowArea()
	}
}
