// Testing functions, arrays, slices...

package main

import "fmt"

func main() {
	x, y := 16, 25
	baz := add(x, y)
	foo, bar := multi_add(x, y)

	// Functions
	fmt.Println("Functions...")
	fmt.Printf("add() is %d\n", baz)
	fmt.Printf("multi_add() is %d and %d\n", foo, bar)

	// Arrays
	fmt.Println("\nArrays...")
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("len(a) = %d\n", len(a))
	fmt.Printf("a[5] = %d\n", a[5])
	fmt.Printf("a[] = { ")
	for _, v := range a {
		fmt.Printf("%d ", v)
	}
	fmt.Println("}")

	// Slices
	var b []int = a[1:4]
	fmt.Printf("\nSlices...\nb[] = a[1:4]\nlen(b) = %d\n", len(b))
	fmt.Printf("b[1] = %d\n", b[1])
	fmt.Printf("b[] = { ")
	for _, v := range b {
		fmt.Printf("%d ", v)
	}
	fmt.Println("}")

	b[1] = 77
	fmt.Printf("b[1] = 77\na[] = { ")
	for _, v := range a {
		fmt.Printf("%d ", v)
	}
	fmt.Println("}")

	// make()
	c := make([]int, 5, 5)
	fmt.Println(c)
}

func add(a int, b int) (sum int) {
	sum = a + b
	return sum
}

func multi_add(a int, b int) (sum, product int) {
	sum = a + b
	product = a * b
	return sum, product
}
