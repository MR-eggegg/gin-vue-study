package main

import (
	"unsafe"
	"fmt"
)


//go:noescape
func _benchmark(buf unsafe.Pointer)

func benchmark() {
var d0  uint64 = 0
var d1  uint64 = 0
var d2  uint64 = 0
var d3  uint64 = 0
var d4  uint64 = 0
var d5  uint64 = 0
var d6  uint64 = 0
var d7  uint64 = 0
var d8  uint64 = 0
var d9  uint64 = 0
var d10 uint64 = 0
var d11 uint64 = 0

a := []unsafe.Pointer{
unsafe.Pointer(&d0),unsafe.Pointer(&d1),unsafe.Pointer(&d2),
unsafe.Pointer(&d3),unsafe.Pointer(&d4),unsafe.Pointer(&d5),
unsafe.Pointer(&d6),unsafe.Pointer(&d7),unsafe.Pointer(&d8),
unsafe.Pointer(&d9),unsafe.Pointer(&d10),unsafe.Pointer(&d11)}

p1 := unsafe.Pointer(&a[0])
// emit
_benchmark(p1)
// output
fmt.Println("-----------------------------")
fmt.Printf("a:%d \t b:%d \t c:%d\n", d0, d1, d2);
fmt.Printf("a:%d \t b:%d \t c:%d\n", d3, d4, d5);
fmt.Printf("a:%d \t b:%d \t c:%d\n", d6, d7, d8);
fmt.Printf("a:%d \t b:%d \t c:%d\n", d9, d10, d11);

}


func main(){
benchmark()
}
