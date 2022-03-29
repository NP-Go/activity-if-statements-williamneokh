package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

/*1.Given the numbers is 17 and 24 and using preceding conditionals
for at least 1 of the numbers.a.Which number is bigger?b.By how much is it bigger by?*/
func main() {
	var number1 int
	var number2 int
	fmt.Println("Please key in the first number")
	_, _ = fmt.Scanln(&number1)
	fmt.Println("Please key in the second number.")
	_, _ = fmt.Scanln(&number2)

	//Insert your code here
	message, err := numberCheck(number1, number2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
	//Hint: You may wish to make use of strconv.Itoa to convert int to string

}

func numberCheck(number1, number2 int) (string, error) {
	if number1 == number2 {
		message := "First number " + strconv.Itoa(number1) + " is equal to second number " + strconv.Itoa(number2)
		return message, nil
	}
	if number1 > number2 {
		diff := number1 - number2
		message := "First number " + strconv.Itoa(number1) + " is greater than second number " + strconv.Itoa(number2) + " by " + strconv.Itoa(diff)
		return message, nil
	}
	if number1 < number2 {
		diff := number2 - number1
		message := "First number " + strconv.Itoa(number1) + " is smaller than second number " + strconv.Itoa(number2) + " by " + strconv.Itoa(diff)
		return message, nil
	} else {
		return "", errors.New("something went wrong")
	}
}
