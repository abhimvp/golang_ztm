package main
import "fmt"

func printSlice(title string,slice []string){
	fmt.Println()
	fmt.Println("---",title,"---")
	for i:=0;i<len(slice);i++{
		element := slice[i]
		fmt.Println(element)
}
}

func main(){
	route := []string{"grocery","Department Store","Salon"}
	printSlice("Route 1", route)
	route = append(route, "Home") // append is a built-in function that adds an element to the end of a slice
	// route = append(route, "Home", "Home2") // append is a built-in function that adds an element to the end of a slice
	printSlice("Route 2", route)
	fmt.Println(route[0],"visited")
	route = route[2:]
	printSlice("Remaining Locations",route)
}
// // Slices and an underlying array can be created at the same time 
// mySlice := []int{1,3,5,7,9,11,13} //slices - comapring with arrays we don't have anything in squaRE BRACKETS
// myArray := [7]int{1,3,5,7,9,11,13} //arrays
// // Accessing elements in slice is the same as accessing elements in an array
// fmt.Println(mySlice[2])
// fmt.Println(myArray[2])
// // Slice Syntax - can create slices from specific elements in an array
// mySlice := myArray[2:5] // 2 is the starting index and 5 is the ending index - inclusive:exclusive
// fmt.Println(mySlice) 
// mySlice := myArray[2:] // 2 is the starting index and 5 is the ending index -  inclusive:exclusive
// numbers := [...]int{1,2,3,4,5,6,7,8,9,10}
// mySlice := numbers[2:5] // 2 is the starting index and 5 is the ending index -  inclusive:exclusive
// slice1:=numbers[:] // all elements
// slice2:=numbers[2:] // 3,4,5,6,7,8,9,10
// slice3:=numbers[:5] // 1,2,3,4,5
// slice4:=numbers[2:5] // 3,4,5
// // Dynamic arrays - slices can be dynamically resized
// mySlice := []int{1,3,5,7,9,11,13}
// mySlice = append(mySlice, 15) // append is a built-in function that adds an element to the end of a slice
// fmt.Println(mySlice)
// mySlice = append(mySlice, 17,19,21)
// // 3 dots can be used to extend a slice with another slice
// mySlice = append(mySlice, []int{23,25,27}...)
// // Preallocation - can be used to preallocate memory for a slice
// mySlice := make([]int, 5, 10) // 5 is the length of the slice and 10 is the capacity of the slice
// // make() function is used to preallocate a slice
// // Useful when number of elements is known , but their values are still unknown
// slice:= make([]int,10) // will have default values of 0
// slice:= make([]int, 10, 100) // will have 10 elements and capacity of 100
// // Slices to Functions
// func main(){
//   slice:=[]int{1,2,3,4,5}
//   fmt.Println(sum(slice))
// }
// func iterate(slice []int){
//   for i:=0;i<len(slice);i++{
//     fmt.Println(slice[i])
//   }
// }
// small := []int{1}
// big:=[]int{1,2,3,4,5}
// iterate(small)
// iterate(big)
// // Multidimensional Slices
// matrix:=[][]int{{1,2,3},{4,5,6},{7,8,9}}