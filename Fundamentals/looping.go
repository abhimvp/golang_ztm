package main
import "fmt"
func main(){
	//looping
	fmt.Println("looping")
	for i:=0; i<2;	i++{
		fmt.Println(i)
	}
	// for: while
	fmt.Println("for: while")
	j:=0
	for j<2{
		fmt.Println(j)
		j++
	}
	// for: infinite
	fmt.Println("for: infinite")
	k:=0
	for{
		fmt.Println(k)
		k++
		if k==2{
			break
		}
	}
	// continue
	// break
	fmt.Println("continue")
	for l:=0; l<2; l++ {
			if l==1{
				continue
			}
			fmt.Println(l)
		}

}