package main

import "fmt"

func main() {
	// array := [6]int{1: 30, 3: 3, 50}
	// fmt.Println("modify before:", array)
	// for k, v := range array {
	// 	if k == 3 {
	// 		array[k] = 55555
	// 	}
	// 	fmt.Println(k, v)
	// }
	// fmt.Println("modify after:", array)

	//指针
	// array := [5]*int{0: new(int), 1: new(int)}
	// *array[0] = 10
	// *array[1] = 20
	// fmt.Println(*array[0])

	// var array1 [5]string
	// array2 := [5]string{"red", "blue", "green", "yellow", "pink"}
	// array1 = array2
	// fmt.Println(array1 == array2)

	// var array1 [3]*string

	// array2 := [3]*string{new(string), new(string), new(string)}

	// *array2[0] = "red"
	// *array2[1] = "blue"
	// *array2[2] = "green"

	// array1 = array2
	// fmt.Println(array1)
	// fmt.Println(array2)
	// fmt.Println(array1 == array2)
	// fmt.Println(array1[0])
	// fmt.Println(array2[0])
	// fmt.Println(*array1[0])
	// fmt.Println(*array2[0])

	//var array [4][2]int
	// array := [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	// fmt.Println(array)
	// array = [4][2]int{1: {20, 21}}
	// fmt.Println(array)
	// array = [4][2]int{1: {0: 20}, 3: {1, 41}}
	// fmt.Println(array)
	// var array [1e6]int
	// printArr(&array)

	// slice := make([]string, 3, 5)
	// fmt.Println(len(slice))

	// 创建nil切片
	// var slice2 []int
	// fmt.Println(slice2)

	// slice := make([]int, 0)
	// fmt.Println(slice)

	// slice := []int{10, 20, 30, 40, 50}
	// slice[1] = 25
	// fmt.Println(slice)
	// newSlice := slice[1:3]
	// newSlice[0] = 1
	// newSlice = append(newSlice, 45)
	// newSlice = append(newSlice, 55)
	// newSlice = append(newSlice, 65)
	// newSlice = append(newSlice, 75)
	// newSlice[0] = 22
	// fmt.Println("new slice cap:", cap(newSlice))
	// fmt.Println(newSlice)
	// fmt.Println(slice)
	// source := []string{"red", "blue", "green", "yellow", "pink"}
	// slice := source[2:3:4]
	// fmt.Println("slice cap:", cap(slice))
	// fmt.Println("slice len:", len(slice))
	// fmt.Println("slice:", slice)

	// source := []string{"apple", "orange", "plum", "banana", "grape"}
	// slice := source[2:3:3]
	// slice[0] = "xxx"
	// // slice = append(slice, "kiwi")
	// // slice = append(slice, "gr")
	// // slice = append(slice, "gb")
	// // slice = append(slice, "g3")
	// fmt.Println("source:", source)
	// fmt.Println("slice:", slice)

	// slice := []int{10, 20, 30, 40}

	// for index, value := range slice {
	// 	fmt.Printf("value:%d value-addr: %X elementAddr: %X\n", value, &value, &slice[index])
	// }

	// 创建一个映射，存储颜色以及颜色对应的十六进制代码
	colors := map[string]string{
		"AliceBlue":   "#f0f8ff",
		"Coral":       "#ff7F50",
		"DarkGray":    "#a9a9a9",
		"ForestGreen": "#228b22",
	}
	// 显示映射里的所有颜色
	for key, value := range colors {
		fmt.Printf("Key: %s Value: %s\n", key, value)
	}
	fmt.Println("--------------------")
	// 调用函数来移除指定的键
	removeColor(colors, "Coral")
	// 显示映射里的所有颜色
	for key, value := range colors {
		fmt.Printf("Key: %s Value: %s\n", key, value)
	}

}

func removeColor(colors map[string]string, key string) {
	delete(colors, key)
}

func printArr(array *[1e6]int) {
	fmt.Println(array)
}
