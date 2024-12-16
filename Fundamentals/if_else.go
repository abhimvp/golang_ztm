package main
import "fmt"
func main(){
	// if else
	// var a int = 10
	// if a < 20 {
	// 	fmt.Println("a is less than 20")
	// } else {
	// 	fmt.Println("a is not less than 20")
	// }

	// if else if
	var a int = 100
	if a == 10 {
		fmt.Println("a is 10")
	} else if a == 20 {
		fmt.Println("a is 20")
	} else if a == 30 {
		fmt.Println("a is 30")
	} else {
		fmt.Println("a is not present")
	}
	// nested if else
	var b int = 100
	if b == 10 {
		fmt.Println("b is 10")
	} else if b == 20 {
		fmt.Println("b is 20")
	} else if b == 30 {
		fmt.Println("b is 30")
	} else {
		if b == 40 {
			fmt.Println("b is 40")
		} else if b == 50 {
			fmt.Println("b is 50")
		} else if b == 60 {
			fmt.Println("b is 60")
		} else {
			fmt.Println("b is not present")
		}
	}
	// switch case
	var c int = 10
	switch c {
	case 10:
		fmt.Println("c is 10")
	case 20:
		fmt.Println("c is 20")
	case 30:
		fmt.Println("c is 30")
	default:
		fmt.Println("c is not present")
	}
	// switch case with multiple cases
	var d int = 10
	switch d {
	case 10, 20:
		fmt.Println("d is 10 or 20")
	case 30, 40:
		fmt.Println("d is 30 or 40")
	default:
		fmt.Println("d is not present")
	}
	// switch case with fallthrough
	var e int = 10
	switch e {
	case 10:
		fmt.Println("e is 10")
		fallthrough
	case 20:
		fmt.Println("e is 20")
	case 30:
		fmt.Println("e is 30")
	default:
		fmt.Println("e is not present")
	}
	// switch case with expression
	var f int = 10
	switch {
	case f == 10:
		fmt.Println("f is 10")
	case f == 20:
		fmt.Println("f is 20")
	case f == 30:
		fmt.Println("f is 30")
	default:
		fmt.Println("f is not present")
	}
	// switch case with expression and fallthrough
	var g int = 10
	switch {
	case g == 10:
		fmt.Println("g is 10")
		fallthrough
	case g == 20:
		fmt.Println("g is 20")
	case g == 30:
		fmt.Println("g is 30")
	default:
		fmt.Println("g is not present")
	}
	// switch case with expression and multiple cases
	var h int = 10
	switch {
	case h == 10, h == 20:
		fmt.Println("h is 10 or 20")
	case h == 30, h == 40:
		fmt.Println("h is 30 or 40")
	default:
		fmt.Println("h is not present")
	}
	// switch case with expression and multiple cases and fallthrough
	var i int = 10
	switch {
	case i == 10, i == 20:
		fmt.Println("i is 10 or 20")
		fallthrough
	case i == 30, i == 40:
		fmt.Println("i is 30 or 40")
	default:
		fmt.Println("i is not present")
	}
	
}