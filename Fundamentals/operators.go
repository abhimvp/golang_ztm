package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20
	var c int = 0

	c = a + b
	fmt.Printf("a + b = %d\n", c)

	c = a - b
	fmt.Printf("a - b = %d\n", c)

	c = a * b
	fmt.Printf("a * b = %d\n", c)

	c = a / b
	fmt.Printf("a / b = %d\n", c)

	c = a % b
	fmt.Printf("a %% b = %d\n", c)

	a++
	fmt.Printf("a++ = %d\n", a)

	a--
	fmt.Printf("a-- = %d\n", a)
	//relational operators
	var x bool = true
	var y bool = false
	if x && y {
		fmt.Printf("Line 1 - Condition is true\n")
	} else {
		fmt.Printf("Line 1 - Condition is not true\n")
	}
	//logical operators
	if x || y {
		fmt.Printf("Line 2 - Condition is true\n")
	} else {
		fmt.Printf("Line 2 - Condition is not true\n")
	}
	if !(x && y) {
		fmt.Printf("Line 3 - Condition is true\n")
	} else {
		fmt.Printf("Line 3 - Condition is not true\n")
	}
	//bitwise operators
	var d int = 5
	var e int = 2
	//bitwise AND
	fmt.Printf("d & e = %d\n", d & e)
	fmt.Printf("d | e = %d\n", d | e)
	//bitwise XOR
	fmt.Printf("d ^ e = %d\n", d ^ e)
	fmt.Printf("d << 1 = %d\n", d << 1)
	//bitwise right shift
	fmt.Printf("d >> 1 = %d\n", d >> 1)

}
