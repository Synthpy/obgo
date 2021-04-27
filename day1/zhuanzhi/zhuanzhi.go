package main

import "fmt"

func main() {
	s := [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(s)
	ns := zz(s)
	fmt.Println(ns)
}

func zz(ar [][]int)[][]int{
	fmt.Println(len(ar))
	fmt.Println(len(ar[0]))
	var nar [][]int
	var nx []int
	for i:=0;i<len(ar[0]);i++{
		for j:=0;j<len(ar);j++{
			nx = append(nx, ar[j][i])
			fmt.Printf("%d  ",ar[j][i])
		}
		nar = append(nar, nx)
		nx = []int{}
		fmt.Printf("\n")
	}
	return nar
}
