package main

import (
	"bytes"
	"fmt"
	"math"
	"unicode/utf8"
)

//iota 在const关键字出现时被重置为0，const中每新增一行常量声明 将使得iota计数一次（iota可以理解位const语句块中的行索引）。使用iota能简化定义。

const (
	a = iota
	_ = 100
	a2
	a3
	a4 = iota
)

const (
	d1, d2 = iota + 1, iota + 2
	d3, d4 = iota + 1, iota + 2
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	//fmt.Println(b,b2,b3)
	fmt.Println(a, a2, a3, a4)
	fmt.Println(d1, d2, d3, d4)
	fmt.Println(KB, GB, TB, PB)
	aa := 5 //101
	bb := aa >> 1
	fmt.Println(aa, bb)
	fmt.Println(math.MaxInt8)
	fmt.Printf("%v\n", "zhidi")
	fmt.Printf("%#v\n", "zhidi")
	fmt.Printf("%#v\n", 1)
	fmt.Printf("%#v\n", []int{1, 2, 3})
	fmt.Printf("%#v\n", []int{1, 2, 3})
	z := 'f'
	zz := '芳'
	fmt.Printf("%#v,%T\n", z, z)
	fmt.Printf("%#v,%T\n", zz, zz)
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	fmt.Println(sample)
	var f = '芳'
	fmt.Printf("%x\n", f)
	fmt.Printf("%q\n", f)
	var s rune = '芳'
	fmt.Printf("%T\n", s)
	fmt.Printf("%d\n", s)
	buf := bytes.NewBufferString("芳")
	fmt.Println(buf.Bytes())
	path := "''"
	fmt.Println(path)
	pp := '\''
	fmt.Println(pp)
	pt := `";'\n
	fjdls\njfld\s
	`
	fmt.Println(pt)
	ddd := "知道了"
	fmt.Println(len(ddd))
	fmt.Println(len([]rune(ddd)))
//	dc := []rune(ddd)
	fmt.Println(ddd)
	fmt.Println("................")
	fmt.Println(utf8.RuneCountInString("小的了"))
	for _,i := range ddd{
		fmt.Printf("%c\n",i)
	}
	fmt.Printf("%v,%v",'之',"之")

}
