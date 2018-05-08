package main

import "fmt"

func main(){

	var a int = 42
	var b *int = &a
	fmt.Println(a,*b)
	a = 27
	*b = 99
	fmt.Println(a,*b)

	array := [3]int { 1, 2 ,3}
	value1 := &array[0]
	value2 := &array[1]

	fmt.Printf("%v %p %p\n",array, value1, value2)

	var ms *myStruct
	ms = &myStruct{foo :42}
	fmt.Println(ms)
	
	var wut *myStruct
	wut = new(myStruct)
	fmt.Println(wut)

	

}

type myStruct struct {
	foo int
}