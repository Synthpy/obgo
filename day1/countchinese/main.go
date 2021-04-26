package main

import (
	"fmt"
	"unicode/utf8"
)


func main() {
	var chineseCounter int =0
	s := "hello,世界,ᢇ"
	b := []byte(s)
	for len(b) > 0 {
		r, size := utf8.DecodeRune(b)
		if r >  0x4E00 && r<=0x9FCB{
			chineseCounter ++
			fmt.Printf("%c---%x\n",r, r)
		}
		b = b[size:]
	}
	fmt.Println(chineseCounter)
}