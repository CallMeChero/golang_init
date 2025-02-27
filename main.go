package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}

func main() {
	title, content := getUserNoteData()
	todoText := getUserInput("Todo text: ")
	todo, err := todo.New(todoText)
	userNote, err := note.New(title, content)

	todo.Display()
	err = saveData(todo)

	if err != nil {
		return
	}

	userNote.Display()
	err = saveData(userNote)

	if err != nil {
		return
	}

}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving note failed")
		return err
	}

	fmt.Println("Saving note successful")

	return nil
}

func getUserNoteData() (string, string) {
	title := getUserInput("Note title:")

	content := getUserInput("Note input:")

	return title, content
}

func getUserInput(promt string) string {
	fmt.Print(promt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
