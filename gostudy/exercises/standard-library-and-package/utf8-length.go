package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println(utf8.RuneCountInString("你好"))

	//regex
	// 	s := `100
	// 00
	// 0
	// 1
	// 0

	// 0
	// 1`
	// 	pattern := regexp.MustCompile("(?m:^0(0|\n)*0)")
	// 	s = pattern.ReplaceAllString(s, "0")
	// 	fmt.Println(s)

	// bufio.Scanner.Scan() //best suited
	// 	fmt.Fscanf()         //only applicable to formated lines

	// 	bufio.Reader.ReadLine() //very low level
	// 	bufio.ReadString('\n')  //cannot handle EOF
}
