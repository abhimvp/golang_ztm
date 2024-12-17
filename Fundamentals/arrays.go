package main
import "fmt"
type Room struct{
	name string
	cleaned bool
}

func checkCleanliness(rooms [4]Room){
	for i:=0;i<len(rooms);i++{
		room:=rooms[i]
		if room.cleaned{
			fmt.Println(room.name,"is clean")
		}else{
			fmt.Println(room.name,"is dirty")
		}
	}
}
func main(){
	roomsList := [...]Room{
		{name:"Office"},
		{name:"Living Room"},
		{name:"BathRoom"},
		{name:"BedRoom"},
	}
	checkCleanliness(roomsList)
	fmt.Println("Performing Cleaning...")
	roomsList[0].cleaned=true
	roomsList[2].cleaned=true
	roomsList[3].cleaned=true
	checkCleanliness(roomsList)


}
// // create an Array
// var myArray0 [3]int // unitialized
// myArray1 := [3]int{1,2,3} // create and assign
// myArray2 := [...]int{1,2,3} // three dots will look at how many items created and substitue that value in that place.  
// myArray3 := [4]int{7,8,9} // ELements not addressed in array initialization will be set to default values
// // Accessing Array Elements
// myArray := [3]int
// myArray[0] = 10
// myArray[1] = 20
// myArray[2] = 30
// item1:= myArray[0] // item1 = 10
// // iteration - Good practuce to assign the element to a variable during iteration
// // Easier to read in large functions / nested loops
// myArray4 = [...]int{7,8,9}
// for i=0;i<len(myArray4);i++{
// 	item:=myArray4[i] // item = 7,8,9
// 	fmt.Println(item)
// }
// // Bounds = Attempting to access an element outside the array bounds will cause a runtime error
// myArray5 := [...]int{7,8,9}
// // runtime error: index out of range
// for i=0;i<10;i++{
// 	fmt.Println(myArray5[i])
// }
// // compile time error
// myArray6 := [3]int{7,8,9}
// item3:=myArray6[3] // compile time error: invalid array index 3 (out of bounds for 3-element array)