package main
import "fmt"
func main(){
	slice := []string{"Hello","world","!"}
	// range
	for i , element := range slice{
		fmt.Println(i,element,":")
		for _, ch := range element{
			fmt.Printf("  %q\n",ch)
		}
	}
	
}