package main

import (
	"fmt"
	"os"
	"time"
)

type note struct {
	title string
	date  string
	task  map[string]string
}

// create a new note
func createNote(title string) note {
	n := note{
		title: title,
		date:  "",
		task:  map[string]string{},
	}
	return n
}

func (n note) format() string {
	fs := fmt.Sprintf("Title: %v \n", n.title)

	fs += fmt.Sprintf("%v \n \n", n.date)

	for k, v := range n.task {
		fs += fmt.Sprintf("%-20v => %v \n", k, v)
	}

	return fs
}

func (n *note) updateDate() string {
	month := time.Now().Month()
	day := time.Now().Day()
	year := time.Now().Year()
	n.date = fmt.Sprintf("%v %v, %v", month, day, year)
	return n.date
}

func (n *note) addTask(s string, t string) {
	n.task[s] = t
}

// save to .txt file
func (n *note) save() {
	data := []byte(n.format())

	err := os.WriteFile("notes/"+n.title+".txt", data, 0644)

	if err != nil {
		panic(err)
	}
	fmt.Println("New note saved!")
}
