package main

import"fmt"

func main(){
     x:=1
     var p *int
     p = &x
      q:=&p
     fmt.Println(p,x,*p,q,*q,**q)
}
