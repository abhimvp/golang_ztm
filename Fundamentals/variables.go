package main // specifies we have a main package and allows us to have the func main
import "fmt" // fmt is a package that allows us to print to the console
// the main function is the first function that's ran whenever you run a Go program
func main(){
    var myName = "Abhishek"
    fmt.Println("my name is",myName)

    var name string = "Kavitha"
    fmt.Println("name =", name)

    userName := "abhimvp"
    fmt.Println("userName =", userName)

    var sum int
    fmt.Println("sum is",sum)

    part1 , other := 1,5
    fmt.Println("part1 =", part1, "other =", other)

    part2, other := 2,6
    fmt.Println("part2 =", part2, "other =", other)

    sum = part1 + part2
    fmt.Println("sum =", sum)

    var (
        lessonName = "Variables"
        lessonType = "Demo"
    )
    fmt.Println("lessonName =", lessonName)
    fmt.Println("lessonType =", lessonType)

    word1, word2, _ := "hello", "world", "!"
    fmt.Println(word1, word2)

}
// $ go run variables.go
// my name is Abhishek
// name = Kavitha
// userName = abhimvp
// sum is 0
// part1 = 1 other = 5
// part2 = 2 other = 6
// sum = 3
// lessonName = Variables
// lessonType = Demo
// hello world
// // Single creation
// var example = 3
// var example2 int = 3 // Hi gives compilation error
// var example3 int // deafults to zero
// // example3 = 3
// // compound creation
// var a,b,c = 1,2,"sample"
// // block creation
// var (
//     a int = 1
//     b int = 2
//     c = "sample"
// )
// // create and assign 
// example4 := 4 // `:=` we'll commonly see in Go code , because it's shorter and easier to write and you get type inference as well
// a,b := 1,"sample"
// // defaults - variables that are declared but not assigned values are given their zero value.
// var a int // 0
// var b string // ""
// // other default nil equivalant to null in other languages
// var a *int // nil 
// // comma ok idiom - allows a varibale to be reassigned in a creation statement.
// a,b := 1,2
// c,b := 3,4
// // Naming - Go variable naming conventions
// // camelCase - first word lower case, then each word's first letter capitalized
// myLongVariableName := 3 
// // use names appropriate for the data
// // Good
// totalGuests := 12
// // Bad
// ttl :=2
// // constants - Useful when declaring some value that needs to be utilized throughout some or all of the program - we start with capital letter
// const MaxSpeed = 30
// const MinPurchasePrice = 7.50
// const AppAuthor = "Bob"
 

 