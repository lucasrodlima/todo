package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/lucasrodlima/todo/internal/cli"
)

var commands = map[string]func([]string) error{
	"add":      cli.Add,
	"list":     cli.List,
	"complete": cli.Complete,
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load environment variables")
	}

	if len(os.Args) == 1 {
		log.Fatal("No command")
	}
	args := os.Args[1:]

	if fn, ok := commands[args[0]]; ok {
		if err := fn(args); err != nil {
			log.Fatal(fmt.Errorf("Error running the program: %w", err))
		}
	} else {
		log.Fatal("Wrong command!")
	}

}
