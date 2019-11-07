package other

import "fmt"

func add(x int, y int) int {
	return x + y
}

func nouse() {
	var no string
	_ = no
}

func empty() {
}

func vary(args ...int) {
	fmt.Println(args)
}

func fix_vary(fix int, args ...int) {
	fmt.Println(args)
}

func swap(x, y int) (int, int) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum / 5
	y = sum % 5
	return
}

// f := func (x, y int) int {
//     return x + y
// }(1,2)

// func (x, y int)int {
//     return x+y
// }(2,3)

// f := func(x, y int) int {
//      return x+y
// }

// if i:=j;i>0 {
//     // do sth
// }

// switch i {
//     case 0:
//         //
//     case 1:
//         //
//     default:
//         //
// }

// switch{
//     case char == 'a':
//         //
//     case char == 'b':
//         //
// }

// switch 1 {
//     case 1,2,3:
//         //
//     default:
//         //
// }

// switch os:=runtime.GOOS;os{
//     case "darwin":
//     case "linux":
// }

// for {}

// a := []int{1,5,3,4,2,6}
// for i, j:=0,len(a) -1; i < j; i, j = i +1, j -1  {
//     a[i],a[j] = a[j],a[i]
// }
// fmt.Println(a)

// for index, value := range(alist) {}

// for key,value := range(amap) {}

// f, err := os.Open(filename)
// if err != nil {
//     log.Println("")
//     return
// }

// defer f.Close()
// //

// type Myint int
// func (m Myint)Less(b Myint) bool {
//     return m < b
// }

// var a Myint = 1
// if a.Less(2) {
//     //do sth
// }

// type Person struct{
//     Name string
//     Age int
// }

// func (p Person)GetName()string{
//     return p.Name
// }

// func (p *Person)SetName(name string){
//     p.Name = name
// }

// type Base struct{
//     Name string
// }
// func (b *Base) Name()string{}

// type Foo struct{
//     Base
// }

// var foo Foo
// foo.Name()

// type IFile interface {
//     Read(buf []byte) (n int, err error)
//     Write(buf []byte) (n int, err error)
//     Seek(off int64, whence int) (pos int64, err error)
//     Close() error
// }

// type IReader interface {
//     Read(buf []byte) (n int, err error)
// }

// type IWriter interface {
//     Write(buf []byte) (n int, err error)
// }

// type ICloser interface {
//     Close() error
// }

// type File struct {
// }

// func (f *File) Read(buf []byte) (n int, err error)
// func (f *File) Write(buf []byte) (n int, error)
// func (f *File) Seek(off int64, whence int) (pos int64, err error)
// func (f *File) Close() error

// var file1 IFile = new(File)
// var file2 IReader = new(File)
// var file3 IWriter = new(File)
// var file4 ICloser = new(File)

// var file1 Writer = new(file)
// if file5, ok := file1.(two.IStream); ok {
// }

// type ReadWriter interface {
//     IReader
//     IWriter
// }

// type ReadWriter interface {
//     Read(p []byte) (n int, err error)
//     Write(p []byte) (n int, err error)
// }

// {
//     var v1 interface{} = 1
//     var v2 interface{} = "abc"
//     var v3 interface{} = &v2
//     var v4 interface{} = struct {X int} {1}
//     var v5 interface{} = &struct{X int}{1}
// }

// var v1 interface{} = ...
// switch v :=ve.(type) {
//     case int:
//     //
//     case string:
//     //
// }

// tes := make(map[string]interface{})
// tes['a'] = "abc"
// tes['b'] = 123

// value, ok := tes["a"].(string)
// value2, ok2 := tes["a"].(int)

// var var1 int = 7
// var2 := float64(var1)
// var3 := int64(var1)

// var4 := new(int32)
// var5 := (*int32)(var4)

// package main
// import "fmt"
// func Add(x, y int) {
// z := x+y
// fmt.Println(z)
// }

// func main() {
//     for i := i<10;i++ {
//         go Add(i, i)
//     }
// }

// time.Sleep()

// package main
// import "fmt"
// import "sync"
// import "runtime"

// var counter int = 0
// func Count(lock *sync.Mutex) {
//     lock.Lock()
//     counter++
//     fmt.Println(counter)
//     lock.Unlock()
// }

// func main() {
//       lock := &sync.Mutex{}
//       for i := 0; i < 10; i++ {
//           go Count(lock)
//       }
//       for {
//           lock.Lock()
//           c := counter
//           lock.Unlock()
//           runtime.Gosched()
//           if c >= 10 {
//               break
//           }
//       }
// }

// //go 思路
// package main
// import "fmt"
// func Count(ch chan int) {
//     ch <- 1
//     fmt.Println("Counting")
// }

// function main() {
// chs :=make([]chan int 10)
// for i := 0; i < 10; i++ {
//     chs[i] = make(chan int)
//     go Count(chs[i])
// }
// for _, ch := range(chs) {
//     <-ch
// }
// }

// var ch1 chan int
// var ch2 chan<- float64 //写float64数据
// var ch3 <-chan int //读int数据

// ch := make(chan int)

// go func() {
//     ch <- 1 //写入数据
// }

// i := <-ch //取数据
// <- ch   //取数据不使用 将死锁

// x, ok = <- ch //ok 表示有咩有数据

// chbuf := make(chan int, 100) //没有数据时阻塞读取  数据塞满100歌时阻塞写入

// //单向通道一般用于函数参数定义

// func single_r(ch <-chan int) {
//     //do sth 只读取数据
// }

// func single_w(ch chan<- int) {
//     //do sth 只能写入数据
// }

// var ch chan int //读写双向通道
// single_r(ch)
// single_w(ch)

// LOOP:
//     for {
//         select {
//             case <-chan1:
//                 //<-chan1的处理过程
//             case chan2 <- 1:
//                 break LOOP
//             default:
//              <-time.After(duration)
//         }
//     }

// for i:= range(ch) {
//     //do
// }

// //信号量处理
// chan <- struct{}{}

// package main

// /*
// #include <stdio.h>
// void hello() {
//     printf("Hello,Cgo")
// }
// import "C"

// func Hello() {
// C.hello()
// }

// func main() {
//  Hello()
// }

// C.char()
// C.schar()
// C.uchar()
// C.short()
// C.ushort()
// C.int()
// C.uint()
// C.long()
// C.ulong()
// C.longlong()
// C.ulonglong()
// C.float()
// C.double()

// defer func() {
//     if err := recover(); err != nil {
//     }
// }

// panic("something error")
