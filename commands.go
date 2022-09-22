package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

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

func (c *SingleCommand) String() string {
	if c.Comment == "" {
		return c.Command
	} else {
		return "# " + c.Comment + "\n" + c.Command
	}
}

func (c *MultiCommands) String() string {
	if c.Comment == "" {
		return strings.Join(c.Commands, "\n")
	} else {
		return "# " + c.Comment + "\n" + strings.Join(c.Commands, "\n")
	}
}

func RunCommands(cmdBlocks []fmt.Stringer) {
	input := bufio.NewScanner(os.Stdin)

	for _, cmdBlk := range cmdBlocks {
		var commands []string
		switch v := cmdBlk.(type) {
		case *SingleCommand:
			commands = append(commands, v.Command)
		case *MultiCommands:
			commands = append(commands, v.Commands...)
		}

		for _, cmdString := range commands {
			fmt.Println("### Executing the following command ###")
			fmt.Println(cmdString)
			fmt.Print("[y/n] ")

			input.Scan()
			text := input.Text()
			switch text {
			case "y":
				fmt.Println("executing")
				execCmd := exec.Command("sh", "-c", cmdString)
				output, _ := execCmd.CombinedOutput()
				fmt.Println(string(output))
			case "n":
				fmt.Println("skipping")
			default:
				fmt.Print("[y/n] ")
			}
		}
	}
}
