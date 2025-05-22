package main

import (
	"fmt"
	"math/cmplx"
	"math/rand"
	"runtime"
	"time"
)

func add(x int, y int) (int, int) {
	return x, y
}
func swap(x, y int) (int, int) {
	return y, x
}
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return //naked return
}

func forloop() {
	fmt.Println("simple for loop")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println("simple while loop")
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}
	fmt.Println("simple do-while loop")
	k := 0
	for {
		fmt.Println(k)
		k++
		if k >= 5 {
			break
		}

	}

}
func conditionals() string {
	isSunday := false
	if !isSunday {
		fmt.Println("Its working day")
	}
	canDrive := true
	if age := rand.Intn(60); age >= 18 && canDrive { //short if statement
		fmt.Printf("Your age is %v\n", age)
		fmt.Println("You have/should have driving license")
	} else if age < 18 && isSunday { // can access here
		fmt.Println(age)
		fmt.Println("Cycling time")
	} else {
		fmt.Println("oh no")
	}
	// return fmt.Sprintf("Exit from the function with age %v\n", age) // Can't access age here
	return fmt.Sprint("Why I need this?") // this is for demo purpose you can return direct string too, Sprintf is helpful tho

}
func switcher() { // expected INDENT, found switch which means I can't be asshole and use keywords
	fmt.Println("Go runs on", runtime.GOOS, "at", runtime.GOARCH)
	days := [8]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday", "Test"}
	switch choice := days[rand.Intn(len(days))]; choice {
	case "Monday":
		fmt.Println("Work starts")
	case "Tuesday", "Wednesday", "Thursday": // or statements
		fmt.Println("Hell work")
	case "Friday":
		fmt.Println("Production day")
	case "Saturday":
		fmt.Println("Production will break")
	case "Sunday":
		fmt.Println("Fuck production, its funday")

	default:
		fmt.Println("What are you smoking bro?")

	}
	fmt.Println("today is ", time.Now().Weekday(), "tomorrow is, ", time.Now().Weekday()+1)
	// lots to explore in time module, it will come in experience
	n := byte('1') // returns in ascii
	fmt.Println(n)
	fmt.Println("HOLD ON another switch")
	//switch with no condition, it will act as if-else if-else
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	case t.Hour() < 21:
		fmt.Println("Good evening")
	case t.Hour() > 22:
		fmt.Println("Good night")
	}

}
func deferredThing() {
	//defer panic and recover
	/* till now i think defer as this
	   defer is like a stack which executes in lifo order
	*/
	fmt.Println("first")
	defer callingfunction()
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("second")
	//there is a defer stack for every stack frame, not to be confused with js's eventloop which executes after empty stack, but here it depends on stack frame every stack frame have defer stack and it will execute then it will pop and stack's top will point to calling function

}
func callingfunction() {
	fmt.Println("third")
	for i := 10; i < 20; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("fourth")
}
func foo() (i int) {
	// in golang we can change return value which is being set, with defer means the function executes, return is about to being returned then we can do defer to change the return value
	defer func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)

		}
	}()
	defer func() { fmt.Println("helloworld"); i++ }() // function created and called like iife in js, its called closure in golang
	return 1                                          // returns 2

}
func anExample() {
	i := 0
	defer fmt.Println(i) // it prints 0, because the value is captured immediately and its there in defer list
	i++
	j := 0
	defer func() { // this is a closure and in closure the value is accessed by reference not by values so it prints 1
		fmt.Println("here it is", j)
	}()
	j++
	return
}

func main() {
	fmt.Printf("helloworld")
	fmt.Println("time is ", time.Now())
	fmt.Println("My favourite number is ", rand.Intn(10))
	hello := "asdf"
	fmt.Printf("dingle nuts %v\n", hello)
	c, d := add(4, 5)
	fmt.Println(c, d)
	fmt.Println(split(17))
	var str, b, f, i = "hello", false, 3.44, 2
	fmt.Println(str, b, f, i)
	var a, b1, c1 int = 1, 2, 3
	fmt.Println(a, b1, c1)
	var comp = cmplx.Sqrt(-16 + 9i)
	fmt.Println(comp)
	fmt.Println(int(1 + 10e1))
	forloop()
	val := conditionals()
	fmt.Println(val)
	// defer switcher()
	switcher()
	deferredThing()
	fmt.Println(foo())
	anExample()

}
