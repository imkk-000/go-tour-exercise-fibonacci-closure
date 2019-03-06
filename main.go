package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var counter, resultNumber int
	f := [2]int{0, 1}
	return func() int {
		switch counter {
		case 0:
			resultNumber = f[0]
		case 1:
			resultNumber = f[1]
		default:
			temp := f[1]
			f[1] = f[0] + temp
			f[0] = temp
			resultNumber = f[1]
		}
		counter++
		return resultNumber
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
