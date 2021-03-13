package main

import "fmt"

func inc_original_array(a *[3]int){
	for i,_:= range a{
		a[i] = a[i] + 1
	}
}
func main(){

	var a[3] int;
	fmt.Println("Enter the values in array")
	for i,_ := range a{
		fmt.Scanln(&a[i])
	}
	
	inc_original_array(&a)
	fmt.Println("After increment in original array")
		for _,v := range a{
			fmt.Println(v)
		}
}
