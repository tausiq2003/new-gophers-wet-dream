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
	// should structs be global in golang, generally yes, it can be in function scope, but its limited, like only be used in function
	coordinate := Coordinate{1, 2}
	fmt.Println(coordinate)
	fmt.Printf("%T", coordinate)                        //type of Coordinate
	fmt.Printf("\n%T\n", fmt.Sprintf("%v", coordinate)) // will give string
	fmt.Println(coordinate.x)                           //thats how we access the structs
	// coordinate.x = "string" //can't do this
	coordinate.x = 3 //changes the value
	p := coordinate  //creates a copy
	p.x = 4
	q := &coordinate
	q.x = 4 // now changes
	fmt.Println(coordinate.x)
	fmt.Println((*q).x)     // explicit dereferencing remember what we used to do in cpp? -> thats right but here we can do below,
	fmt.Println(q.x)        // all will give same results, implicit dereferencing
	fmt.Printf("\n%T\n", q) // although the type is not same, its a syntactic sugar
	coordinate1 := Coordinate{y: 1, x: 2}
	fmt.Println(coordinate1) // this is print {2 1}
	coordinate2 := Coordinate{}
	fmt.Println(coordinate2) // will print {0 0}
	r := new(Coordinate)
	fmt.Printf("\nthis is r %v and its type is %T\n", r, r) // new returns a pointer
	t := new(int)
	fmt.Println(*t) // it will give 0

}
func arrays() {
	//only one datatype
	var a [5]int = [5]int{1, 2, 3} // this is the bad declaration imo, disappointed
	fmt.Println(a)
	b := [5]int{2, 3, 4, 5} //this seems aight
	fmt.Println(b)          //access same as other langs
	//slices, kind of like vectors in c++, don't have fixed size
	c := []int{3, 3, 4, 52, 3, 42, 5, 4, 56, 2, 3, 2, 352, 53, 4, 2, 3, 5, 23, 52, 25, 3, 563, 6, 6, 3233} //this is called slice literal
	fmt.Println(c)
	d := c[4:9] //these points to reference copy
	//value copy
	vc := make([]int, len(d))
	copy(vc, d)
	vc[0] = 69 //doesn't change the d
	fmt.Println(vc)
	fmt.Println(d)
	//struct array
	s := []struct { // anonymous struct
		i   int
		str string
	}{
		{1, "a"},
		{2, "b"},
		{3, "c"},
		{4, "d"},
	}
	fmt.Println(s)
	//like every vector, arraylist, dynamic arrays, slices has both len and cap
	c = c[:5:5]
	fmt.Println(c)
	c = c[:5] // it will only run till this
	// c = c[:6] // this will give error
	fmt.Println(c)
	//below is the gotour example
	goTourEg := []int{2, 3, 5, 7, 11, 13}
	fmt.Printf("len=%d cap=%d %v\n", len(goTourEg), cap(goTourEg), goTourEg)

	goTourEg = goTourEg[:0]
	fmt.Printf("len=%d cap=%d %v\n", len(goTourEg), cap(goTourEg), goTourEg)

	// you might think how this work? but in golang if there is cap, capacity it references that array till 0->cap,
	//if you make cap 0 then it wont refer the original array, its like from: to/len (to - from): cap, from <= to <= cap
	goTourEg = goTourEg[:4]
	fmt.Printf("len=%d cap=%d %v\n", len(goTourEg), cap(goTourEg), goTourEg)

	goTourEg = goTourEg[2:]
	fmt.Printf("len=%d cap=%d %v\n", len(goTourEg), cap(goTourEg), goTourEg)
	//make
	foo := make([]int, 6)
	fmt.Println(foo, cap(foo))
	// you can also do like this but
	bar := new([]int)
	*bar = append(*bar, 2)

	fmt.Println(*bar, len(*bar), cap(*bar), (*bar)[0])
	/* why make is preferred to create slices, maps and channels, because
		   first you can see the new returns a dereferenced type or pointer to type,
		   where it can be error prone to work with it. Second, you can't specify len
		   or cap here. you have to manually append numbers which isn't a good approach,
		   i mean good approach but you have to make it again or append again and again to
		   increase the cap, also do you want to do like this? append(*bar, 0, 0, 0, ... n items).
		   No right, so use make, and new initializes a nil slice, it might be mem efficient,
	       but whats the reason for mem efficient, if you want to use n items. make is just
	       one-liner. Just a doubt i had so wrote
	*/
	baz := make([]int, 5, 5)
	// if there is not enough space in slice, then append creates a new slice, not always, here it appends, not inserts
	baz = append(baz, 1, 2, 3, 4, 5, 6)
	fmt.Println(baz, len(baz), cap(baz))

}
func ranges() {
	//only iterates over slice and map
	foo := []int{5, 6, 6, 7, 8, 9, 10}
	for i, v := range foo {
		foo[min(i+1, len(foo)-1)] = 4
		fmt.Println(i, v)
	}
	bar := []int{2, 45, 25, 32, 3, 5}
	for i := range bar {
		fmt.Println(i)
	}

	fmt.Println(foo)
}
func maps() {
	// m := make(map[string]int, 2) // atp i know it allocates cap, but won't explore more
	m := make(map[string]int)
	m["foo"] = 1
	m["bar"] = 2
	m["baz"] = 4
	fmt.Println(m)
	var m1 = map[string]Coordinate{
		// "fan": Coordinate{
		// 5, 4,
		// },
		// "photo": Coordinate{
		// 4, 4,
		// },
		// this gives a warning
		"fan":   {5, 4},
		"light": {4, 4},
	}
	fmt.Println(m1)
	m1["fan"] = Coordinate{6, 7}

	delete(m1, "light")
	// elem, ok := m1["fan"] // ok says whether the key is present or not
	elem, foo := m1["fan"] //needn't be be ok can be anything
	fmt.Println(elem, " exists in the map so ok's value is ", foo)

}
func main() {
	fmt.Println("hello")
	pointers()
	structs()
	arrays()
	ranges()
	maps()
}
