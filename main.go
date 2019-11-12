package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"

	"github.com/jung-kurt/gofpdf"
)

type dataStruct struct {
	Name string `json:"name"`
	path, fileName string
}

func main() {

	http.HandleFunc("/", templateHandler)
	http.ListenAndServe(":8080", nil)

}

//templateHandler renders a template and returns as http response
func templateHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		data := dataStruct{
			Name: "world", 
			path: "template.html",
			fileName: "template.pdf",
		}
		data.pdfDownload(w, r)
	}
}

func (data dataStruct) pdfDownload(w http.ResponseWriter, r *http.Request) {

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)

	t, err := template.ParseFiles(data.path)
	if err != nil {
		fmt.Println(w, "Unable to load template")
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, data); err != nil {
		fmt.Println(err)
	} else {
		htmlStr := tpl.String()
		html := pdf.HTMLBasicNew()
		html.Write(40, htmlStr)
		fmt.Println("Success reading html from", data.path)
	}

	w.Header().Set("Content-Disposition", "attachment; filename=" + data.fileName)
	w.Header().Set("Content-Type", "application/pdf")
	pdf.Output(w)
	fmt.Println("Success downloading file", data.fileName)

}
