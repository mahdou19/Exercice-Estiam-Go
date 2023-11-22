package main

import (
	"encoding/json"
	"estiam/dictionary"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	d.Add(entry.Word, entry.Definition)
	response := map[string]string{"message": "Entrée ajoutée avec succès"}
	json.NewEncoder(w).Encode(response)
}

func actionDefine(d *dictionary.Dictionary, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	word := vars["word"]

	entry, err := d.Get(word)
	if err != nil {
		http.Error(w, "Mot non trouvé", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(entry)
}

func actionRemove(d *dictionary.Dictionary, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	word := vars["word"]

	d.Remove(word)
	response := map[string]string{"message": "Supprimé avec succès"}
	json.NewEncoder(w).Encode(response)
}

func actionList(d *dictionary.Dictionary, w http.ResponseWriter, r *http.Request) {

	entries := d.List()
	fmt.Println("Liste des entrées du dictionnaire :", entries)

	json.NewEncoder(w).Encode(entries)
}
