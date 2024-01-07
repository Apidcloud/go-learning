package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"-"`
	Id        int    `json:"id"`
	Title     string `json:"title,omitempty"`
	Completed bool   `json:"completed"`
}

func main() {
	url := "https://jsonplaceholder.typicode.com/todos/2"

	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		todoItem := Todo{}

		decoder := json.NewDecoder(response.Body)

		if err := decoder.Decode(&todoItem); err != nil {
			log.Fatal(err)
		}

		todo, err := json.MarshalIndent(todoItem, "", "\t")

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(todo))
	}

	// empty string field will be ommited
	anotherTodo := &Todo{1, 1, "", false}

	todo, err := json.MarshalIndent(anotherTodo, "", "\t")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(todo))
}
