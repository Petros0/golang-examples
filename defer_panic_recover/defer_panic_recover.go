package main

import (
	"log"
	"fmt"
	// "net/http"
)

func main(){
	fmt.Println("PAAANIC!")

	fmt.Println("start0")
	defer fmt.Println("middle0") //defer, execute after the main function, before it returns.
	fmt.Println("end0")

	//basically defer is for resources on the end.

	a := "start1"
	defer fmt.Println(a) //it will print "start" and not the "end"
	a = "end1"
/* 

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Hey go!"))
	})

	err := http.ListenAndServe(":8080",nil)
	
	if err != nil {
		panic(err.Error()) //a use of the panic, where an error occurs.
	}
*/

	fmt.Println("start2")
	defer func(){
		if err := recover(); err != nil{
			log.Println("Error...",err)
		}
	}()
	//commented because we want the next code to be executed
	//panic("An error occureeed..") //the defer function, executes before the panic action.



	fmt.Println("func start")
	panicker()
	fmt.Println("func end")
}

func panicker(){ // the main function will continue working.
	fmt.Println("about to panic")
	defer func(){
		if err := recover(); err != nil{
			log.Println("Error...",err)
			// panic(err) -> but if want not to continue we have to panic again ...
		}
	}()
	panic("An error occureeed..") 
	fmt.Println("done panicking....")
}