package main


import (
	"fmt"
)

func main(){

	//bools 
	var n bool 
	fmt.Printf("%v \n",n)

	//ints 
	anInt := 42
	fmt.Printf("%v, %T\n",anInt, anInt)

	var uanInt uint16 = 42
	fmt.Printf("%v, %T\n",uanInt, uanInt)
	

	var aint int = 10
	var bint int8 = 3
	fmt.Println(aint+int(bint))

	//floats
	afloat := 3.14
	afloat = 2.1E14
	fmt.Printf("%v, %T\n",afloat, afloat)

	//complexes
	var acomplex64 = 5i +1
	fmt.Printf("%v, %T\n",acomplex64, acomplex64)
	fmt.Printf("%v, %T\n",real(acomplex64), real(acomplex64))
	fmt.Printf("%v, %T\n",imag(acomplex64), imag(acomplex64))

	//strings
	s := "this is a string"
	fmt.Printf("%v, %T\n",string(s[2]), s[2])

	s2 := "second string..."
	fmt.Printf("%v, %T\n",s+s2, s[2])

	bytes := []byte(s)
	fmt.Printf("%v, %T\n",bytes,bytes)

	var r rune = 'a'
	fmt.Printf("%v, %T\n",r,r)


}