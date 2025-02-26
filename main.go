package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
)

func main() {
	title, content := getUserNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()

	err = userNote.Save()

	if err != nil {
		fmt.Println("Saving note failed")
		return
	}

	fmt.Println("Saving note successful")
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
