package main
import ("fmt")
// Defining a structure
// type Sample struct {
// 	field string
// 	a,b int
// }
type Passenger struct{
	Name string
	TicketNumber int
	Boarded bool
}
type Bus struct{
	FrontSeat Passenger
}
type Rectangle struct{
	Length int
	Width int
}
func area(rect Rectangle)int{
	return rect.Length*rect.Width
}
func perimeter(rect Rectangle)int{
	return 2*(rect.Length+rect.Width)
}
func printInfo(rect Rectangle){
	fmt.Println("Area:", area(rect))
	fmt.Println("Perimeter:", perimeter(rect))
}
func main(){

	casey:= Passenger{"Casey", 1, false}
	fmt.Println(casey)

	var(
		bill= Passenger{Name: "Bill", TicketNumber: 2}
		ella= Passenger{Name: "Ella", TicketNumber: 3}
	)
	fmt.Println(bill,ella)
	var heidi Passenger // all defaults
	heidi.Name = "Heidi"
	heidi.TicketNumber = 4
	fmt.Println(heidi)

	casey.Boarded = true
	bill.Boarded = true
	if bill.Boarded {
		fmt.Println("Bill has boarded the bus")
	}
	if casey.Boarded {
		fmt.Println(casey.Name,"Casey has boarded the bus")
	}
	heidi.Boarded = true
	bus := Bus{heidi}
	fmt.Println(bus) 
	fmt.Println(bus.FrontSeat.Name, "is in the front seat")

	rect:= Rectangle{10, 5}
	printInfo(rect)
	rect.Length*=2
	rect.Width*=2
	printInfo(rect)
	

}
// Instantiating a structure
// sample := Sample{"abc", 1, 2}
// // Default Values - Any fields not indicated during instantiation will be set to their default values
// s := Sample{a: 1, b: 2} // field will be set to its default value
// data := Sample{}
// // Accessing fields - Fields can be read from and written to
// word:= sample.field
// data.field = "def"
// Anonymous Structures - Possible to create anonymous/inline structures inside of a function
// useful when working with library functions or when shipping data across a network , can easily define the data structure as needed.
// inline structs created using var will have default values
// var sample1 struct {
// 	field string
// 	a,b int
// }
// sample1.field = "abc"
// sample1.a = 1
// // this below needs to be give values and you can't be missing any fields
// sample2 := struct { // := create and assign
// 	field string
// 	a,b int
// }{"abc", 1, 2}
