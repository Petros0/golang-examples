package main

import "fmt"
import "io"

func main(){
	fmt.Println("Hey INTERFACES!!")


	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))

	myInt := IntCounter(0)
	var inc Incrementer = &myInt

	for i:=0; i<10; i++{
		fmt.Println(inc.Increment())
	}

	r, ok := inc.(io.Reader) //it will fail, but like this we do a downcast. --> w.(A_TYPE)
	if ok{
		fmt.Println("converted", r)
	}else {
		fmt.Println("nope.")
	}

	var myObj interface{} = ConsoleWriter{}
	if r, ok := myObj.(ConsoleWriter); ok{
		fmt.Println(r)
		r.Write([]byte("Hello Go!"))
	}else{
		fmt.Println("conversion failed..")
	}

	var i interface{} = 0
	switch i.(type){
	case int:
		fmt.Println("its an int!")
	default:
		fmt.Println("something else!")
	} 
}

type Writer interface{
   	Write([]byte) (int, error)
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int,error){
	n, err := fmt.Println(string(data))
	return n,err
}

type Incrementer interface{
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int{
	*ic++
	return int(*ic)
}