package main

//packages
//import "fmt"
//To import more packages need to use paranthesis shown below
import (
	"fmt"
	"math"
)

//functions
func greeting(name string) string {
	return "Hello " + name
}

func main() {
	fmt.Println("Hello")

	/*Variable starts here*/

	var name = "Harsha"
	var age = 24
	const isCool = true

	fmt.Println(name, age, isCool)
	fmt.Printf("%T\n", isCool)

	//shorthand method to create Variable
	newName, email := "Harsha", "harsha@setmore.com"
	size := 1.4
	fmt.Println(newName, email, math.Floor(size))
	fmt.Printf("%T\n", size)

	fmt.Println(greeting("Harsha"))

	/*Variable end here*/

	/*Arrays and Slices starts here*/

	/*var fruitArr [2]string

	//Assing Values
	fruitArr[0] = "Apple"
	fruitArr[1] = "PineApple"

	//Declare and Assign
	fruitArr := [2]string{"Apple", "Orange"}*/

	fruitSlice := []string{"Apple", "Orange", "Grape"}

	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:2])

	/*Arrays and Slices ends here*/

	/*Condtionals starts here*/

	x := 51
	y := 10

	//If else
	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	//else if
	color := "green"

	if color == "red" {
		fmt.Println("Color is Red")
	} else if color == "blue" {
		fmt.Println("Color is Blue")
	} else {
		fmt.Println("Color is not Blue or Red")
	}

	switch color {
	case "red":
		fmt.Println("Color is Red")
	case "blue":
		fmt.Println("Color is Blue")
	default:
		fmt.Println("Color is not Blue or Red")
	}

	/*Condtionals ends here*/

	/*Loops starts here*/

	//Long Method
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	//short method
	for i := 1; i <= 10; i++ {
		fmt.Printf("Number %d\n", i)
	}

	//Fizz Buzz
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

	/*Loops ends here*/

	/*maps starts here*/
	/*emails := make(map[string]string)

	//Assgin Key Values
	emails["ding"] = "ding@setmore.com"
	emails["dong"] = "dong@setmore.com"
	emails["dingdong"] = "dingdong@setmore.com"*/

	emails := map[string]string{"ding": "ding@setmore.com", "dong": "dong@setmore.com"}

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["ding"])

	//Deleting from map
	delete(emails, "ding")
	fmt.Println(emails)

	/*maps ends here*/

	/*Ranges starts Here*/
	ids := []int{1, 2, 3, 4, 5, 6, 7, 8}

	//Loop through ids using Range
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	//Not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	//Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum :: ", sum)

	//Range with Map
	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	/*Ranges starts Here*/

	/*Pointers starts here*/

	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)

	//Use * to read value from address
	fmt.Println(*b)
	fmt.Println(*&a)

	//Change value with pointer
	*b = 10
	fmt.Println(a)

}
