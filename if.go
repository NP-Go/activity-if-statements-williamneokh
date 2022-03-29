package main

import "fmt"

/*1.Given the numbers is 17 and 24 and using preceding conditionals
for at least 1 of the numbers.a.Which number is bigger?b.By how much is it bigger by?*/
func main() {
	var number1 int
	var number2 int
	resultMessage := "No outcome due to first input is greater then 2nd input."
	fmt.Println("Please input 2 numbers, 1 greater and 1 smaller.\nPlease input a smaller number")
	fmt.Scanln(&number1)
	fmt.Println("Please key in a larger number.")
	fmt.Scanln(&number2)

	//Insert your code here
	if number1 < number2 {
		diff := number2 - number1

		fmt.Printf("%v is greater than %v and is by: %v", number2, number1, diff)

	} else {
		fmt.Println(resultMessage)
	}
	//Hint: You may wish to make use of strconv.Itoa to convert int to string

}
