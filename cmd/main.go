package main

import (
	"fmt"
	"log"
	"net/http"
)

type User struct {
	name string
	age uint16
	money int16
	avgGrades float64
	happiness float64
}

// func (user User) getAllInfo() string {
// 	return fmt.Sprintf("user name is: %s. He is %d and he has manoy equal: %d", user.name, user.age, user.money)
// }

// func (user *User) setNewName(newName string)  {
// 	user.name = newName
// }

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "hello world")
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func handleRequest() {
	http.HandleFunc("/", homePageHandler)
	fmt.Println("Server listening on port 3000")
	log.Panic(
		http.ListenAndServe(":8080", nil),
	)
}

func main() {
	handleRequest()
}