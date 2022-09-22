package main

func main() {
	var cmds []CommandGroup
	// f := fmt.Sprintf

	cmds = append(cmds, Commands(
		"echo abc",
		"echo def",
		"",
		"# this is a comment",
		"echo jaff",
	))

	RunCommands(cmds)
}
