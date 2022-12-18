package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const notefile = "notes.txt"

func main() {
	// Open the notes file
	file, err := os.OpenFile(notefile, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Display menu and handle user input
	for {
		fmt.Println("\nWelcome to the notepad program!")
		fmt.Println("1. View all notes")
		fmt.Println("2. Add a new note")
		fmt.Println("3. Edit a note")
		fmt.Println("4. Quit")
		fmt.Print("Enter a number (1-4): ")

		var input int
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number (1-4).")
			continue
		}

		switch input {
		case 1:
			viewNotes(file)
		case 2:
			addNote(file)
		case 3:
			editNote(file)
		case 4:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid input. Please enter a number (1-4).")
		}
	}
}

func viewNotes(file *os.File) {
	// Read all lines from the file
	scanner := bufio.NewScanner(file)
	var notes []string
	for scanner.Scan() {
		notes = append(notes, scanner.Text())
	}

	// Display all notes
	fmt.Println("\nAll notes:")
	for i, note := range notes {
		fmt.Printf("%d. %s\n", i+1, note)
	}
}

func addNote(file *os.File) {
	// Get the new note from the user
	fmt.Print("\nEnter the new note: ")
	reader := bufio.NewReader(os.Stdin)
	note, _ := reader.ReadString('\n')
	note = strings.TrimSpace(note)

	// Append the new note to the file
	_, err := file.WriteString(note + "\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Note added successfully!")
}

func editNote(file *os.File) {
	// Read all lines from the file
	scanner := bufio.NewScanner(file)
	var notes []string
	for scanner.Scan() {
		notes = append(
