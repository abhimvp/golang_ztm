package main
import "fmt"

type Counter struct{
	hits int
}

func increment(counter *Counter){
	// if we have a pointer to a structure , we don't have to dereference it with asterisk
	counter.hits+=1
	// (*counter).hits+=1
	fmt.Println("counter hits",counter)

}

func replace(old *string, new string, counter *Counter){
	*old = new
	increment(counter)
}

func main(){
	counter := Counter{}
	hello:="Hello"
	world := "World"
	fmt.Println(hello, world)
	replace(&hello, "Hi", &counter)
	replace(&world, "Mars", &counter)
	fmt.Println(hello, world)
	phrase := []string{hello, "*", world}
	fmt.Println(phrase)
	replace(&phrase[1], "Go!", &counter)
	fmt.Println(phrase)
	fmt.Println("counter hits", counter.hits)


}
// // using pointers
// // Asterisk when used with a pointer will dereference the pointer and get the value at the address
// func updateNumber(x *int){
// 	*x +=1 // deferencing x which points to i and updates i value.
// }
// i:=1
// updateNumber(&i) // this will have access to i memory address and this function will have access to i but not copy
// fmt.Println(i) //2