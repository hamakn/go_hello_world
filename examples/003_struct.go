package examples

import "fmt"

type Vertex struct {
	X int
	Y int
}

func passByValue(v Vertex) {
	v.X += 1
	v.Y += 1
	fmt.Println(v)
}

func passByReference(v *Vertex) {
	v.X += 1
	v.Y += 1
	fmt.Println(*v)
}

func StructExample() {
	v := Vertex{1, 2}
	fmt.Println(v)

	fmt.Println("-- passByValue --")
	passByValue(v)
	fmt.Println(v)

	fmt.Println("-- passByReference --")
	passByReference(&v)
	fmt.Println(v)
}
