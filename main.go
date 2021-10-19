package main

import (
	"ToDoProject/Config"
	"ToDoProject/Routes"
	"fmt"
)

var err error

func main() {
	// Create and connect to todo.db
	Config.InitDB()
	Config.DB_CreateTable()

	messages := make(chan string)
	go Endpoint()
	msg := <-messages
	fmt.Println(msg)
}

func Endpoint() {
	e := Routes.Setup()
	e.Logger.Fatal(e.Start(":8080"))
}
