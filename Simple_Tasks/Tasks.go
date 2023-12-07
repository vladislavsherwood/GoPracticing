package main

import "fmt"

func main() {
	a := 17
	b := 13

	sum := SumOfTwoNumbers(a, b)
	fmt.Printf("Sum of Two Numbers: %v\n", sum)

	maximum := MaximumOfTwoNumbers(a, b)
	fmt.Printf("Maximum of Two Numbers: %v\n", maximum)

	even := EvenNumber(a)
	fmt.Printf("Number %v is even: %v\n", a, even)
}

// Task 1: Sum of Two Numbers
// Write a function that takes two arguments and returns their sum.

func SumOfTwoNumbers(a, b int) int {
	return a + b
}

// Task 2: Maximum of Two Numbers
// Create a function that takes two numbers and returns the maximum of the two.

func MaximumOfTwoNumbers(a, b int) int {
	if a > b {
		return a
	} else if a < b {
		return b
	}
	return 0
}

// Task 3: Check for Even Number
// Implement a function that checks if a number is even.

func EvenNumber(a int) bool {
	if a%2 == 0 {
		return true
	}
	return false
}

// Task 4: Print Numbers from 1 to 100
// Write a program, using a loop, to print all numbers from 1 to 100.

// Task 5: Factorial of a Number
// Create a function that calculates the factorial of a number.

// Task 6: Palindrome Check
// Implement a function that checks if a string is a palindrome.

// Task 7: Average of an Array
// Write a program that finds the average value of an array of numbers.

// Task 8: Nth Fibonacci Number
// Create a function that returns the Nth Fibonacci number.

// Task 9: Least Common Divisor
// Implement a function that finds the least common divisor of two numbers.

// Task 10: Primality Check
// Write a program that checks if a number is prime.

// Task 11: Swap Variable Values
// Create a function that swaps the values of two variables.

// Task 12: Count Words in a String
// Implement a function that returns the number of words in a string.

// Task 13: Multiplication Table
// Write a program that prints the multiplication table for a given number N.

// Task 14: Remove Duplicates from an Array
// Create a function that removes duplicates from an array.

// Task 15: Sum of Array Elements
// Implement a function that returns the sum of the elements in an array.

// Task 16: Leap Year Check
// Write a program that checks if a year is a leap year.

// Task 17: Reverse a String
// Create a function that reverses a string.

// Task 18: Quadratic Equation Roots
// Implement a program that finds the roots of a quadratic equation.

// Task 19: Array Sorting
// Write a function that sorts an array of integers.

// Task 20: String Number Check
// Create a program that determines if a string is a number.