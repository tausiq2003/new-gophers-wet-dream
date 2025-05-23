package main

import (
	"fmt"
	"math"
	"strings"

	"golang.org/x/tour/pic"
	"golang.org/x/tour/wc"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		prevZ := z
		z -= (z*z - x) / (2 * z)
		if math.Abs(prevZ-z) < 1e-15 {
			break
		}
		fmt.Println("Testing", z)
	}
	return z
}
func Pic(dx, dy int) [][]uint8 {
	// this is the most fun exercise till now,
	//thanks go tour, can't run in alacritty, can do in chafa but leave it
	res := make([][]uint8, dy)
	for i := range res {
		res[i] = make([]uint8, dx)

	}
	for i := range res {
		for j := range res[i] {
			res[i][j] = uint8((i & j)) //try good in future

		}
	}

	return res
}
func WordCount(s string) map[string]int {
	words := strings.Split(s, " ")
	m := make(map[string]int, len(words))
	for _, v := range words {
		_, ok := m[v]
		if !ok {
			m[v] = 1
		} else {
			m[v]++
		}

	}
	return m
}
func fibonacci() func() int {
	a := 0
	b := 1

	return func() int {
		c := a
		a = b
		b = a + c
		return c
	}

}

func main() {
	fmt.Println(Sqrt(10))
	fmt.Println(Pic(2, 3))
	pic.Show(Pic)
	wc.Test(WordCount)
	f := fibonacci()
	fmt.Println(f) //this stores everything, so its working, it keeps its state
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

}
