//  Basic building blocks of Go programs
// Allows functionality to be isolated, which makes programs easier to: Test,debug,modify,read,write,document
// Functions are simple : They take data as input and return data as output.
package main
import "fmt"
func add(x int, y int) int {
    return x + y
}

func double(x int) (int) {
     
    return x + x
}

func greet(){
	fmt.Println("Hello, World! , printing from greet function")
}

// excercise
//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
func greetPerson(name string) {
	fmt.Printf("Hello, %s!\n", name)
}
//* Write a function that returns any message, and call it from within
//  fmt.Println()
func message() string {
	return "This is a message from the message function."
}
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
func addThreeNumbers(a, b, c int) int {
	return a + b + c
}
//* Write a function that returns any number

//* Write a function that returns any two numbers

//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once
	
func main() {
    greet()
	dozen := double(6)
	fmt.Println("Your dozen is:", dozen)
	bakersDozen := add(dozen, 1)
	fmt.Println("And your bakers dozen is:", bakersDozen)

	anotherBakersDozen := add(double(6), 1)
	fmt.Println("And another bakers dozen is:", anotherBakersDozen)

	greetPerson("abhi")
	//fmt.Println(message())
	fmt.Println(message()) // Print the message returned by the message function
	sum := addThreeNumbers(1, 2, 3)
	fmt.Println("Sum of three numbers:", sum)

	}
// $ go run functions.go
// Hello, World! , printing from greet function
// Your dozen is: 12
// And your bakers dozen is: 13
// And another bakers dozen is: 13