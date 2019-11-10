package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"

	"github.com/jung-kurt/gofpdf"
)

type jsonStruct struct {
	Name string `json:"name"`
}

func main() {

	http.HandleFunc("/", templateHandler)
	http.ListenAndServe(":8080", nil)

}

//templateHandler renders a template and returns as http response
func templateHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t, err := template.ParseFiles("template.html")
	if err != nil {
		fmt.Fprintf(w, "Unable to load template")
	}

	data := jsonStruct{Name: "world"}
	t.Execute(w, data)

	if r.Method == http.MethodPost {

		var tpl bytes.Buffer
		if err := t.Execute(&tpl, data); err != nil {
			fmt.Println(err)
		}
		htmlStr := tpl.String()
		htmlToPDF("template.pdf", htmlStr)

	}
}

func htmlToPDF(path, htmlStr string) {

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)

	html := pdf.HTMLBasicNew()
	html.Write(40, htmlStr)

	err := pdf.OutputFileAndClose(path)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("PDF file has generated.")
	}

}
