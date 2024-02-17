package main

import (
	"fmt"
	"coding/arrays"
)

var input = []int{1, 2, 3, 4, 5, 1}

func main(){
	fmt.Println("---")
	fmt.Println("Main function input :")
	fmt.Println(input)
	fmt.Println("---")

	result := arrays.ContainsDuplicate(input)

	fmt.Println(result)
}
