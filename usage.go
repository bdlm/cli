package cli

import (
	"log"
	"strings"
	"text/template"
)

var (
	usageTemplate     *template.Template
	templateFunctions = template.FuncMap{
		"ToLower": strings.ToLower,
	}
)

func init() {
	var err error

	// Usage text
	usageTemplate, err = template.
		New("usage").
		Funcs(templateFunctions).
		Parse(`
Usage: {{ .Name }} {{ .CLISummary }}

{{if .ShortDescription}}{{ .ShortDescription }}

{{end}}

{{if .Flags}}Flags:
{{range .Flags}}    {{ .Flag }}{{ .Description }}{{end}}

{{end}}
{{if .Commands}}Commands:
{{range .Commands}}    {{ .Command }}{{ .Description }}{{end}}

{{end}}
		`)
	if nil != err {
		log.Fatal(err)
	}
}
