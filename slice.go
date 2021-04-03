package main

import "fmt"

func main(){

	months :=[...]string{1:"Jan",2:"Feb",3:"Mar",4:"Apr",5:"May",6:"June",7:"July",8:"Aug",9:"Sep",10:"Oct",11:"Nov",12:"Dec"}
	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2)
	fmt.Println(summer)
	for _,v := range Q2{
		for _,s := range summer{
			if s == v {
				fmt.Println(s," Appears to be same in both slices")
			}
		}
	}
	fmt.Println("Slices Have their own Index numbers different from Original Array from which slice made")
	for i,v := range summer{
		fmt.Println(i,v)
	}
	for i,v := range months{      // Implicitly intialized by Empty string at 0
		fmt.Println(i,v)
	}
	endlesssummer := summer[:-1]
	fmt.Println(endlesssummer)

}
	