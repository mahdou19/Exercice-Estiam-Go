package dictionary

type Entry struct {
	word       string
	definition string
}

func (e Entry) String() string {
	return e.word + ":" + e.definition
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
	d.entries[word] = Entry{definition: definition}
}
