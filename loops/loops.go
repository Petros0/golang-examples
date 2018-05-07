package main

import (
	"fmt"
)

func main(){
	fmt.Println("Hey Loops!")
	
	for i := 0; i< 5; i++{
		fmt.Println(i)
	}

	for i, j :=0, 0; i<5; i, j = i+1, j+2{
		fmt.Println(i,j)
	}

	i:=0
	for ; i<5; i++{
		fmt.Println(i)
	}
	j:=0
	for j<5{ //while
			fmt.Println(j)
			j++
	}

	i=0
	for{
		fmt.Println(i)
		i++
		if i==5{
			break
		}
	}


	for i := 0; i < 10; i++ {
		if(i%2 == 0){
			continue
		}
		fmt.Println(i)
	}
//a tag, that if we say 'break Loop' it will break also the outer loop.
Loopbreaker:
	for i := 0; i <4; i++ {
		for j := 0; j < 3 ; j++ {
			fmt.Println(i * j)
			if i*j >=3{
				// with a 'break' only we exit only the inner loop
				// but if we add a "tag" -- such as Loop -- it will break also the outer loop.
				fmt.Println("BREAKING -- LAST STATEMENT OF THE TWO LOOPS")
				break Loopbreaker 
				
			}
		}
	}


	//loop through lists.
	fmt.Println("\nLOOP THROUGH LISTS")

	s := []int{1,2,3}

	for index, value := range s {
		fmt.Println(index, value)
	}

	aString := "Hello GO!"

	for index, value := range aString {
		fmt.Println(index, string((value)))
	}
}