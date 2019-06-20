package internal

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type Arc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Book map[string]Arc

func New(file http.File) (Book, error) {
	fs, err := file.Stat()
	if err != nil {
		return nil, err
	}
	bookData := make([]byte, fs.Size())
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	_, err = file.Read(bookData)

	if err != nil {
		return nil, err
	}

	var book Book

	err = json.Unmarshal(bookData, &book)

	if err != nil {
		return nil, err
	}

	return book, nil
}

func (book Book) LoadStory(arc string) (Arc, error) {
	nArc, ok := book[arc]

	if !ok {
		return nArc, errors.New("could not load story")
	}

	return nArc, nil
}
