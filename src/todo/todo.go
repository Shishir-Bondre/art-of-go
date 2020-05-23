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
	fmt.Println("Enter the choice \n1. Create Note \n2. Retrieve Note \n3. Retrieve all Notes" +
		"\n4. Update note \n5. Delete note by id or title \n6. Delete all notes \nPress any key to exit")
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
		fmt.Println("==============================")
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
		fmt.Println("==============================")
		commandLine(notes)
	case "3":
		getAllNotes(notes)
		fmt.Println("==============================")
		commandLine(notes)
	case "4":
		fmt.Println("Title: ")
		title, _ := reader.ReadString('\n')
		fmt.Println("Id:")
		id, _ := reader.ReadString('\n')
		fmt.Println("-------------------------")
		fmt.Println("New Title: ")
		newTitle, _ := reader.ReadString('\n')
		fmt.Println("New Description:")
		newDescription, _ := reader.ReadString('\n')
		modifiedNote := Note{title:newTitle, description:newDescription}
		if newTitle ==""{
			modifiedNote = Note{title:title, description:newDescription}
		}
		modified, note := notes.updateNoteByIdOrTitle(id, title, modifiedNote)
		fmt.Println("Note modified:", modified)
		if modified{
			noteContents(note)
		}
		fmt.Println("==============================")
		commandLine(notes)
	case "5":
		fmt.Println("Title: ")
		title, _ := reader.ReadString('\n')
		fmt.Println("Id:")
		id, _ := reader.ReadString('\n')
		deleted := notes.deleteNoteByIdOrTitle(id, title)
		fmt.Println("Note deleted:", deleted)
		fmt.Println("==============================")
		commandLine(notes)
	case "6":
		notes = &Notes{}
		fmt.Println("==============================")
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

func (n *Notes) updateNoteByIdOrTitle(id string, title string, modifiedNote Note) (bool, Note) {
	for i:=0;i<len(n.notes);i++{
		if n.notes[i].title == title || n.notes[i].id.String() == id {
			n.notes[i].title = modifiedNote.title
			n.notes[i].description = modifiedNote.description
			n.notes[i].updateAt = time.Now()
			return true, n.notes[i]
		}
	}

	return false, Note{}
}

func (n *Notes) deleteNoteByIdOrTitle(id string, title string) bool {
	found := -1
	fmt.Println(title)
	for i:=0;i<len(n.notes);i++{
		if n.notes[i].title == title || n.notes[i].id.String() == id{
			found = i
			break
		}
	}
	fmt.Println(found)
	if found > -1{
		n.notes = append(n.notes[:found], n.notes[found+1:]...)
		return true
	}
	return false
}


func getAllNotes(notes *Notes){
	for _, note := range notes.notes {
		noteContents(note)
	}
}