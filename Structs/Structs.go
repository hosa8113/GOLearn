package main

import "fmt"

//A struct is a collection of fields.

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	p  = &Vertex{1, 2} // has type *Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
)

func main() {
	letsGO()
	structFields()
	pointerToStructs()
	structLiterals()
}

func structLiterals() {
	fmt.Println(v1, p, v2, v3)
}

func pointerToStructs() {
	//Struct fields can be accessed through a struct pointer.
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

func structFields() {
	//Struct fields are accessed using a dot.
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}

func letsGO() {
	fmt.Println(Vertex{1, 2})
}
