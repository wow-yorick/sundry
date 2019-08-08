package main

import "fmt"

// type S struct {
// 	m string
// }

// func (s S) F() {

// }

// type IF interface {
// 	F()
// }

// func InitType() S {
// 	var s S
// 	return s
// }

// func InitPointer() *S {
// 	var s *S
// 	return s
// }

// func InitEfaceType() interface{} {
// 	var s S
// 	return s
// }

// func InitEfacePointer() interface{} {
// 	var s *S
// 	return s
// }

// func InitIfaceType() IF {
// 	var s S
// 	return s
// }

// func InitlfacePointer() IF {
// 	var s *S
// 	return s
// }

// func main() {
// fmt.Println(InitType() == nil) //cannot convert nil to type S
// 	fmt.Println(InitPointer() == nil)
// 	fmt.Println(InitEfaceType() == nil)
// 	fmt.Println(InitEfacePointer() == (*S)(nil))
// 	fmt.Println(InitIfaceType() == nil)
// 	fmt.Println(InitlfacePointer() == nil)
//OkIdiom()
// p := f()
// fmt.Println(p.m)

// s := S{}
// p := &s
// var i interface{}

// f(s)
// //g(p)
// g(&i)
// f(p)
// //g(p)
// }

// func f() *S {
// 	return &S{
// 		m: "foo",
// 	}
// }

// func OkIdiom() {
// 	var m = make(map[string]int) //A
// 	m["a"] = 1
// 	if v, ok := m["b"]; !ok { //B
// 		fmt.Println(v)
// 	}
// }

// func f(x interface{}) {
// }

// func g(x *interface{}) {
// }

// const N = 3

//临时指针
// func main() {
// 	m := make(map[int]*int)

// 	for i := 0; i < N; i++ {
// 		j := int(i)
// 		m[i] = &j
// 	}

// 	for _, v := range m {
// 		fmt.Println(*v)
// 	}
// }

//标签
// func main() {

// outer:
// 	for i := 0; i < 3; i++ {
// 		for j := 0; j < 3; j++ {
// 			fmt.Println(i, ",", j, " ")
// 			break outer
// 		}
// 		fmt.Println("运行不到")
// 	}

// 	fmt.Println("finish")

// LABEL1:
// 	for i := 0; i <= 5; i++ {
// 		for j := 0; j <= 5; j++ {
// 			if j == 4 {
// 				break LABEL1
// 			}
// 			fmt.Printf("i is: %d, and j is: %d\n", i, j)
// 		}
// 	}

// 	fmt.Println("finish")
// 	a := 1
// 	goto TARGET // compile error
// TARGET:
// 	var b int
// 	b += a
// 	fmt.Printf("a is %v *** b is %v", a, b)

// i := 0
// for { //since there are no checks, this is an infinite loop
// 	if i >= 3 {
// 		break
// 	}
// 	//break out of this for loop when this condition is met
// 	fmt.Println("Value of i is:", i)
// 	i++
// }
// fmt.Println("A statement just after for loop.")

// for i := 0; i < 7; i++ {
// 	if i%2 == 0 {
// 		continue
// 	}
// 	fmt.Println("Odd:", i)
// }

//defer 需要在判断错误之后
// f, err := os.Open("file")
// if err != nil {
// 	return
// }
// defer f.Close()
// b, err := ioutil.ReadAll(f)
// println(string(b))

// }

func f1() {
	defer println("f1-begin")
	f2()
	defer println("f1-end")
}

func f2() {
	defer println("f2-begin")
	f3()
	defer println("f2-end")
}

func f3() {
	defer println("f3-begin")
	panic(0)
	defer println("f3-end")
}

func main() {
	//f1()
	//a()
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	return
}

func a() int {
	i := 0
	defer fmt.Println(i)
	i++
	return i
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
