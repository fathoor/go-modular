package main

import (
	"fmt"
	"github.com/fathoor/go-modular/scripts/templates"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
	"text/template"
)

func main() {
	// Ensure module name is provided
	if len(os.Args) < 2 {
		println("Usage: go run scripts/module.go <module name>")
		return
	}

	cmd := exec.Command("go", "list", "-m")
	out, err := cmd.Output()
	if err != nil {
		log.Fatalf("Failed to obtain module: %v\n", err)
	}

	pkg := strings.TrimSpace(string(out))
	module := os.Args[1]
	modulePath := fmt.Sprintf("internal/modules/%s", module)

	// Initialize module directory
	directories := []string{
		modulePath,
		path.Join(modulePath, "internal/entity"),
		path.Join(modulePath, "internal/model"),
		path.Join(modulePath, "internal/repository"),
		path.Join(modulePath, "internal/repository/postgres"),
		path.Join(modulePath, "internal/usecase"),
		path.Join(modulePath, "internal/controller"),
		path.Join(modulePath, "internal/router"),
	}

	for _, directory := range directories {
		err := os.MkdirAll(directory, os.ModePerm)
		if err != nil {
			fmt.Printf("Failed to create directory %s: %v\n", directory, err)
			return
		}
	}

	// Initialize module files
	files := map[string]string{
		path.Join(modulePath, "internal/router", "router.go"): templates.RouterTmpl,
		path.Join(modulePath, fmt.Sprintf("%s.go", module)):   templates.ProviderTmpl,
	}

	for file, content := range files {
		tmpl, err := template.New(file).Parse(content)
		if err != nil {
			fmt.Printf("Failed to parse template %s: %v\n", file, err)
			return
		}

		f, err := os.Create(file)
		if err != nil {
			fmt.Printf("Failed to create file %s: %v\n", file, err)
			return
		}
		defer f.Close()

		data := struct {
			Package    string
			ModuleName string
			Name       string
		}{
			Package:    pkg,
			ModuleName: module,
			Name:       cases.Title(language.Indonesian).String(module),
		}

		if err = tmpl.Execute(f, data); err != nil {
			fmt.Printf("Failed to execute template %s: %v\n", file, err)
			return
		}

		fmt.Printf("Created %s\n", file)
	}

	fmt.Printf("Module %s has been initialized\n", module)
}
