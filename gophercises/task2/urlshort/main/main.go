package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/asutptec4/go-training/gophercises/task2/urlshort"
)

func main() {
	yamlFileName := flag.String("yaml", "../../pathToUrls.yaml", "a yaml file with urls")
	useYaml := flag.Bool("useYaml", true, "use yaml file")
	flag.Parse()

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the JSONHandler using the mapHandler as the fallback
	json := `[{"path":"/urlshort-json", "url":"https://github.com/asutptec4/go-training"}]`
	jsonHandler, err := urlshort.JSONHandler([]byte(json), mapHandler)
	if err != nil {
		panic(err)
	}

	// Build the YAMLHandler using the jsonHandler as the fallback
	if *useYaml {
		yaml, err := ioutil.ReadFile(*yamlFileName)
		if err != nil {
			panic(err)
		}
		yamlHandler, err := urlshort.YAMLHandler(yaml, jsonHandler)
		if err != nil {
			panic(err)
		}
		startServer(yamlHandler)
		return
	}

	startServer(jsonHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func startServer(handler http.Handler) {
	// Strarting server
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", handler)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
