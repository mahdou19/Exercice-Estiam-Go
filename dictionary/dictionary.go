// dictionary.go
package dictionary

import (
	"encoding/json"
	"errors"
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

func (d *Dictionary) loadEntries() error {
	file, err := ioutil.ReadFile(d.filename)
	if err != nil {
		// Ignore error if file doesn't exist yet
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	err = json.Unmarshal(file, &d.entries)
	if err != nil {
		return err
	}

	return nil
}

func (d *Dictionary) saveEntries() error {
	data, err := json.Marshal(d.entries)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(d.filename, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (d *Dictionary) Add(word string, definition string) {
	err := d.loadEntries()
	if err != nil {
		panic(err)
	}

	d.entries[word] = Entry{Word: word, Definition: definition}

	err = d.saveEntries()
	if err != nil {
		panic(err)
	}
}

func (d *Dictionary) List() []Entry {
	err := d.loadEntries()
	if err != nil {
		panic(err)
	}

	var entryList []Entry
	for _, entry := range d.entries {
		entryList = append(entryList, entry)
	}
	return entryList
}

func (d *Dictionary) Remove(word string) {
	err := d.loadEntries()
	if err != nil {
		panic(err)
	}

	delete(d.entries, word)

	err = d.saveEntries()
	if err != nil {
		panic(err)
	}
}

func (d *Dictionary) Get(word string) (Entry, error) {
	err := d.loadEntries()
	if err != nil {
		return Entry{}, err
	}

	entry, found := d.entries[word]
	if !found {
		return Entry{}, errors.New("Mot non trouvé")
	}
	return entry, nil
}
