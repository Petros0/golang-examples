package main

import "fmt"


//enumerated constants

const (
	_ = iota
	a 
	b
	c
)

const(
	a2 = iota
)

const(
	_ = iota //ignore the first value 
	KB = 1 << (10 * iota)
	MB
	GB
	TB
)

func main(){

	//name conventions
	const myConst int = 42
	fmt.Printf("%v, %T\n",myConst,myConst)

	fmt.Printf("%v, %T\n",a,a)
	fmt.Printf("%v, %T\n",b,b)
	fmt.Printf("%v, %T\n",c,c)
	fmt.Printf("%v, %T\n",a2,a2)


	fileSize := 1024.
	fmt.Printf("%.2fKB",fileSize/KB)
	
}