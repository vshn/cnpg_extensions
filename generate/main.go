// Scaffold a new extension directory from templates.
// Usage: go run ./generate -name <ext> [-package <pkg>] [-distros <d1,d2>] [-versions <v1,v2>]
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type TemplateData struct {
	Name           string
	SQLName        string
	Package        string
	Distros        []string
	Versions       []string
	DefaultDistro  string
	DefaultVersion string
}

func main() {
	name := flag.String("name", "", "Extension name, also used as directory name (required)")
	sqlName := flag.String("sql-name", "", "SQL name for CREATE EXTENSION (defaults to -name)")
	pkg := flag.String("package", "", `Debian package name pattern, e.g. "postgresql-%version%-repack" (defaults to postgresql-%version%-<name>)`)
	distros := flag.String("distros", "bookworm,trixie", "Comma-separated list of distros")
	versions := flag.String("versions", "18", "Comma-separated list of PostgreSQL major versions")
	outDir := flag.String("output", ".", "Parent directory where the extension directory is created")
	flag.Parse()

	if *name == "" {
		fmt.Fprintln(os.Stderr, "error: -name is required")
		flag.Usage()
		os.Exit(1)
	}

	if *sqlName == "" {
		*sqlName = *name
	}
	if *pkg == "" {
		*pkg = fmt.Sprintf("postgresql-%%version%%-%s", strings.ReplaceAll(*name, "_", "-"))
	}

	distroList := strings.Split(*distros, ",")
	versionList := strings.Split(*versions, ",")

	data := TemplateData{
		Name:           *name,
		SQLName:        *sqlName,
		Package:        *pkg,
		Distros:        distroList,
		Versions:       versionList,
		DefaultDistro:  distroList[0],
		DefaultVersion: versionList[0],
	}

	funcMap := template.FuncMap{
		"toTitle": func(s string) string {
			words := strings.Fields(strings.ReplaceAll(s, "_", " "))
			for i, w := range words {
				if len(w) > 0 {
					words[i] = strings.ToUpper(w[:1]) + strings.ToLower(w[1:])
				}
			}
			return strings.Join(words, " ")
		},
		"replaceAll": strings.ReplaceAll,
	}

	extDir := filepath.Join(*outDir, *name)
	if err := os.MkdirAll(extDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "error creating directory %s: %v\n", extDir, err)
		os.Exit(1)
	}

	tmplFiles := map[string]string{
		"templates/Dockerfile.tmpl":   filepath.Join(extDir, "Dockerfile"),
		"templates/README.md.tmpl":    filepath.Join(extDir, "README.md"),
		"templates/metadata.hcl.tmpl": filepath.Join(extDir, "metadata.hcl"),
	}

	for tmplPath, outPath := range tmplFiles {
		t, err := template.New(filepath.Base(tmplPath)).Funcs(funcMap).ParseFiles(tmplPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error parsing template %s: %v\n", tmplPath, err)
			os.Exit(1)
		}
		f, err := os.Create(outPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error creating %s: %v\n", outPath, err)
			os.Exit(1)
		}
		if err := t.Execute(f, data); err != nil {
			f.Close()
			fmt.Fprintf(os.Stderr, "error rendering %s: %v\n", tmplPath, err)
			os.Exit(1)
		}
		f.Close()
		fmt.Printf("generated %s\n", outPath)
	}
}
