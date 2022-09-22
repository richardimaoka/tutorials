package main

import (
	"fmt"
	"io"
)

func WriteMarkdown(w io.Writer, cmdBlocks []fmt.Stringer) {
	for _, cmdBlk := range cmdBlocks {
		var commands []string
		switch v := cmdBlk.(type) {
		case *SingleCommand:
			commands = append(commands, v.Command)
		case *MultiCommands:
			commands = append(commands, v.Commands...)
		}

		fmt.Fprintln(w, "```sh:コピペして実行")
		for _, cmdString := range commands {
			fmt.Fprintln(w, cmdString)
		}
		fmt.Fprint(w, "```\n\n")
	}
}

func main() {
	fmt.Println("tutorials executed")
}
