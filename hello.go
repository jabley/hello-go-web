package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	indexHTML = `<!DOCTYPE html>
<html>
<head>
<title>Hello World!</title>
</head>
<body>
<article>
<h1>Simple app showing environmental variables</h1>
<table>
<tr>
<th>key</th>
<th>value</th>
</tr>
{{range .}}
<tr>
<td>{{.Key}}</td>
<td>{{.Value}}</td>
</tr>
{{end}}
</table>
</article>
</body>
</html>
`
)

// KeyValue makes the ENV vars into a first-class data structure
type KeyValue struct {
	Key   string
	Value string
}

// KeyValues is a shorter way of referencing an array
type KeyValues []*KeyValue

var (
	tmpl = template.Must(template.New("index.html").Parse(indexHTML))
)

func main() {
	flag.Parse()

	port := getDefaultConfig("PORT", "8080")

	http.HandleFunc("/", mainHandler)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func getDefaultConfig(name, fallback string) string {
	if val := os.Getenv(name); val != "" {
		return val
	}
	return fallback
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.Header().Set("Cache-Control", "private, no-cache, no-store, must-revalidate")
	values := getKeyValues()
	if err := tmpl.Execute(w, values); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getKeyValues() KeyValues {
	result := make(KeyValues, 2)
	result[0] = &KeyValue{"PORT", os.Getenv("PORT")}
	result[1] = &KeyValue{"HELLO_VERSION", os.Getenv("HELLO_VERSION")}
	return result
}

func newKeyValue(kv string) *KeyValue {
	s := strings.Split(kv, "=")
	return &KeyValue{Key: s[0], Value: s[1]}
}
