package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
)

type singleCommand struct {
	command string //if empty string, ignored in run, and empty line in markdown
}

func (c *singleCommand) isComment() bool {
	for _, c := range c.command {
		if c == ' ' {
			continue
		} else if c == '#' {
			return true
		}
	}

	return false
}

type CommandGroup struct {
	title    string
	commands []singleCommand // can be only one singleCommand
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
			if cmd.command == "" {
				continue // empty Command is ignored
			}

			fmt.Println("### Executing the following command ###")
			fmt.Println(cmd.command)
			fmt.Print("[y/n] ")

			input.Scan()
			text := input.Text()
			switch text {
			case "y":
				fmt.Println("executing")
				execCmd := exec.Command("sh", "-c", cmd.command)
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
				fmt.Fprintln(w, cmd.command)
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
