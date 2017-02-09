package cookbook

import (
	"fmt"
	"os"
)

type CookBook struct {
	Name    string
	content []record
}

type record struct {
	text      string
	imagePath string
}

func New(name string) CookBook {
	return CookBook{name, make([]record, 0)}
}

func (book *CookBook) Record(step string) {
	if book.content == nil {
		book.content = make([]record, 0)
	}
	book.content = append(book.content, record{text: step})
}

func (book *CookBook) String() string {
	recipe := ""
	for index, step := range book.content {
		recipe += fmt.Sprintf("%v.  %v\n", index+1, step.text)
		if step.imagePath != "" {
			recipe += fmt.Sprintf("+\nimage::%v[height=256,width=256,link=%v]\n", step.imagePath, step.imagePath)
		}
	}
	return recipe
}

func (book *CookBook) Clear() string {
	defer func() { book.content = nil }()
	return book.String()
}

func (book *CookBook) AtStep() int {
	return len(book.content)
}
func (book *CookBook) RegisterImage() string {
	name := fmt.Sprintf("screenshots/%v-%v.png", book.Name, book.AtStep())
	book.content[len(book.content)-1].imagePath = name
	return name
}

func (book *CookBook) ClearToFile() {
	f, err := os.Create("recipes/" + book.Name + ".adoc")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	f.Chmod(0777)
	f.WriteString(book.Clear())
}
