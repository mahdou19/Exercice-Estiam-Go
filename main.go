package main

import (
	"encoding/json"
	"estiam/dictionary"
	"estiam/middleware"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func errorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("Uncaught error:", err)
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
func main() {
	r := mux.NewRouter()
	r.Use(middleware.LoggingMiddleware)
	r.Use(middleware.AuthenticationMiddleware)
	r.Use(errorHandler)
	dict := dictionary.NewDictionary("./data.json")

	r.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		actionAdd(dict, w, r)
	}).Methods("POST")

	r.HandleFunc("/define/{word}", func(w http.ResponseWriter, r *http.Request) {
		actionDefine(dict, w, r)
	}).Methods("GET")

	r.HandleFunc("/remove/{word}", func(w http.ResponseWriter, r *http.Request) {
		actionRemove(dict, w, r)
	}).Methods("DELETE")

	r.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		actionList(dict, w, r)
	}).Methods("GET")

	http.Handle("/", r)

	fmt.Println("Server listening on :3000...")
	http.ListenAndServe(":3000", nil)
}

func actionAdd(d *dictionary.Dictionary, w http.ResponseWriter, r *http.Request) {
	var entry dictionary.Entry
	err := json.NewDecoder(r.Body).Decode(&entry)
	if err != nil {
		http.Error(w, "JSON decoding error", http.StatusBadRequest)
		log.Println("JSON decoding error:", err)
		return
	}

	if len(entry.Word) < 3 || len(entry.Definition) < 5 {
		http.Error(w, "Data does not meet validation rules", http.StatusBadRequest)
		log.Println("Data validation error")
		return
	}

	d.Add(entry.Word, entry.Definition)
	response := map[string]string{"message": "Entry added successfully"}
	json.NewEncoder(w).Encode(response)
}

func actionDefine(d *dictionary.Dictionary, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	word := vars["word"]

	entry, err := d.Get(word)
	if err != nil {
		http.Error(w, "Word not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(entry)
}

func actionRemove(d *dictionary.Dictionary, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	word := vars["word"]

	d.Remove(word)
	response := map[string]string{"message": "Delete successful"}
	json.NewEncoder(w).Encode(response)
}

func actionList(d *dictionary.Dictionary, w http.ResponseWriter, r *http.Request) {

	entries := d.List()
	fmt.Println("List of dictionary entries:", entries)

	json.NewEncoder(w).Encode(entries)
}
