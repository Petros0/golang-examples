package main

import (
	"fmt"
)

func main(){

	fmt.Println("Hey Control Flow!")

	if true{
		fmt.Println("I AM TRUE")
	}

	statePopulations := map[string]int{
		"California": 3943101,
		"Texas": 43242,
	}
	//the pop variable can be used only inside the if statement
	if pop, ok := statePopulations["Texas"]; ok{
		fmt.Println(pop)
	}

	number := 50
	guess := 10

	if guess < number{
		fmt.Println("IT'S SMALLER!")
	}

	fmt.Println(number<=guess, number>=guess, number != guess)

	// operators are like normal: || && !

	if returnTrue() && number == 50{
		fmt.Println("WORKS!")
	}

	//switch statement

	switch 2{
	case 1:
		fmt.Println("1")
	case 2,3,4: //checks for these values it's like : 
	//case 2:
	//case 3:
	//case 4:	 
		fmt.Println("2")
	default:
		fmt.Println("nothing")
	}

	i := 10

	switch{
	case i<=10:
			fmt.Println("i<=10")
			fallthrough //the next is executing even if the case is false
	case i<=20:
			fmt.Println("i<=20")
	default:
		fmt.Println("whatever")
	}

	var j interface{} = 1

	switch j.(type){
	case int:
		fmt.Println("IT's an int!")
		break
		fmt.Println("will not be showed")
	case float64:
		fmt.Println("FLOAT!")
	default:
		fmt.Println("whatever")
	}	
}

func returnTrue() bool {
	fmt.Println("returning true")
	return true
}