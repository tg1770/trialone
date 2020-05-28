package main

import (
	"fmt"
)


func main() {

	s:="Hello playground"
	fmt.Println(s)

	bs := []byte(s) //coverting to type uint8
	fmt.Println(bs)
	fmt.Printf("%T\n",bs)

	for i:=0;i<len(s);i++{
		fmt.Printf("%#U",s[i]) //prints it as the Utf8 encoding scheme (rune)
	}

	for i, v := range s {
		fmt.Printf("at index position %d we have hex %#x\n", i, v)
	}
}
