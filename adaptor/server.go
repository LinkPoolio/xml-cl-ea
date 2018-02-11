package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"io/ioutil"
	"strings"
	"encoding/json"
	"log"
	"net/http"
	xj "github.com/basgys/goxml2json"
)

func Api() *rest.Api{
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/xmltojson", ConvertEndpoint),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	return api
}

func ConvertEndpoint(w rest.ResponseWriter, r *rest.Request) {
	// GET params
	endpoint := r.URL.Query().Get("endpoint")
	if endpoint == "" {
		rest.Error(w, "Invalid request, `endpoint` needs to be set.", http.StatusBadRequest)
		return
	}

	// Request the API
	rs, err := http.Get(endpoint)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Read the body of the request
	bodyBytes, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert XML from API to JSON
	rawJSON, err := xj.Convert(strings.NewReader(string(bodyBytes)))
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Replace the escaped backslashes and the XML at symbols
	replaced := strings.Replace(rawJSON.String(), "\\", "", -1)
	replaced = strings.Replace(replaced, "\"@", "\"", -1)

	//Convert string into golang map, prep for JSON output
	var jsonMap map[string]*json.RawMessage
	err = json.Unmarshal([]byte(trimQuotes(replaced)), &jsonMap)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteJson(jsonMap)
}

func trimQuotes(s string) string {
	if len(s) > 0 && s[0] == '"' {
		s = s[1:]
	}
	if len(s) > 0 && s[len(s)-1] == '"' {
		s = s[:len(s)-1]
	}
	if len(s) > 0 && s[len(s)-2] == '"' {
		s = s[:len(s)-2]
	}
	return s
}