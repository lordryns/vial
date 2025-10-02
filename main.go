package main

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
	"vial/filesystem"
	"vial/logger"
)

func main() {
	var newLogger = logger.NewLogger()
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

		filesystem.CreateProject(project_name)
		newLogger.Success(fmt.Sprintf("%v: Project created successfully!", project_name))
	} else if arg == "help" {
		newLogger.Info("HELP")
		fmt.Println(`help             => Displays this message
create <name>    => Creates new app`)
	} else {
		newLogger.Error("Not a valid command!")
		fmt.Println("Use 'help' to get the list of commands!")
	}
}
