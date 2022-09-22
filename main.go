package main

import "fmt"

type SingleCommand struct {
	Title   string
	Comment string
	Command string
}

type MultiCommands struct {
	Title    string
	Comment  string
	Commands []string
}

func main() {
	fmt.Println("tutorials executed")
}
