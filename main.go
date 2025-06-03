package main

import (
	"fmt"
	"os"

	"vial/filesystem"

	"github.com/fatih/color"
)

func main() {
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
			color.Red("Not enough arguments to execute command!")
			fmt.Println("Hint: Use vial create <project-name> instead!")
			return
		}

		var res = filesystem.CreateProject(project_name)
		fmt.Println(res)
		color.Green("%v: Project created successfully!", project_name)
	}
}
