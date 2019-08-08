package main

//go build -buildmode=plugin -o str.so str.go
import "strings"

func UpperCase(s string) string {
	return strings.ToUpper(s)
}
