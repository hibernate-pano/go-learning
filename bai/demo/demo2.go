package main

import (
	"fmt"
)

func main() {
	//var name = `asndkasndkanslk12312312salads阿是豆豆`
	//fmt.Println(len(name))
	//for i := 0; i < len(name); i++ {

	//	fmt.Printf("0x%x ", name[i]) // 0xe4 0xb8 0xad 0xe5 0x9b 0xbd 0xe4 0xba 0xba}fmt.Printf("\n")
	//
	//}
	//fmt.Printf("\n")
	//fmt.Println(utf8.RuneCountInString(name))
	//for _, c := range name {
	//	fmt.Printf("0x%x ", c) // 0x4e2d 0x56fd 0x4eba}fmt.Printf("\n")
	//
	//}
	var s = "中国人"
	fmt.Printf("0x%x\n", s[0]) // 0xe4：字符“中” utf-8编码的第一个字节
	fmt.Println(s[0])
	fmt.Printf("%T", s[0])
}
