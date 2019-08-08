package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	const tpl = `
{{.Name}}{{range $child := .Children}}
 {{$child.Name}}{{range $grandchild := $child.Children}}
  {{$grandchild.Name}}{{end}}{{end}}
`
	type Tree struct {
		Name     string
		Children []Tree
	}

	root := Tree{
		"A",
		[]Tree{
			Tree{"B", []Tree{
				Tree{"C", []Tree{}},
			}},
			Tree{"D", []Tree{}},
		},
	}

	t := template.Must(template.New("tree").Parse(tpl))

	err := t.Execute(os.Stdout, root)
	if err != nil {
		log.Fatalf("executing template", err)
	}
}
