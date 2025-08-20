package book

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Book struct {
	Name   string
	Author string
}


func NewBook(name, author string) (Book,error) {
	if name == "" || author == "" {
		return Book{},errors.New("Invalid Inputs")
	}
	return  Book{
		Name:   name,
		Author: author,
	},nil
}

func (book Book)Display(){
	fmt.Printf("The book '%s' is written by %s.\n", book.Name, book.Author)
}

func (book Book)Save()(error){
	saveToFile,err := json.Marshal(book)
	fileName := strings.ReplaceAll(book.Name," ","_")
	if err != nil{
		return err
	}

	return os.WriteFile(fileName,saveToFile,0644)
}
