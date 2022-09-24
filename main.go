package main

func main() {
	var cmds []CommandGroup
	// f := fmt.Sprintf

	cmds = append(cmds, Commands(
		"echo abc",
		"echo def",
		"",
		"# this is a comment",
		"echo jaff1",
		"echo jaff2",
		"echo jaff3",
		"echo jaff4",
	))

	RunCommands(cmds)
}
