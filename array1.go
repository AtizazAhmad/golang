package main

import "fmt"

func inc_original_array(a *[3]int){
	for i,_:= range a{
		a[i] = a[i] + 1
	}
}
func main(){

	var a[3] int;
	var output[500] int;
	fmt.Println("Enter the values in array")
	for i,_ := range a{
		fmt.Scanln(&a[i])
	}
	
	inc_original_array(&a)
	fmt.Println("After increment in original array")
		for _,v := range a{
			fmt.Println(v)
		}
	fmt.Println("Reverse order")
	for i,_:=range a{
	    output[i] = a[len(a) - i -1]
	    fmt.Println(output[i]);
	}

}
