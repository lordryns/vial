package filesystem

import (
	"os"
	"vial/boilerplate"

	"github.com/fatih/color"
)

// assume everything is being handled from within the parent directory

var FOLDERS = []string{"website", "website/templates", "website/static"}

var FILES = map[string]string{
	"website/__init__.py":          boilerplate.INIT_PY,
	"website/views.py":             boilerplate.VIEWS_PY,
	"website/auth.py":              boilerplate.AUTH_PY,
	"app.py":                       boilerplate.APP_PY,
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

	for _, path := range FOLDERS {
		var res = createFolder(projectpath + "/" + path)
		createResults[path] = res
	}

	for path, content := range FILES {
		var res = createFile(projectpath+"/"+path, content)
		createResults[path] = res
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
