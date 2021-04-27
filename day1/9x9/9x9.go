package main

import "fmt"

func main() {
	for i := 0; i < 9; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Printf("%dx%d=%d ",i+1,j+1,(i+1)*(j+1))
		}
		fmt.Printf("\n")
	}
	for i := 0; i < 9; i++ {
		for j := 0; j <9-i ; j++ {
			fmt.Printf("%dx%d=%d ",9-i,j+1,(9-i)*(j+1))
		}
		fmt.Printf("\n")
	}
	s := "知道了"
	fmt.Println(len(s))
	for _,v := range s{
		fmt.Printf("%c",v)
	}


}