package main

import (
	"bufio"
	"fmt"
	"github.com/google/uuid"
	"os"
	"time"
)

type Note struct {
	title string
	description string
	createdAt time.Time
	updateAt time.Time
	id uuid.UUID
}

type Notes struct {
	notes [] Note
}

func main(){
	notes := &Notes{}
	fmt.Println("Todo list program")
	commandLine(notes)
}

func noteContents(note Note){
	fmt.Println("-------------------------",
		"\nId:", note.id,
		"\nTitle:", note.title,
		"\ndescription:", note.description,
		"\nCreated at: ", note.createdAt.Format("2006-01-02 15:04"),
		"\nUpdated at:", note.updateAt.Format("2006-01-02 15:04"),
		"\n--------------------------")
}

func commandLine(notes *Notes){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the choice \n1. Create Note \n2. Retrieve Note \n3. Retrieve all Notes \nPress any key to exit")
	var choice string
	fmt.Scanf("%s",&choice)
	switch choice {
	case "1":
		fmt.Println("Title:")
		title, _ := reader.ReadString('\n')
		fmt.Println("Description:")
		description, _ := reader.ReadString('\n')
		created, note := notes.addNote(title, description)
		fmt.Println("Note created:", created)
		if created {
			noteContents(note)
		}
		commandLine(notes)
	case "2":
		fmt.Println("Title: ")
		title, _ := reader.ReadString('\n')
		fmt.Println("Id:")
		id, _ := reader.ReadString('\n')
		found, note := notes.getNoteByIdOrTitle(id, title)
		fmt.Println("Note present:", found)
		if found{
			noteContents(note)
		}
		commandLine(notes)
	case "3":
		getAllNotes(notes)
		commandLine(notes)
	default:
		fmt.Printf("%T\n", choice)
		fmt.Println("Bye..")
	}
}

func (n *Notes) addNote(title, description string) (bool, Note){
	if title == ""{
		return false, Note{}
	}
	note := Note{
		title,
		description,
		time.Now(),
		time.Now(),
		uuid.New(),
	}
	n.notes = append(n.notes, note)
	return true, note
}


func (n *Notes) getNoteByIdOrTitle(id string, title string) (bool, Note) {
	for _, note := range n.notes {
		if note.title == title || note.id.String() == id {
			return true, note
		}
	}
	return false, Note{}
}

func getAllNotes(notes *Notes){
	for _, note := range notes.notes {
		noteContents(note)
	}
}