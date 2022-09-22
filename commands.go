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
	Title    string
	Commands []SingleCommand // can be only one SingleCommand
}

func RunCommands(cmdGroups []CommandGroup) {
	input := bufio.NewScanner(os.Stdin)

	for _, grp := range cmdGroups {
		for _, cmd := range grp.Commands {
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
		fmt.Fprintln(w, grp.Title)

		// code block starts ```
		fmt.Fprintln(w, "```sh:コピペして実行")
		for _, cmd := range grp.Commands {
			if cmd.Comment != "" {
				fmt.Fprintln(w, cmd.Comment)
			}
			fmt.Fprintln(w, cmd.Command)
		}
		fmt.Fprint(w, "```\n\n")
		//  code block ends ```

	}
}
