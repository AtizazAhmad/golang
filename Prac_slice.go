package main

import "fmt"

func main(){
	b :=[]string{"g","o","l","a","n","g"}
	b[1] = "m"
	fmt.Println(b[:])
	q2 := b[1:4]
	fmt.Println(q2)	
	q3 := b[:2]
	fmt.Println(q3)
	q4 := b[2:]
	fmt.Println(q4)
	q5 := b[:]
	fmt.Println(q5)
	x :=[3]string{"Kutta","Very Kutta","Very very kutta"}
	s := x[:]
	for i,v := range x{
		fmt.Println(i)
		fmt.Println(v)
	}
	fmt.Println(s)
}