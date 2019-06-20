package app

import (
	"encoding/json"
	"fmt"
	"github.com/alemjc/gophercises/cyoa/rest/internal"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func loadStory(w http.ResponseWriter, r *http.Request, bookName, storyTitle string) {
	resourcesPath := os.Getenv("RESOURCES")
	files := http.Dir(resourcesPath)
	bookFile, err := files.Open(fmt.Sprintf("%s.%s", bookName, "json"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	mBook, err := internal.New(bookFile)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	storyArc, err := mBook.LoadStory(storyTitle)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	marshaled, err := json.Marshal(storyArc)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(marshaled)

}

func bookLists(w http.ResponseWriter, r *http.Request) {

	resourcesPath := os.Getenv("RESOURCES")
	files, err := ioutil.ReadDir(resourcesPath)
	response := make([]string, 0, len(files))

	if err != nil {
		http.Error(w, fmt.Sprintf("%s, %s", resourcesPath, err.Error()), http.StatusInternalServerError)
	}

	for _, file := range files {
		response = append(response, strings.TrimRight(file.Name(), ".json"))
	}

	marshaled, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(marshaled)

}
