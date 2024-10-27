package main

import (
	ccwccommand "ccwc/command"
	"os"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	programArgs := os.Args[1:]
	command, err := ccwccommand.NewCommandParser().Parse(programArgs[0:])
	if err != nil {
		panic(err)
	}
	for _, filePath := range command.InputFilesPaths {
		ccwccommand.NewWcFromFile(command.Options, filePath).Exec()
	}
}
