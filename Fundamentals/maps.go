package main
import "fmt"
func main(){
	shoppingList := make(map[string]int)
	// adding items to the map
	shoppingList["eggs"] = 7
	shoppingList["milk"] = 1
	shoppingList["bread"] += 2
	// updating items in the map
	shoppingList["eggs"] += 1
	// printing the map
	fmt.Println(shoppingList)
	// removing items from the map
	delete(shoppingList, "milk")
	// printing the map
	fmt.Println(shoppingList)
	eggs, ok := shoppingList["eggs"]
	// checking if the key exists
	if ok {
		fmt.Println(eggs)
	}
	cereal, ok := shoppingList["cereal"]
	if !ok {
		fmt.Println("cereal not found")
	}else{
		fmt.Println(cereal,"cereal found")
	}

	totalItems := 0
	for _, amount := range shoppingList {
		totalItems += amount
	}
	fmt.Println("Total items:", totalItems)
}
// // mkaing a Map
// myMap := make(map[string]int)
// myMap["key1"] = 1
// myMap["key2"] = 2
// myMap["key3"] = 3
// myMap["key4"] = 4
// myMap2 := map[string]int{"key 1":1,"key 2":2,"key 3":3,"key 4":4}
// // Map Operations
// fmt.Println(myMap["key1"])
// // Insert
// myMap["key5"] = 5
// delete(myMap,"key5")
// // Read
// fav:= myMap["key1"]
// missing := myMap["key6"] // default value
// // check existence
// _, ok := myMap["key1"]
// if ok {
//     fmt.Println("key1 exists")
// }
// // iterate
// for key, value := range myMap {
//     fmt.Println(key, value)
// }