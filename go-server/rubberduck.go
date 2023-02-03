package main

import "fmt"

type RubberDuck struct {
	name         string
	color        string
	age          int
	isJavaScript bool
}

func (rubberDuck *RubberDuck) makeJavaScript() {
	rubberDuck.name = "JavaScript"
	rubberDuck.isJavaScript = true
	rubberDuck.color = "yellow"
	rubberDuck.age = 0
}

func (rubberDuck *RubberDuck) isDead() bool {
	return rubberDuck.age > 20
}

func (rubberDuck *RubberDuck) kill() {
	rubberDuck.age = 21
	rubberDuck.color = "gone"
}

func (rubberDuck *RubberDuck) makeAngry() {
	rubberDuck.color = "red"
}

func (rubberDuck *RubberDuck) makeHappy() {
	rubberDuck.color = "yellow"
}

func (rubberDuck *RubberDuck) changeName(newName string) {
	rubberDuck.name = newName
}

func (rubberDuck *RubberDuck) isBaby() {
	rubberDuck.name = "baby duck"
	rubberDuck.isJavaScript = false
	rubberDuck.color = "yellow"
	rubberDuck.age = 0
}

func (rubberDuck *RubberDuck) printAge() {
	if rubberDuck.isDead() {
		fmt.Println("I am dead")
	} else {
		fmt.Println("I am", rubberDuck.age, "years old")
	}
}

func printDuckInfo(rubberDuck RubberDuck) {
	fmt.Println("Hi! My name is", rubberDuck.name)
	rubberDuck.printAge()
	fmt.Println("My feathers are", rubberDuck.color)

	if rubberDuck.isJavaScript && !rubberDuck.isDead() {
		fmt.Println("And I am the chosen one")
	} else {
		fmt.Println("And that's it.")
	}
}

func (rubberDuck *RubberDuck) growUp() {
	rubberDuck.color = "white"
}
