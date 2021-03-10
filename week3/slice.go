/*
Write a program which prompts the user to enter integers and stores the integers in a sorted slice.
The program should be written as a loop. Before entering the loop, the program should create an empty integer slice of size (length)
3. During each pass through the loop, the program prompts the user to enter an integer to be added to the slice.
The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order.
The slice must grow in size to accommodate any number of integers which the user decides to enter.
The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.
*/

package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
)

func main() {
	intSli := make([]int, 3)
	for {
		var input string
		fmt.Println("Please enter integers(enter 'X' to exit):")
		fmt.Scanf("%s\n", &input)
		if input == "X" {
			break
		}
		fmt.Println("Your number is ", input)
		num, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		intSli = append(intSli, num)
		fmt.Println("Before sort:", intSli)
		sort.Ints(intSli)
		fmt.Println("After sort: ", intSli)
	}

}
