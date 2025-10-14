package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, Frontend Masters Go!")

	// Variables
	var name string = "Olivier"
	// prints according to format specifier
	/*
		Format Specifiers
		-----------------
			%v for value (default format)
			%#v for Go-syntax representation of the value
			%T for type of the value
			%% for a literal percent sign
			%t for boolean
			%c for character (rune)
			%s for string
			%d for integer
			%b for binary representation of an integer
			%x for hexadecimal representation of an integer
			%f for floating-point number
			%e for scientific notation of a floating-point number
			%g for the most compact representation of a floating-point number (either %f or %e)
			%p for pointer address
	*/

	fmt.Printf("Name: %s\n", name)

	// type inferencing:
	age := 32
	fmt.Printf("Age: %d\n", age)

	// Constants
	const isCool = true
	fmt.Printf("isCool: %t\n", isCool)

	// Multiple variables
	var (
		firstName = "Olivier"
		lastName  = "Coq"
	)
	// Multiple variable declaration and assignment, order matters:
	fmt.Printf("Full Name: %s %s\n", firstName, lastName)

	var city string
	city = "Tampa"
	fmt.Printf("City: %s\n", city)

	var (
		isEmployed = true
		salary     = 61000
		position   = "Full-Stack Developer"
	)

	fmt.Printf("Employed: %t\nSalary: %d\nPosition: %s\n", isEmployed, salary, position)

	/*
	 Zero values. If you don't initialize a variable, it gets the zero value of its type
	 int = 0, float = 0.0, bool = false, string = ""
	 Useful when you want to declare a variable but don't have an initial value for it yet
	 Avoids "undefined variable" errors
	*/
	var height int
	var weight float64
	var isMarried bool
	var middleName string

	fmt.Printf("Height: %d\nWeight: %f\nMarried: %t\nMiddle Name: %s\n", height, weight, isMarried, middleName)

	// Constants:
	const pi = 3.14
	fmt.Printf("Pi: %f\n", pi)

	// Uncommenting the next line will cause a compile-time error
	// pi = 3.14159 // Error: cannot assign to pi (constant)

	// enums using iota
	const (
		Jan int = iota + 1 // iota starts at 0, so Jan is 1
		Feb
		Mar
		Apr
	)

	fmt.Printf("Months: Jan=%d, Feb=%d, Mar=%d, Apr=%d\n", Jan, Feb, Mar, Apr)

	result := add(3, 4)
	fmt.Printf("3 + 4 = %d\n", result)

	// Multiple return values. variables can be assigned to multiple return values
	s, p := calculateSumAndProduct(5, 6)
	fmt.Printf("Sum: %d, Product: %d\n", s, p)

	adultCheck := checkAdult(32)
	fmt.Printf("Age 32 is: %s\n", adultCheck)

	dayOfWeekStr := dayOfWeek(3)
	fmt.Printf("Day 3 of the week is: %s\n", dayOfWeekStr)

	gradeStr := grade(85)
	fmt.Printf("Score 85 gets grade: %s\n", gradeStr)

	sumUpTo10 := sumUpTo(10)
	fmt.Printf("Sum of numbers from 1 to 10 is: %d\n", sumUpTo10)

	factorialOf5 := factorial(5)
	fmt.Printf("Factorial of 5 is: %d\n", factorialOf5)

	fmt.Println("Loop with break and continue:")
	loopWithBreakAndContinue()

	// Uncomment the next line to see the infinite loop in action
	// infiniteLoop()

	arrayAndSliceDemo()
	matrixDemo()
}

// Functions
// function with parameters and return type
func add(a, b int) int {
	return a + b
}

func calculateSumAndProduct(x, y int) (int, int) {
	sum := x + y
	product := x * y
	return sum, product
}

// Control structures (if, for, switch)
// Adult age check
func checkAdult(age int) string {
	if age >= 18 {
		return "Adult"
	}
	return "Minor"
}

// switch statement.
// In go, you don't need to use 'break' at the end of each case. You also don't need to use 'default' case.
// Cases do not fall through by default. Use 'fallthrough' keyword if you want to.
func dayOfWeek(day int) string {
	switch day {
	case 1:
		return "Monday"
	case 2:
		return "Tuesday"
	case 3:
		return "Wednesday"
	case 4:
		return "Thursday"
	case 5:
		return "Friday"
	case 6:
		return "Saturday"
	case 7:
		return "Sunday"
	default:
		return "Invalid day"
	}
}

// If -else if, else
func grade(score int) string {
	if score >= 90 {
		return "A"
	} else if score >= 80 {
		return "B"
	} else if score >= 70 {
		return "C"
	} else if score >= 60 {
		return "D"
	} else {
		return "F"
	}
}

// For loop
func sumUpTo(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

// While loop (using for). There's no while keyword in Go
func factorial(n int) int {
	result := 1
	for n > 1 {
		result *= n
		n--
	}
	return result
}

// Infinite loop (using for)
func infiniteLoop() {
	for {
		fmt.Println("This will print forever. Use Ctrl+C to stop.")
	}
}

// Loop with break and continue
func loopWithBreakAndContinue() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue // Skip even numbers
		}
		if i > 7 {
			break // Stop the loop if i is greater than 7
		}
		fmt.Println(i) // Print odd numbers less than or equal to 7
	}
}

// Arraysssss and slices. Yay! :D
// Arrays have fixed size, slices are dynamic
func arrayAndSliceDemo() {
	// Array

	arr := [3]int{1, 2, 3} // Array of 3 integers
	fmt.Printf("Array: %v\n", arr)
	// Accessing array elements
	fmt.Printf("First element: %d\n", arr[0])
	fmt.Printf("Array length: %d\n", len(arr))

	// Another format:
	// var arr [5]int // Array of 5 integers
	// arr[0] = 1
	// arr[1] = 2
	// fmt.Printf("Array: %v\n", arr)

	// Slice
	slice := []int{1, 2, 3}
	slice = append(slice, 4)
	fmt.Printf("Slice: %v\n", slice)
	// Slicing a slice
	subSlice := slice[1:3] // elements from index 1 to 2 (3 is excluded). A range is half-open [start, end)
	fmt.Printf("Sub-slice: %v\n", subSlice)

	// last element of an array or slice: arr[len(arr)-1]
	fmt.Printf("Last element of slice: %d\n", slice[len(slice)-1])

	// Slices can be created from arrays. They can also grow and shrink dynamically
	arr2 := [5]int{10, 20, 30, 40, 50}
	sliceFromArray := arr2[1:4] // elements from index 1 to 3
	fmt.Printf("Slice from array: %v\n", sliceFromArray)
}

// Matrix:
func matrixDemo() {
	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Printf("Matrix: %v\n", matrix)
}
