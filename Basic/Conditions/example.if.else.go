package main

import "fmt"

func main(){
	var a int = 100;

	if ( a < 20 ) {
		fmt.Printf("a는 20보다 작다. \n")
	} else {
		fmt.Printf("a is not less than 20 \n")
	}
	fmt.Printf("a의 값은 %d \n", a)
}