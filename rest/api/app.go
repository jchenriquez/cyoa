package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func makeHandler(route func(w http.ResponseWriter, r *http.Request, book string, title string)) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		route(w, r, vars["book"], vars["story"])

	}

}

func middleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		next.ServeHTTP(w, r)
	})

}

func Run(port int) error {
	router := mux.NewRouter()
	router.Host("http://localhost")
	router.Use(middleWare)
	//router.Handle("/", http.FileServer(http.Dir("/Users/alemjc/go/src/github.com/alemjc/gophercises/cyoa/web"))).Methods("GET", "OPTIONS")
	router.HandleFunc("/booklist", bookLists).Methods("GET")
	router.HandleFunc("/{book:[a-zA-Z]+}/{story}", makeHandler(loadStory)).Methods("GET")
	http.Handle("/", router)

	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
