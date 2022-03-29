package main

import (
	"errors"
	"fmt"
	"log"
)

/*1.Given the numbers is 17 and 24 and using preceding conditionals
for at least 1 of the numbers.a.Which number is bigger?b.By how much is it bigger by?*/
func main() {
	var number1 int
	var number2 int
	fmt.Println("Please key in first number")
	fmt.Scanln(&number1)
	fmt.Println("Please key in second number.")
	fmt.Scanln(&number2)

	//Insert your code here
	err := numberCheck(number1, number2)
	if err != nil {
		log.Fatal(err)
	}

	//Hint: You may wish to make use of strconv.Itoa to convert int to string

}

func numberCheck(number1, number2 int) error {
	if number1 == number2 {
		fmt.Printf("First number %v is equal to second number %v", number1, number2)
		return nil
	}
	if number1 > number2 {
		diff := number1 - number2
		fmt.Printf("First number %v is greater than %v by %v", number1, number2, diff)
		return nil
	}
	if number1 < number2 {
		diff := number2 - number1
		fmt.Printf("First number %v is smaller than %v by %v", number1, number2, diff)
		return nil
	} else {
		return errors.New("Something went wrong!")
	}
}
