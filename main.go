package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	counter := -1
	f := []int{0, 1}
	return func() int {
		if counter >= 1 {
			f = append(f, f[counter]+f[counter-1])
		}
		counter++
		return f[counter]
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
