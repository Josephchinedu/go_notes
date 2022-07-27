// mini golang note taking script.
package main

import (
	"fmt"
	"strings"
)

func main() {
	notebook := notesApplication{author: "Joseph Chinedu"}
	notebook.addNote("This is my first note")
	notebook.addNote("This is my second note")
	notebook.addNote("This is my third note")
	notebook.listNotes()
	notebook.get(2)
	notebook.search("third")
	notebook.edit(2, "The fundamental of javascript.")
	notebook.listNotes()
	notebook.delete(1)
}

// instantiate the noteapplication struct object
type notesApplication struct {
	author string
	notes  []string
}

// defines struct methods.

func (napp *notesApplication) addNote(note string) {
	napp.notes = append(napp.notes, note)
}

func (napp *notesApplication) listNotes() {
	for idx, note := range napp.notes {
		label1 := fmt.Sprintf("%s%d", "Note Id: ", idx+1)
		label2 := fmt.Sprintf("%s\n", note)
		fmt.Println(label1)
		fmt.Println(label2)
	}
}

func (napp *notesApplication) get(noteID int) {
	note := napp.notes[noteID-1]
	label := fmt.Sprintf("%s\n", note)
	fmt.Println(label)
}

func (napp *notesApplication) search(searchText string) {
	label := fmt.Sprintf("%s%s\n", "Showing results for search ", searchText)
	fmt.Println(label)

	for idx, note := range napp.notes {
		if strings.Contains(note, searchText) {
			label1 := fmt.Sprintf("%s%d", "Note ID: ", idx+1)
			label2 := fmt.Sprintf("%s\n", note)
			fmt.Println(label1)
			fmt.Println(label2)
		}
	}
}

func (napp *notesApplication) edit(noteId int, newContent string) {
	napp.notes[noteId-1] = newContent
}

func (napp *notesApplication) delete(noteId int) {
	napp.notes = append(napp.notes[:noteId-1], napp.notes[noteId:]...)
	fmt.Println(napp.notes)
}
