package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)

	base := `
{{- define "aaa" -}}
    <title>{{ template "bbb" . -}}</title>
{{ end -}}

{{ define "bbb" }}{{ end -}}
`

	page := `
{{- template "aaa" . }}
{{ define "bbb" }}OK{{ end -}}
`

	t1 := template.Must(template.New("aaa").Parse(base))

	for i := 0; i < 2; i++ {
		t2 := template.Must(t1.Clone())
		t2 = template.Must(t2.New(fmt.Sprintf("%d", i)).Parse(page))

		fmt.Println("--- exec 1 ---")

		if err := t2.Execute(os.Stdout, nil); err != nil {
			log.Fatalf("err=%q", err)
		}

		fmt.Println("--- exec 2 ---")

		if err := t2.Execute(os.Stdout, nil); err != nil {
			log.Fatalf("err=%q", err)
		}
	}
}
