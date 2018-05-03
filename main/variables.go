package main

import (
	"fmt"
	test_variables "../test_variables"
)

var i = 3

var j = "Visible only in package." //lowercase variables are only package visible
var J = "Visible out of package too." //uppercase variables are visible out of package too


func main() {
	//scope
	fmt.Println(i)
	i := 42
	fmt.Println(i)

	fmt.Println(j)
	//global variable
	fmt.Println(test_variables.A_GLOBAL_VARIABLE)

	//naming
	var theHTTP = "Keep acronyms capitals"
	fmt.Println(theHTTP)


}
