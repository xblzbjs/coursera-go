/* Write a program which prompts the user to enter a floating point number and prints the integer
   which is a truncated version of the floating point number that was entered.
   Truncation is the process of removing the digits to the right of the decimal place.
*/

package main

import (
	"fmt"
	"math"
)

var firstFloatNumber float64

func main() {

	fmt.Println("Please enter a floating point number: ")
	num, err := fmt.Scanln(&firstFloatNumber)
	if err != nil || num != 1 {
		panic(err)
	}
	interNumber := int(math.Floor(firstFloatNumber))
	fmt.Println("Print the integer,which is a truncated version of the floating point number that was entered:", interNumber)

}
