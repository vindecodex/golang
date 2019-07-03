package main

import "fmt"

type Number int //created a new type of int that will be called as Number

type Hero struct { //created a new type that has a property of Name and Power that will be called as Hero
	Name  string
	Power string
	Age   int
}

func main() {
	var n Number                //declaring n with a type of Number
	n = 100                     //giving a value of n
	fmt.Printf("n: %v \n\n", n) //printing the value of n

	//method 1 of assigning value to a struct
	var hero = Hero{
		"Golu",
		"Spirit Ball",
		120,
	}
	fmt.Println("METHOD 1")
	fmt.Println(hero.Name)
	fmt.Println(hero.Power)
	fmt.Println(hero.Age)

	//method 2
	hero2 := Hero{
		"Guko",
		"Spirit Ball",
		120,
	}

	fmt.Println("METHOD 2")
	fmt.Println(hero2.Name)
	fmt.Println(hero2.Power)
	fmt.Println(hero2.Age)

	//method 3
	hero3 := Hero{"Goku", "Spirit Ball", 120}

	fmt.Println("METHOD 3")
	fmt.Println(hero3.Name)
	fmt.Println(hero3.Power)
	fmt.Println(hero3.Age)

	//method 4
	hero4 := Hero{
		Name:  "Goku",
		Power: "Spirit Ball",
		Age:   120,
	}

	fmt.Println("METHOD 4")
	fmt.Println(hero4.Name)
	fmt.Println(hero4.Power)
	fmt.Println(hero4.Age)

	//method 5
	var hero5 Hero
	hero5.Name = "Goku"
	hero5.Power = "Spirit Ball"
	hero5.Age = 120

	fmt.Println("METHOD 5")
	fmt.Println(hero5.Name)
	fmt.Println(hero5.Power)
	fmt.Println(hero5.Age)
}
