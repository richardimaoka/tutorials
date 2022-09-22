package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
)

type SingleCommand struct {
	Comment string
	Command string //if empty string, ignored in run, and empty line in markdown
}

type CommandGroup struct {
	title    string
	commands []SingleCommand // can be only one SingleCommand
	results  string
}

func commandGroup(cmd ...string) CommandGroup {
	return commandGroup()
}

func (grp *CommandGroup) addTitle(title string) {
	grp.title = title
}

func (grp *CommandGroup) addResults(results string) {
	grp.results = results
}

func RunCommands(cmdGroups []CommandGroup) {
	input := bufio.NewScanner(os.Stdin)

	for _, grp := range cmdGroups {
		for _, cmd := range grp.commands {
			if cmd.Command == "" {
				continue // empty Command is ignored
			}

			fmt.Println("### Executing the following command ###")
			fmt.Println(cmd.Command)
			fmt.Print("[y/n] ")

			input.Scan()
			text := input.Text()
			switch text {
			case "y":
				fmt.Println("executing")
				execCmd := exec.Command("sh", "-c", cmd.Command)
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

func WriteMarkdown(w io.Writer, cmdGroups []CommandGroup) {
	for _, grp := range cmdGroups {
		fmt.Fprintln(w, grp.title)

		if len(grp.commands) > 0 {
			fmt.Fprintln(w, "```sh:コピペして実行")
			for _, cmd := range grp.commands {
				if cmd.Comment != "" {
					fmt.Fprintln(w, cmd.Comment)
				}
				fmt.Fprintln(w, cmd.Command)
			}
			fmt.Fprint(w, "```\n\n")
		}

		if grp.Results != "" {
			fmt.Fprintln(w, "```sh:コピペして実行")
			fmt.Fprintln(w, grp.results)
			fmt.Fprint(w, "```\n\n")
		}
	}
}
