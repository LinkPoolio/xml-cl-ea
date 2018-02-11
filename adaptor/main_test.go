package main

import (
	"testing"
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
)

type Book struct {
	Id              int    `json:"id"`
	Type            string `json:"type"`
	Isbn            string `json:"isbn"`
	Description     string `json:"description"`
	Title           string `json:"author"`
	PublicationDate string `json:"publicationDate"`
}

type Books struct {
	Books []Book `json:"hydra:member"`
}

func TestXMLToJSON(t *testing.T) {
	server := httptest.NewServer(Api().MakeHandler())
	defer server.Close()

	res, err := http.Get(server.URL + "/xmltojson?endpoint=https://demo.api-platform.com/books")
	if err != nil {
		log.Fatal(err)
	}
	books, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var booksObj Books
	json.Unmarshal(books, &booksObj)

	assert.NotZero(t, booksObj.Books)
	assert.Equal(t, 1, booksObj.Books[0].Id)
}
