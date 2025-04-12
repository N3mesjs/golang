package main

import ("fmt")

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type wheel struct {
	Radius int
	Material string
}

type car struct {
	Make  string
	Model string
	Year  int
	FrontWheel wheel
	BackWheel  wheel
}

// embedding struct, i can use the wheel struct as a field of the carFiga struct, and i can access the fields of the wheel struct directly from the carFiga struct
type carFiga struct {
	Colore string
	wheel
}

func main(){
	person := Person{
		FirstName: "John",
		Age: 32,
	}
	// same as var person Person
	person.LastName = "Doe"
	fmt.Println(person)

	var person2 Person = Person{
		FirstName: "Jane",
		LastName:  "Doe",
		Age:       32,
	}
	fmt.Println(person2)

	printPerson(person)
	printPerson(Person{FirstName: "Skibidi", LastName: "Toeletta", Age: 69})

	fmt.Print("------- # ------\n")

	toyotaWheel := wheel{
		Radius: 15,
		Material: "Aluminum",
	}

	car1 := car{
		Make:  "Toyota",
		Model: "Corolla",
		Year:  2020,
		FrontWheel: toyotaWheel,
		BackWheel: toyotaWheel,
	}
	// here i call the method i defined for the car struct
	car1.printCar()

	carFiga1 := carFiga{
		Colore: "Blu",
		wheel: wheel{
			Radius: 15,
			Material: "Aluminum",
		},
	}
	fmt.Println(carFiga1)
	fmt.Println(carFiga1.Radius) // i can access the fields of the wheel struct directly from the carFiga struct


	// anonymous struct, i would likely use it if i only need it once
	pisello := struct {
		Lunghezza string
		Colore    string
	} {
		Lunghezza: "10cm",
		Colore:    "Verde",
	}

	fmt.Println(pisello)
}

// we can define methods for structs, the method is a function that has a receiver, in this case the receiver is a pointer to the Person struct, so i can modify the struct directly

func printPerson(p Person) {
	fmt.Println(p.FirstName, p.LastName, p.Age)
}

func (c car) printCar() {
	fmt.Println(c.Make, c.Model, c.Year)
	fmt.Println("Front Wheel:", c.FrontWheel.Radius, c.FrontWheel.Material)
	fmt.Println("Back Wheel:", c.BackWheel.Radius, c.BackWheel.Material)
}