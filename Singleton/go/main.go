package main

import (
	"fmt"
	"sync"
)

type User struct {
	Name string
	Age  int
}

var singleton *User
var once sync.Once

func GetUser() *User {
	once.Do(func() {
		singleton = &User{"hoony", 24}
	})
	return singleton
}

func (user *User) getName() string {
	return user.Name
}

func (user *User) setName(name string) {
	user.Name = name
}

func main() {
	user1 := GetUser()
	user2 := GetUser()

	if user1 == user2 {
		fmt.Println("같다")
	}
}
