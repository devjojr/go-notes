package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getUserInput(i string, r *bufio.Reader) (string, error) {
	fmt.Print(i)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func newNote() note {
	reader := bufio.NewReader(os.Stdin)

	title, _ := getUserInput("Add the title for your note: ", reader)

	n := createNote(title)
	fmt.Println("Title added:", n.title)

	return n
}

func noteOptions(n note) {
	reader := bufio.NewReader(os.Stdin)

	choice, _ := getUserInput("Press (t - Add your task, s - To save your note): ", reader)

	fmt.Println(choice)

	switch choice {
	case "t":
		name, _ := getUserInput("Task name: ", reader)

		nameLength := len(name)
		if nameLength == 0 {
			fmt.Println("Error: Task name is empty. Please enter a valid task name")
			noteOptions(n)
		}

		task, _ := getUserInput("Task: ", reader)

		taskLength := len(task)
		if taskLength == 0 {
			fmt.Println("Error: Task is emtpy. Please enter a valid task")
			noteOptions(n)
		}
		n.addTask(name, task)

		noteOptions(n)
	case "s":
		n.updateDate()
		n.save()
		fmt.Println("Your note has been saved as ", n.title)
	default:
		fmt.Println("Invalid option, please select again")
		noteOptions(n)
	}
}

func main() {
	myNote := newNote()
	noteOptions(myNote)
}
