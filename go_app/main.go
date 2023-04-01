// Created by: Janet
// Created on: Sep 2020
//
// This program finds the area of a trapezoid

package main

import "fmt"

func main() {
	// This function finds the area of a trapezoid

	var height float64
	var base1 float64
	var base2 float64
	var area float64

	// input
	fmt.Println("This program finds the area of a trapezoid.")
	fmt.Println()
	fmt.Print("Enter the height (in cm): ")
	fmt.Scanln(&height)
	fmt.Print("Enter the length of base 1 (in cm): ")
	fmt.Scanln(&base1)
	fmt.Print("Enter the length of base 2 (in cm): ")
	fmt.Scanln(&base2)
	fmt.Println()

	// process
	area = 0.5 * (base1 + base2) * height

	// output
	fmt.Println("The area is:", area, "cm².")
	fmt.Println()
	fmt.Println("Done.")
}
