package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var commands = map[string]func([]string) error{
	"add": createTodo,
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load environment variables")
	}

	args := os.Args[1:]

	if fn, ok := commands[args[0]]; ok {
		if err := fn(args); err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal("Wrong command!")
	}

}
