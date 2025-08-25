package main

import "fmt"

type EmployeeGoogle struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// Pascal case
	// eg. CalculateArea, UserInfo, NewHTTPRequest
	// Structs, interfaces, enums

	// snake_case
	// eg. user_id, first_name, http_request
	//

	// UPPERCASE
	// constants

	// mixedCase
	// eg. javaScript, htmlDoc, isValid
	// variables

	const MAXRETRIES = 5
	var employeeID = 1001
	fmt.Println(employeeID)

}
