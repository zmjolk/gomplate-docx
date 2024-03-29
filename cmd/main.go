package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type docMap struct {
	template string
	idx      int
}

func main() {

	// User picks a document to format
	fmt.Println("Hello, which document do you need to compile?")
	docList := []docMap{}
	documents, err := ioutil.ReadDir("./documents")
	if err != nil {
		panic("Can't find Documents dir? Something has gone seriously wrong")
	}
	for idx, document := range documents {
		idx = idx + 1
		docList = append(docList, docMap{document.Name(), idx})
		fmt.Println("[" + strconv.Itoa(idx) + "]: " + document.Name())
	}

	var docChoice string
	fmt.Printf("::> ")
	fmt.Scanln(&docChoice)

	for _, list := range docList {
		if s, _ := strconv.Atoi(docChoice); s == list.idx {
			docChoice = list.template
			break
		}
	}

	_, ok := os.Lstat("documents/" + docChoice)
	if ok != nil {
		panic("Can't find that Document! :O")
	}

	fmt.Println("You have picked " + docChoice)

	var docStructure []string
	templates, _ := ioutil.ReadDir("documents/" + docChoice)

	for {
		chosenParagraph, err := chooseParagraph(templates)
		if err != nil {
			fmt.Printf("%s", err)
		}
		if chosenParagraph == "q" {
			fmt.Println("Done, compiling document")
			break
		}
		docStructure = append(docStructure, chosenParagraph)
	}

	fmt.Printf("Document structure is: %s \n", docStructure)

	// Template in the required names

	// delete dupes in chosen paragraphs

	// discover required template vars

	// type Template struct {
	// 	Name string
	// }

	// val := Template{Name: "yes"}
	// for _, templateName := range docStructure {
	// 	// add documents/docChoice/ to front of all elements for templating
	// 	fmt.Println(templateName)
	// }

	// // probably want to exec all templates at once

	// t := template.New("doc")
	// t.ParseFiles(docStructure...) // should be able to splat this somehow
	// err = t.Execute(os.Stdout, val)
	// if err != nil {
	// 	fmt.Printf("%e", err)
	// }
	// ask user for vars

	// execute templates with given vars

	// load doc xml

	xmlFile, err := os.Open("templates/document.tmpl")
	if err != nil {
		panic("Cannot open document xml?!?")
	}
	defer xmlFile.Close()

	// oldDoc, _ := ioutil.ReadAll(xmlFile)
	// newDoc := strings.Replace(string(oldDoc), "REPLACE", "<w:p w:rsidR=\"00D171E0\" w:rsidRDefault=\"00E81949\"><w:r><w:t>Exciting New Text</w:t></w:r></w:p>", 1)

	f, _ := os.Create("base_docx/word/document.xml")
	defer f.Close()

	// var docContent []string
	// // add the actual content
	// for _, para := range docStructure {
	// 	paraText, _ := ioutil.ReadFile("/documents/" + docChoice + "/" + para)
	// 	docContent = append(docContent, "<w:p w:rsidR=\"00D171E0\" w:rsidRDefault=\"00E81949\"><w:r><w:t>"+string(paraText)+"</w:t></w:r></w:p>")
	// }
	// f.Write([]byte(docContent))
	// zipDoc()

}

func chooseParagraph(templates []os.FileInfo) (chosenTemplate string, err error) {

	// User picks the next paragraph
	templateList := []docMap{}
	var paraChoiceNo string

	fmt.Println("Pick the next paragraph:")
	for idx, template := range templates {
		idx = idx + 1
		templateList = append(templateList, docMap{template.Name(), idx})
		fmt.Println("[" + strconv.Itoa(idx) + "]: " + template.Name())
	}
	fmt.Println("[q]: finished")

	fmt.Scanln(&paraChoiceNo)
	if paraChoiceNo == "q" {
		return "q", nil
	}
	for _, list := range templateList {
		if s, _ := strconv.Atoi(paraChoiceNo); s == list.idx {
			return list.template, nil
		}
	}
	return "", errors.New("[Error]: unable to index paragraphs, are you sure that number exists?")
}
