// dictionary.go
package dictionary

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type Entry struct {
	Word       string `json:"word"`
	Definition string `json:"definition"`
}

type Dictionary struct {
	filename string
	entries  map[string]Entry
}

func NewDictionary(filename string) *Dictionary {
	return &Dictionary{
		filename: filename,
		entries:  make(map[string]Entry),
	}
}

func (d *Dictionary) loadEntries() {
	_, err := os.Stat(d.filename)
	if os.IsNotExist(err) {
		file, err := os.Create(d.filename)
		if err != nil {
			fmt.Println("Erreur créatsion fichier:", err)
			return
		}
		defer file.Close()
		return
	} else if err != nil {
		fmt.Println("Erreur:", err)
		return
	}

	data, err := ioutil.ReadFile(d.filename)
	if err != nil {
		fmt.Println("Erreur:", err)
		return
	}

	if len(data) == 0 {
		return
	}

	var entries []Entry
	if err := json.Unmarshal(data, &entries); err != nil {
		fmt.Println("Erreur dans les données:", err)
		return
	}

	for _, entry := range entries {
		d.Add(entry.Word, entry.Definition)
	}
}

func (d *Dictionary) saveEntries() error {

	entries := make([]Entry, 0, len(d.entries))
	for _, entry := range d.entries {
		entries = append(entries, entry)
	}

	data, err := json.MarshalIndent(entries, "", "	")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(d.filename, data, 0644)
}

func (d *Dictionary) Add(word string, definition string) {
	entry := Entry{Word: word, Definition: definition}
	d.entries[word] = entry

	err := d.saveEntries()
	if err != nil {
		fmt.Println("Erreur lors de la sauvegarde des entrées :", err)
	}
}

func (d *Dictionary) List() []Entry {
	d.loadEntries()
	var entryList []Entry
	for _, entry := range d.entries {
		entryList = append(entryList, entry)
	}
	return entryList
}

func (d *Dictionary) Remove(word string) {
	d.loadEntries()

	delete(d.entries, word)

	d.saveEntries()

}

func (d *Dictionary) Get(word string) (Entry, error) {
	d.loadEntries()

	entry, found := d.entries[word]
	if !found {
		return Entry{}, errors.New("Mot non trouvé")
	}
	return entry, nil
}
