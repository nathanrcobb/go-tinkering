package main

import (
	"fmt"
	"sync"
)

// Simple struct
type person struct {
	firstName string
	lastName  string
	age       int
}

// Nested structs
type animal struct {
	name            string
	characteristics []string
}

type herbivore struct {
	animal
	eatHuman bool
	sync.Mutex
}

func main() {
	// Regular print messages
	//println("Hello World!")

	// Deferred function calls, using LIFO
	defer printMsg("Hope you had fun!")
	defer printMsg("That's all for now.")

	// Variable declaration, and assignment
	var b string
	b = " World!"

	// Variable declaration/assignment; equivalent to above
	a := "Hello"

	println(a, b)

	// Integer declaration
	sum := 0

	// Simple for loop
	for i := 0; i < 5; i++ {
		sum += i
	}
	println(sum)

	n := 1

	// Simple while loop
	for j := 1; j < 5; j++ {
		n *= 2
	}
	println(n)

	/*
		m := 1

		Infinite loop
		for {
			m /= 2
		}
	*/

	// Array
	arr := [5]int{0, 1, 2, 3, 4}

	for k := 0; k < 5; k++ {
		print(arr[k])
	}
	print("\n")

	// Slices (dynamic arrays)
	slice := []int{1, 10, 100, 1000, 10000}

	println(slice)

	// Under the hood, creates a new array with the added element
	slice = append(slice, -1)

	println(slice[len(slice)-1:])

	// Map (dictionary)
	currency := map[string]string{
		"AUD": "Australian Dollar",
		"GBP": "Great Britain Pound",
		"USD": "United States Dollar",
		"JPY": "Japan Yen",
	}

	// Delete a map element
	delete(currency, "JPY")

	// Iterate through a map
	for k, v := range currency {
		fmt.Printf("%v is %v\n", k, v)
	}

	// Struct member
	p1 := person{
		firstName: "Nathan",
		lastName:  "Cobb",
		age:       32,
	}

	fmt.Printf("%v %v is %v years old.\n", p1.firstName, p1.lastName, p1.age)

	// Golang uses "promotion" to enable access members of sub-classes at the top-member-level
	a1 := herbivore{
		animal: animal{
			name: "Goat",
			characteristics: []string{
				"eat grass", "lack sense", "are lazy",
			},
		},
		eatHuman: false,
	}

	fmt.Printf("Everyone knows that %vs", a1.name)
	for l := 0; l < len(a1.characteristics); l++ {
		if l < len(a1.characteristics)-1 {
			fmt.Printf(" %v,", a1.characteristics[l])
		} else {
			fmt.Printf(" and %v.\n", a1.characteristics[l])
		}
	}

	var t1 string
	if a1.eatHuman {
		t1 = "DO"
	} else {
		t1 = "DO NOT"
	}
	// Can lock mutex as a top-level-member
	a1.Lock()
	fmt.Printf("Oh, and %vs %v eat humans!\n", a1.name, t1)
}

func printMsg(a string) {
	println(a)
}
