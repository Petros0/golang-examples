package main

import "fmt"

func main(){
	sayMessage("Hey Functions!!")
	sayGreeting("Hey","Petros")

	greeting := "Hello"
	name := "Stacey"
	sayGreetingWithPointers(&greeting,&name)
	result := sum(1,2,3,4,5)
	fmt.Println(result)

	wut := returnAPointer()
	fmt.Println(*wut)

	//multiple value functions
	d, err:= divide(3.1,0.0)

	if err != nil{
		fmt.Println(err)
		// return
	}
	fmt.Println(d)

	//anonymous func

	func(){
		fmt.Println("Hey GO")
	}()

	var f func() = func(){
		fmt.Println("Hey")
	}

	f()

	g := Greeter{
		greeting: "Hello",
		name: "Go",
	}
	g.greet()
}


type Greeter struct{
	greeting string
	name string
}

//methods
func (g Greeter) greet(){ // if "Greeter" has an '*', then we can manipulate it (passes the reference)
	fmt.Println(g.greeting,g.name)
}

func divide(a, b float64) (float64, error){
	if b== 0.0{
		return 0.0, fmt.Errorf("Cannot devide by zero")
	}
	return a/b, nil
}

func returnAPointer() *int{
	var wut int = 3
	return &wut

}

func sum(values ...int) int{
	result := 0
	for _,v := range values {
		result += v
	}

	fmt.Println("sum is: ",result)
	return result
}

func sayGreetingWithPointers(greeting, name *string){
	fmt.Println(*greeting,*name)
	*name = "Petros"
}
func sayGreeting(greeting, name string){
	fmt.Println(greeting,name)
}

func sayMessage(msg string){
	fmt.Println(msg)
}
