package main

import (
	"fmt"
	"strings"
)

/*
Este es un comentario de bloque.
Puedes escribir múltiples líneas aquí
sin tener que poner barras en cada una.
*/
func main() {
	var name = "test"
	fmt.Println(len(name))

	fmt.Println("Hello, World!")
	var same = strings.HasPrefix("test", "te") // true
	fmt.Println(same)
	var myArray = [3]string{"First", "Second", "Third"}
	fmt.Println(len(myArray))
	myArrayCopy := myArray
	fmt.Println(myArrayCopy)
	//agesMap := make(map[string]int)
	//fmt.Println(agesMap)
	//for loop
	for i := 0; i < len(myArray); i++ {
		fmt.Println(i)

	}
	i := 0
	// while loop
	for i < 10 {
		fmt.Println(i)
		i++
		if i == 2 {
			break
		}
	}
	numbers := []int{1, 2, 3}

	for _, num := range numbers {
		fmt.Printf("%d: %d\n", i, num)
	}
	var age int
	if age < 12 {
		//child
	} else if age < 18 {
		//teen
	} else {
		//adult
	}

	switch age {
	case 0:
		fmt.Println("Zero years old")
	case 1:
		fmt.Println("One year old")
	case 2:
		fmt.Println("Two years old")
	case 3:
		fmt.Println("Three years old")
	case 4:
		fmt.Println("Four years old")
	default:
		fmt.Println(" years old")
	}
	// declare a variable
	//var a = 1
	// intialize a variable
	//b := 1
	/**
		    flavio := Person{Name: "Flavio", Age: 39}
			fmt.Println(flavio.Age)
			fmt.Println(flavio.Name)

			var pablo Person
			pablo.Name = "Pablo"
			pablo.Age = 33
			result := sumTwoNumbers(1, 2)
			fmt.Println("Suma = ", result)
			sum, diff := performOperation(5, 2)
			fmt.Print(sum, diff)
	        ageptr := &age
		agevalue := *ageptr
		fmt.Println(age)
		fmt.Println(ageptr)
		fmt.Println(agevalue)
	*/

	age = 20
	increment(&age)
	flavio := Person{Name: "Flavio", Age: 39}
	flavio.Speak()
}

type FullName struct {
	FirstName string
	LastName  string
}

func (p Person) Speak() {
	fmt.Println("Hello from " + p.Name)
}

// type is like a structure or like collection of variables
type Person struct {
	Name string
	Age  int
}

func sumTwoNumbers(a int, b int) int {
	return a + b
}
func performOperation(a int, b int) (int, int) {
	return a + b, a - b
}

// function with pointers to increment the value of an int in 1
func increment(a *int) {
	*a = *a + 1
}

// interface
type Speaker interface {
	Speak()
}
