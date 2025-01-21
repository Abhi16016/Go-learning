package main

import "fmt"

//NOTE: Go does not have ternary operator, instead we have to use if-else for it.

func main() {

	//standard if-else
	age := 22

	if age >= 18 {
		fmt.Println("Person is an adult")
	} else if age >= 12 {
		fmt.Println("Person is teenager")
	} else {
		fmt.Println("Person is a kid")
	}

	//example-1
	var role = "admin"
	var harPermission = true

	if role == "admin" && harPermission {
		fmt.Println("YES")
	}

	//example-2 (We can declare variables inside "if" construct)
	if age := 16; age >= 18 {
		fmt.Println("Person is an adult", age)
	} else if age >= 12 {
		fmt.Println("Is teen")
	}
}
