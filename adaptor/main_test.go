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

type Agency struct {
	Tag              string `json:"-tag"`
	Title            string `json:"-title"`
	RegionTitle      string `json:"-regionTitle"`
}

type Body struct {
	Agencies []Agency `json:"agency"`
}

type Response struct {
	Body Body `json:"body"`
}

func TestXMLToJSON(t *testing.T) {
	server := httptest.NewServer(Api().MakeHandler())
	defer server.Close()

	res, err := http.Get(
		server.URL + "/xmltojson?endpoint=http://webservices.nextbus.com/service/publicXMLFeed?command=agencyList")
	if err != nil {
		log.Fatal(err)
	}
	books, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var responseObj Response
	json.Unmarshal(books, &responseObj)

	assert.NotZero(t, responseObj.Body.Agencies)
	assert.Equal(t, "actransit", responseObj.Body.Agencies[0].Tag)
}
