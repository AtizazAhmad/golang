package main

import "fmt"  

func main() {
var n,sum,x int
fmt.Println("Enter number of integers :")
fmt.Scanf("%d",&n)
fmt.Println("Enter the Integer whose sum You want")
for i:=1;i<=n;i++{
fmt.Println("Integer",i)
fmt.Scanf("%d",&x)
sum+=x
}
fmt.Println("The Sum of Above Intgers") 
fmt.Println(sum)}
