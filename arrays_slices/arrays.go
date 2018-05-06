package main

import (
	"fmt"
)

func main() {

	fmt.Println("hey Arrays & Slices!\n#############")

	// arrays
	grades := [...]int{97, 30, 21} // '...' any length (only for) initializing hard an array
	fmt.Println(grades)

	var students [3]string
	students[0] = "Hello"
	students[1] = "Student World"
	students[2] = "!"
	fmt.Println(students)
	fmt.Println(students[1])
	fmt.Printf("Length:%v\n", len(students))

	//2 dimensional
	var identityMatrix [3][3]int = [3][3]int{
		[3]int{1,0,0},
		[3]int{0,1,0},
		[3]int{0,0,1},
	}

	fmt.Println(identityMatrix)

	aArray := [...]int{1,2,3}
	//no reference!!! for reference should -> '&a'
	bArray := aArray // copy array a to b
	bArray[1] = 5 //change value, !!changes the value only to the b array

	fmt.Println(aArray)
	fmt.Println(bArray)

	//slices

	aSlice := []int{1,2,3}
	bSlice := aSlice // slices are only with references, so 
	bSlice[1] = 99 //changes the values also in 'aSlice'
	fmt.Println(aSlice)
	fmt.Println(bSlice)
	fmt.Printf("Length: %v\n", len(aSlice))
	fmt.Printf("Capacity: %v\n", cap(aSlice))

	//ways to create slices
	
	a := []int{1,2,3,4,5,6,7,8,9,10}
	b := a[:] //slice of all elemetns
	c := a[3:] //slice from 4th element to the end
	d := a[:6] //slice first 6 elements
	e := a[3:6] //slice the 4th, 5th, 6th elements

	//the value still changes.
	//e.g.: b[1] = 99 then a[1] == 99

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	//make function

	aMake := make([]int, 3, 100) //3 length && 100 capacity
	fmt.Println(aMake)

	//append function
	aGreatSlice:= []int{}
	fmt.Println(aGreatSlice)
	fmt.Printf("Length: %v\n", len(aGreatSlice))
	fmt.Printf("Capacity: %v\n", cap(aGreatSlice))

	//appends to the slice 'aGreatSlice' the values '1','2','3','4','5' and returns the new array
	
	aGreatSlice = append(aGreatSlice,1,2,3,4,5) 
	fmt.Println(aGreatSlice)
	fmt.Printf("Length: %v\n", len(aGreatSlice))
	fmt.Printf("Capacity: %v\n", cap(aGreatSlice))

	//concatanate slices
	concSlice := []int{}
	concSlice = append(concSlice,1)
	fmt.Println(concSlice)

	// appends all the values from the argument slice
	// we need the 3 'dots', otherwise GO will tell us
	// that it's not the same type. -- makes sense, right?
	concSlice = append(concSlice, []int{2,3,4,5}...) 
	fmt.Println(concSlice)

	//stack operations

	stackA := []int{1,2,3,4,5}
	stackB := stackA[1:] // remove the first element
	fmt.Println(stackB)

	stackA = []int{1,2,3,4,5}
	stackB = stackA[:len(stackA)-1] // remove the last element
	fmt.Println(stackB)

	stackA = []int{1,2,3,4,5}
	stackB = append(stackA[:2], stackA[3:]...) // remove the third element
	fmt.Println(stackB)

}
