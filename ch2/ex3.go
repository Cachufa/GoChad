package main

import "fmt"

func main() {
	var b byte = 255
	var smallI int32 = 2147483647
	var bigI uint32 = 4294967295
	b += 1
	smallI += 1
	bigI += 1
	fmt.Println(b, smallI, bigI)
}
