package main

import "fmt"

// Creating a child structure to implement basic Human actions
type Action struct {
	sentence string
}

// Declaring functionality of our Action structure
// Here we will be able to implement other actions i.e. walk, detect etc
func (action Action) speak() string {
	return fmt.Sprintf("%v", action.sentence)
}

// Creating a parent structure. This structure contains an embedded structure
// that suppose perform actions for our parent struct.
type Human struct {
	name string
	Action
}

func main() {
	// Declaring and initializing a parent struct as well as our embedded one.
	human := Human{
		name: "Mark",
		Action: Action{
			sentence: "Hello there!",
		},
	}

	// Usage of created struct.
	fmt.Printf("%v : %v \n", human.name, human.Action.speak())

	fmt.Printf("Action.sentece: %v \n", human.Action.sentence)

	// Implemening a basic interface
	type speaker interface {
		speak() string
	}

	// Using our embedded struct with an interface
	var s speaker = human
	fmt.Printf("speaker says %v", s.speak())

}
