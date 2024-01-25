package main

import (
	"encoding/json"
	"estiam/dictionary"
	"estiam/middleware"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

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
	mongodbURI := os.Getenv("MONGODB_URI")
	if mongodbURI == "" {
		log.Fatal("MONGODB_URI not set in environment variables")
	}
	dbName := os.Getenv("DB_NAME")
	collectionName := os.Getenv("COLLECTION_NAME")
	dict, err := dictionary.NewDictionary(mongodbURI, dbName, collectionName)
	if err != nil {
		log.Fatal("Error initializing dictionary:", err)
	}

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
func actionAdd(dict *dictionary.Dictionary, w http.ResponseWriter, r *http.Request) {
	var entry dictionary.Entry
	err := json.NewDecoder(r.Body).Decode(&entry)
	if err != nil {
		http.Error(w, "JSON decoding error", http.StatusBadRequest)
		log.Println("JSON decoding error:", err)
		return
	}
	if len(entry.Word) < 3 || len(entry.Definition) < 5 {
		http.Error(w, "Data does not meet validation rules", http.StatusBadRequest)
		log.Println("Data does not meet validation rules:", err)
		return

	}

	err = dict.Add(entry.Word, entry.Definition)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println("Error adding entry:", err)
		return
	}

	response := map[string]string{"message": "Entry added successfully"}
	json.NewEncoder(w).Encode(response)
}

func actionDefine(d *dictionary.Dictionary, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	word := vars["word"]

	if len(word) < 3 {
		http.Error(w, "Invalid word parameter", http.StatusBadRequest)
		log.Println("Invalid word parameter")
		return
	}

	entry, err := d.Get(word)
	if err != nil {
		http.Error(w, "Error, no word corresponding", http.StatusNotFound)
		log.Println("Error", err)
		return
	}

	json.NewEncoder(w).Encode(entry)
}

func actionRemove(d *dictionary.Dictionary, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	word := vars["word"]

	err := d.Remove(word)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println("Error delete word:", err)
		return
	}
	response := map[string]string{"message": "Delete successful"}
	json.NewEncoder(w).Encode(response)
}

func actionList(d *dictionary.Dictionary, w http.ResponseWriter, r *http.Request) {
	entries, err := d.List()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Println("Error getting list:", err)
		return
	}

	fmt.Println("List of dictionary entries:", entries)
	if entries == nil {
		http.Error(w, "No object saved", http.StatusOK)
		log.Println("Error getting list:", err)
		return
	}
	json.NewEncoder(w).Encode(entries)
}
