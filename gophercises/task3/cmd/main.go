package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	storybook "github.com/asutptec4/go-training/gophercises/task3"
)

func main() {
	fileName := flag.String("file", "gopher.json", "the JSON file with the story")
	port := flag.Int("port", 3000, "the port to start server")
	flag.Parse()
	fmt.Printf("Using story in file: %s\n", *fileName)

	f, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}

	story, err := storybook.JsonStory(f)
	if err != nil {
		panic(err)
	}

	// fmt.Printf("%+v\n", story)
	h := storybook.NewHandler(story)
	fmt.Printf("Starting the server on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
