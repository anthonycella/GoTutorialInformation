package main

import (
	"math/rand"
)

type RubberDuck struct {
	name         string
	color        string
	age          int
	isJavaScript bool
}

type Bird interface {
	isDead() bool
	kill()
}

func (rubberDuck *RubberDuck) fillDefaults() {
	names := [5]string{"jack", "jill", "billy", "bob", "joe"}
	colors := [7]string{"black", "blue", "white", "yellow", "gray", "purple", "green"}

	nameChosenIndex := rand.Int() % 5
	// fmt.Println("Name number is", nameChosenIndex)

	colorChosenIndex := rand.Int() % 7
	// fmt.Println("Color number is", colorChosenIndex)

	nameChosen := names[nameChosenIndex]
	colorChosen := colors[colorChosenIndex]
	ageChosen := rand.Int() % 21

	rubberDuck.name = nameChosen
	rubberDuck.color = colorChosen
	rubberDuck.age = ageChosen

	rubberDuck.isJavaScript = false
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

func (rubberDuck *RubberDuck) getAge() string {
	switch {
	case rubberDuck.isDead():
		return "I am dead"
	case rubberDuck.age == 1:
		return "I am " + string(rubberDuck.age) + " year old"
	case rubberDuck.age == 0:
		return "I am a newborn ducky"
	default:
		return "I am " + string(rubberDuck.age) + " years old"
	}
}

func (rubberDuck RubberDuck) String() string {
	name := "Hi! My name is" + rubberDuck.name + "\n"
	age := rubberDuck.getAge() + "\n"
	featherColor := "My feathers are " + rubberDuck.color + "\n"
	isChosenOne := ""

	if rubberDuck.isJavaScript && !rubberDuck.isDead() {
		isChosenOne = "And I am the chosen one\n\n"
	} else {
		isChosenOne = "And that's it.\n\n"
	}

	fullDuckInfo := name + age + featherColor + isChosenOne

	return fullDuckInfo
}

func (rubberDuck *RubberDuck) growUp() {
	rubberDuck.color = "white"
}
