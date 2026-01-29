package main

import (
	"fmt"
	"os"

	// "vial/filesystem"
	"vial/filesystem"
	"vial/logger"

	"github.com/charmbracelet/huh"
	"github.com/common-nighthawk/go-figure"
)


type Configuration struct {
	projectTitle string 
	projectStyle string 
	packageManager string
	installAndStartServer bool
}

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
		var projectConfig = createProjectForm()

		filesystem.CreateProject(projectConfig.projectTitle)
		newLogger.Success(fmt.Sprintf("%v: Project created successfully!", projectConfig.projectTitle))
	} else if arg == "help" || arg == "h" || arg == "-h" || arg == "--help" {
		newLogger.Info("HELP")
		fmt.Println(`help 		=> Displays this message
create 		=> Creates new app`)
	} else if arg == "audit" {

	} else {
		newLogger.Error("Not a valid command!")
		fmt.Println("Use 'help' to get the list of commands!")
	}
}


func createProjectForm() Configuration {
	var config Configuration 
	var form = huh.NewForm(
		huh.NewGroup( 
			huh.NewInput().Title("Project name?").Value(&config.projectTitle). 
				Validate(func(s string) error { 
					if len(s) > 0 {
						return nil 
					} else {
						return fmt.Errorf("Project title cannot be left empty!")
					}
				}),
		),
		huh.NewGroup(	
			huh.NewSelect[string]().Title("Select project type"). 
			Options( 
				huh.NewOption("Simple", "simple"), 
				huh.NewOption("Standard", "standard"), 
				huh.NewOption("Complex", "complex"),
			).Value(&config.projectStyle),
			),

		huh.NewGroup(
			huh.NewSelect[string]().Title("Select package manager"). 
			Options( 
				huh.NewOption("Python Installer Package (pip)", "pip"), 
				huh.NewOption("UV", "uv"), 
			).Value(&config.packageManager),
		),
	)

	if err := form.Run(); err != nil {
		logger.NewLogger().Error("Process exited!")
		os.Exit(1)
	}
	return config
}
