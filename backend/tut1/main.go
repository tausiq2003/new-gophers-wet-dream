package main

import "fmt"

func pointers() {
	//pointers hold the memory address of a value
	//var p *int // this holds nil
	i := 42
	j := 23
	p := &i         // generates a pointer p that stores memory address of i
	fmt.Println(p)  // prints the memory address
	fmt.Println(*p) // prints the stored value
	fmt.Println(&p) // pointer which stores &p which is &(&i)
	//you can check by
	fmt.Println(*&p)
	fmt.Println(*(*&p)) //42 // * known as dereferencing and & known as address of operator
	//now values will change
	p = &j
	fmt.Println(p, ", ", *p)
	*p = *p * 3
	fmt.Println(j) // this will give now 69
	// p++ // we can't do pointer arithmetic, raises error
	// fmt.Println(*p)

}

type Coordinate struct {
	x int
	y int
}

func structs() {
	//like C, its a collection of fields
	// should structs be global in golang, idk for now, it can be in function scope, but its limited, like only be used in function
	coordinate := Coordinate{1, 2}
	fmt.Println(coordinate)
	fmt.Printf("%T", coordinate)                      //type of Coordinate
	fmt.Printf("\n%T", fmt.Sprintf("%v", coordinate)) // will give string

}

func main() {
	fmt.Println("hello")
	pointers()
	structs()
}
