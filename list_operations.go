package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println(a[0])

	fmt.Println(a[3:4])

	n := len(a)

	fmt.Println(n)

	e := a[n-4:n-3][0]
	fmt.Println(e)
}
