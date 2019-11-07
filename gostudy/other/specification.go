package other

import "fmt"

func specification() {
	s := []int{1, 2, 3}
	ss := s[1:]
	ss = append(ss, 4)
	for _, v := range ss {
		v += 10
	}
	for i := range ss {
		ss[i] += 10
	}
	fmt.Printf("s %v", s)

}

func f() (x, y int) {
	return
}
