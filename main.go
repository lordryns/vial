package main

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
	"vial/filesystem"
	"vial/logger"
)

func main() {
	var newLogger = logger.NewLogger()
	var badFiles int
	arg, err := getArg(1)
	if err != nil {
		var myHeader = figure.NewFigure("Vial", "", true)
		myHeader.Print()
		newLogger.Info("use `vial help` for help")
		return
	}

	if arg == "create" {
		project_name, err := getArg(2)
		if err != nil {
			newLogger.Error("Not enough arguments to execute command!")
			fmt.Println("Hint: Use vial create <project-name> instead!")
			return
		}

		var res = filesystem.CreateProject(project_name)
		handleFileState(res, &badFiles)
		if badFiles > 0 {
			fmt.Printf("\nFailed to create %v files!\n", badFiles)
		}
		color.Green("%v: Project created successfully!", project_name)
	} else if arg == "help" {
		newLogger.Info("HELP")
		fmt.Println(`help             => Displays this message
create <name>    => Creates new app`)
	} else {
		color.Red("[ Error ]: Not a valid command!")
		fmt.Println("Use 'help' to get the list of commands!")
	}
}

func handleFileState(files map[string]bool, badFiles *int) {
	for fpath, state := range files {
		if state {
			color.Green("%v: %v", fpath, state)
		} else {
			*badFiles += 1
			color.Red("%v: %v", fpath, state)
		}
	}
}
