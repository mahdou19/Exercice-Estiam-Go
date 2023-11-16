package dictionary

type Entry struct {
	Word       string
	Definition string
}

func (e Entry) String() string {
	return e.Word + ":" + e.Definition
}

type Dictionary struct {
	entries map[string]Entry
}

func NewDictionary() *Dictionary {
	return &Dictionary{
		entries: make(map[string]Entry),
	}
}

func (d *Dictionary) Add(word string, definition string) {
	d.entries[word] = Entry{Word: word, Definition: definition}
}

func (d *Dictionary) List() []Entry {
	var entryList []Entry
	for _, entry := range d.entries {
		entryList = append(entryList, entry)
	}
	return entryList
}
