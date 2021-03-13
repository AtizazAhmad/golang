package main

import "fmt"


func main() {
    number := []int{2, 3, 4}
    sum := 0
    for j, num := range number {
        sum += num
        fmt.Printf("%d\n",j)
    }
    fmt.Printf("The sum of numbers :%d",sum)
}