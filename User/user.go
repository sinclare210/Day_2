package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type User struct {
	Name       string
	University string
	Age        int
}

func (user User) Display() {
	fmt.Printf("Hello, my name is %s, I am %d years old, and I study at %s.\n", user.Name, user.Age, user.University)
}



func NewUser(name, university string, age int) (User,error) {
	if name == "" || university == "" || age == 0 {
		return User{},errors.New("Invalid Inputs")
	}
	return User{
		Name:       name,
		University: university,
		Age:        age,
	}, nil
}



func (user User)Save() (error){
	saveToFile,err := json.Marshal(user)
	if err != nil{
		return err
	}

	return os.WriteFile(user.Name,saveToFile,0644)
}