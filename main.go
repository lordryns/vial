package main

import (
	"fmt"
	"os"

	"vial/filesystem"

	"github.com/fatih/color"
)

func main() {
	var badFiles int
	arg, err := getArg(1)
	if err != nil {
		fmt.Println("**Vial v 0.1**")
		fmt.Println("Welcome to Vial!")
		var current_dir, err = os.Getwd()
		if err == nil {
			fmt.Printf("cwd: %v", current_dir)
		}
		return
	}

	if arg == "create" {
		project_name, err := getArg(2)
		if err != nil {
			color.Red("[ Error ]: Not enough arguments to execute command!")
			fmt.Println("Hint: Use vial create <project-name> instead!")
			return
		}

		var res = filesystem.CreateProject(project_name)
		handleFileState(res, &badFiles)
		if badFiles > 0 {
			fmt.Printf("\nFailed to create %v files!\n", badFiles)
		}
		color.Green("%v: Project created successfully!", project_name)
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
