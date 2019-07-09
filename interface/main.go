// Intefaces is collection of Methods
package main

import "fmt"

//We can add more methods here
type car interface {
	description() string
}

//implements the car interface
type mitsubishi struct {
	color string
	speed int
}

//implements the car interface
type toyota struct {
	color  string
	speed  int
	wheels int
}

func (m mitsubishi) description() string {
	desc, _ := fmt.Printf("Mitsubishi color: %v, speed: %v", m.color, m.speed)
	return string(desc)
}

func (t toyota) description() string {
	desc, _ := fmt.Printf("Toyota color: %v, speed: %v and has %v wheels", t.color, t.speed, t.wheels)
	return string(desc)
}

func main() {
	var c car
	c = mitsubishi{"Red", 1500}
	fmt.Println(c.description())
	c = toyota{"Black", 1500, 4}
	fmt.Println(c.description())
}
