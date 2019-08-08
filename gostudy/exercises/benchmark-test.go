package main

import (
	"fmt"
	"testing"
)

func BenchamarkPrint(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fmt.Println('a')
	}
}
