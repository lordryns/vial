package filesystem

import (
	"os"
	"strings"
	"vial/boilerplate"

	"github.com/fatih/color"
)

// assume everything is being handled from within the parent directory
var PATHS = map[string]string{
	"website":                      "",
	"website/__init__.py":          boilerplate.INIT_PY,
	"website/views.py":             boilerplate.VIEWS_PY,
	"website/auth.py":              boilerplate.AUTH_PY,
	"main.py":                      boilerplate.MAIN_PY,
	"website/templates":            "",
	"website/static":               "",
	"website/templates/base.html":  boilerplate.BASE_HTML,
	"website/templates/index.html": boilerplate.INDEX_HTML,
	"website/static/style.css":     "",
	"requirements.txt":             boilerplate.REQUIREMENTS_TXT,
}

func CreateProject(path string) map[string]bool {
	var isCreated = createFolder(path)
	if isCreated {
		var fullResponse = createEverythingBothFilesAndDir(path)
		return fullResponse
	}

	return map[string]bool{}
}

func createEverythingBothFilesAndDir(projectpath string) map[string]bool {
	var createResults = make(map[string]bool)
	for path, content := range PATHS {
		if strings.Contains(path, ".") {
			var res = createFile(projectpath+"/"+path, content)
			createResults[path] = res
		} else {
			var res = createFolder(projectpath + "/" + path)
			createResults[path] = res
		}
	}

	return createResults
}

func createFolder(path string) bool {
	err := os.Mkdir(path, os.ModePerm)
	if err != nil {
		return false
	}
	return true
}

func createFile(path string, content string) bool {
	file, err := os.Create(path)
	if err != nil {
		color.Red(err.Error())
		return false
	}

	file.WriteString(content)

	defer file.Close()
	return true
}
