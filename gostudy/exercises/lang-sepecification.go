package main

import "fmt"

// type S struct{}

// func (s S) F() {}

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

// func InitIfacePointer() IF {
// 	var s *S
// 	return s
// }

// func main() {
// 	//println(InitType() == nil)
// 	// println(InitPointer() == nil)
// 	//println(InitEfaceType() == nil)
// 	//println(InitEfacePointer() == nil)
// 	//println(InitIfaceType() == nil)
// 	//println(reflect.DeepEqual(InitIfaceType(), S{}))
// 	//println(InitIfacePointer() == nil)
// }

//recover() start
// func f() {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			fmt.Printf("recover:%#v \n", r)
// 			log.Printf("recover:%#v \n", r)
// 		}
// 	}()
// 	panic(1)
// 	panic(2)
// }

// func main() {
// 	f()
// 	fmt.Printf("abc %d \n", 123)
// }

//recover() end

//goroutine closures start
// const N = 10

// func main() {
// 	var mn sync.Mutex
// 	m := make(map[int]int)
// 	wg := &sync.WaitGroup{}
// 	wg.Add(N)
// 	for i := 0; i < N; i++ {
// 		go func(i int) {
// 			defer wg.Done()
// 			mn.Lock()
// 			m[i] = i
// 			mn.Unlock()
// 		}(i)
// 	}
// 	wg.Wait()
// 	fmt.Println(len(m))
// }

//goroutine closures end

//type shadowing start
// type S1 struct{}

// func (s1 S1) f() {
// 	fmt.Println("s1.f()")
// }

// func (s1 S1) g() {
// 	fmt.Println("s1.g()")
// }

// type S2 struct {
// 	S1
// }

// func (s2 S2) f() {
// 	fmt.Println("s2.f()")
// }

// type I interface {
// 	f()
// }

// func printType(i I) {
// 	fmt.Printf("%T\n", i)
// 	if s1, ok := i.(S1); ok {
// 		s1.f()
// 		s1.g()
// 	}
// 	if s2, ok := i.(S2); ok {
// 		s2.f()
// 		s2.g()
// 	}
// }
// func main() {
// 	printType(S1{})
// 	printType(S2{})
// }

// //type shadowing end

//string to bytes start
// func main() {
// 	s := "123"
// 	ps := &s
// 	b := []byte(*ps)
// 	pb := &b

// 	s += "4"
// 	*ps += "5"
// 	b[1] = '0'

// 	fmt.Println(*ps)
// 	fmt.Println(string(*pb))
// }

//string to bytes end

//map mutex start
// const N = 10

// func main() {
// 	m := make(map[int]int)
// 	wg := &sync.WaitGroup{}
// 	sm := &sync.Mutex{}
// 	wg.Add(N)
// 	for i := 0; i < N; i++ {
// 		go func() {
// 			defer wg.Done()
// 			sm.Lock()
// 			m[rand.Int()] = rand.Int()
// 			sm.Unlock()
// 		}()
// 	}
// 	wg.Wait()
// 	fmt.Println(len(m))
// }

//map mutex end

//deepEqual start
// type S struct {
// 	a, b, c string
// }

// func main() {
// 	x := interface{}(&S{"a", "b", "c"})
// 	y := interface{}(&S{"a", "b", "c"})

// 	fmt.Println(reflect.DeepEqual(x, y))
// }

//deepEqual end

//gosched start
// const N = 26

// func main() {
// 	const GOMAXPROCS = 1
// 	runtime.GOMAXPROCS(GOMAXPROCS)

// 	var wg sync.WaitGroup
// 	wg.Add(2 * N)
// 	for i := 0; i < N; i++ {
// 		go func(i int) {
// 			defer wg.Done()
// 			runtime.Gosched()
// 			fmt.Printf("%c\n", 'a'+i)
// 		}(i)

// 		go func(i int) {
// 			defer wg.Done()
// 			fmt.Printf("%c\n", 'A'+i)
// 		}(i)
// 	}
// 	wg.Wait()
// }

//gosched end

// map immutability start
// type S struct {
// 	name string
// }

// func main() {
// 	m := map[string]*S{"x": &S{"one"}}
// 	m["x"].name = "two"
// 	fmt.Printf("the map %s", m)
// }

// map immutability end

//slice sorting start
// type S struct {
// 	v int
// }

// func main() {
// 	s := []S{{1}, {3}, {5}, {2}}
// 	sort.Slice(s, func(i, j int) bool {
// 		return s[i].v < s[j].v
// 	})
// 	fmt.Printf("%#v", s)
// }

//slice sorting end

// func main() {
// 	ret := make(map[string]string, 0)
// 	//mapniltest(ret, map[string]string{"name": "user"})
// 	mapniltest(ret, nil)
// 	fmt.Printf("the ret %#v", ret)
// }

// func mapniltest(ret, setm map[string]string) {
// 	if setm == nil {
// 		println("xxx")
// 	}
// 	for k, v := range setm {
// 		ret[k] = v
// 	}
// }

type T int

func IsClosed(ch <-chan T) bool {
	select {
	case <-ch:
		return true
	default:

	}
	return false
}

func main() {
	c := make(chan T)
	fmt.Println(IsClosed(c))
	close(c)
	fmt.Println(IsClosed(c))
}
