package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/russross/blackfriday"
)

func PresentAll(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "not yet")
}

func PresentSingle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println("got document id = " + vars["documentId"])
	document, err := getDocument(vars["documentId"])
	if err != nil {
		// If we didn't find it, 404
		w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "<html><body><p>not found</p></body></html>")
		return
	}

	contentBytes, _ := getContentBytes(document)

	output := blackfriday.MarkdownCommon(contentBytes)
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(output))
}

func getContentBytes(document Document) ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(document.Content)
	if err != nil {
		return []byte{}, err
	}

	return decoded, nil
}

func getDocument(documentId string) (Document, error) {
	url := "http://ec2-54-190-25-232.us-west-2.compute.amazonaws.com:8080/documents/" + documentId
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Did not get any response.")
		return Document{}, err
	}
	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()

	document := Document{}
	err = json.Unmarshal(body, &document)
	if err != nil {
		fmt.Println("could not unmarshall response: ")
		fmt.Println(err)
		return Document{}, err
	}

	fmt.Println("got " + document.Content)

	return document, nil
}
