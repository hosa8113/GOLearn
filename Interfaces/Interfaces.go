package main

import (
	runtime_utils "GOLearn/Runtime"
	"fmt"
	"math"
)

func main() {

	// taken from https://www.callicoder.com/golang-interfaces/#:~:text=An%20interface%20in%20Go%20is,for%20similar%20type%20of%20objects.&text=An%20interface%20is%20declared%20using,method%20signatures%20inside%20curly%20braces.

	simpleDemoWithShapeInterface()
	usingInterfaceTypesAsArgumentsToFunctions()
	usingInterfaceTypesAsFields()
	howDoesAnInterfaceWorkWithConcreteValues()
}

func simpleDemoWithShapeInterface() {
	fmt.Println("\n-- " + runtime_utils.GetCurrentFunctionName() + "() --")
	var s Shape = Circle{5.0}
	fmt.Printf("Shape Type = %T, Shape Value = %v\n", s, s)
	fmt.Printf("Area = %f, Perimeter = %f\n\n", s.Area(), s.Perimeter())

	s = Rectangle{4.0, 6.0}
	fmt.Printf("Shape Type = %T, Shape Value = %v\n", s, s)
	fmt.Printf("Area = %f, Perimeter = %f\n", s.Area(), s.Perimeter())

}

// Go Interface - `Shape`
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Struct type `Rectangle` - implements the `Shape` interface by implementing all its methods.
type Rectangle struct {
	Length, Width float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

// Struct type `Circle` - implements the `Shape` interface by implementing all its methods.
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) Diameter() float64 {
	return 2 * c.Radius
}

func usingInterfaceTypesAsArgumentsToFunctions() {
	fmt.Println("\n-- " + runtime_utils.GetCurrentFunctionName() + "() --")
	totalArea := CalculateTotalArea(Circle{2}, Rectangle{4, 5}, Circle{10})
	fmt.Println("Total area = ", totalArea)
}

// Generic function to calculate the total area of multiple shapes of different types
func CalculateTotalArea(shapes ...Shape) float64 {
	totalArea := 0.0
	for _, s := range shapes {
		totalArea += s.Area()
	}
	return totalArea
}

func usingInterfaceTypesAsFields() {
	fmt.Println("\n-- " + runtime_utils.GetCurrentFunctionName() + "() --")
	drawing := MyDrawing{
		shapes: []Shape{
			Circle{2},
			Rectangle{3, 5},
			Rectangle{4, 7},
		},
		bgColor: "red",
		fgColor: "white",
	}
	fmt.Println("Drawing", drawing)
	fmt.Println("Drawing Area = ", drawing.Area())
}

type MyDrawing struct {
	shapes  []Shape
	bgColor string
	fgColor string
}

func (drawing MyDrawing) Area() float64 {
	totalArea := 0.0
	for _, s := range drawing.shapes {
		totalArea += s.Area()
	}
	return totalArea
}

func howDoesAnInterfaceWorkWithConcreteValues() {
	fmt.Println("\n-- " + runtime_utils.GetCurrentFunctionName() + "() --")
	var s Shape

	s = Circle{5}
	fmt.Printf("(%v, %T, %#v)\n", s, s, s)
	fmt.Printf("Shape area = %v\n", s.Area())

	s = Rectangle{4, 7}
	fmt.Printf("(%v, %T, %#v)\n", s, s, s)
	fmt.Printf("Shape area = %v\n", s.Area())
}
