// dictionary_test.go
package dictionary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	dict := NewDictionary("./test_data.json")

	err := dict.Add("testWord", "testDefinition")
	assert.NoError(t, err, "Expected no error for valid input")

	err = dict.Add("", "testDefinition")
	assert.Error(t, err, "Expected error for invalid input")

}

func TestGet(t *testing.T) {
	dict := NewDictionary("./test_data.json")
	err := dict.Add("testWord", "testDefinition")
	assert.NoError(t, err, "Expected no error for valid input")

	entry, err := dict.Get("testWord")
	assert.NoError(t, err, "Expected no error for existing entry")
	assert.Equal(t, Entry{Word: "testWord", Definition: "testDefinition"}, entry, "Unexpected entry")

	entry, err = dict.Get("nonexistentWord")
	assert.Error(t, err, "Expected error for non-existing entry")
	assert.Equal(t, Entry{}, entry, "Expected empty entry for non-existing word")
}

func TestRemove(t *testing.T) {
	dict := NewDictionary("./test_data.json")
	err := dict.Add("testWord", "testDefinition")
	assert.NoError(t, err, "Expected no error for valid input")

	dict.Remove("testWord")
	entries := dict.List()
	assert.Empty(t, entries, "Expected empty dictionary after removal")

	dict.Remove("nonexistentWord")
	entries = dict.List()
	assert.Empty(t, entries, "Expected empty dictionary for non-existing word")
}

func TestList(t *testing.T) {
	dict := NewDictionary("./test_data.json")
	err := dict.Add("word1", "definition1")
	assert.NoError(t, err, "Expected no error for valid input")

	err = dict.Add("word2", "definition2")
	assert.NoError(t, err, "Expected no error for valid input")

	err = dict.Add("word3", "definition3")
	assert.NoError(t, err, "Expected no error for valid input")

	entries := dict.List()
	assert.Len(t, entries, 3, "Unexpected number of entries")
	assert.Contains(t, entries, Entry{Word: "word1", Definition: "definition1"}, "Missing entry word1")
	assert.Contains(t, entries, Entry{Word: "word2", Definition: "definition2"}, "Missing entry word2")
	assert.Contains(t, entries, Entry{Word: "word3", Definition: "definition3"}, "Missing entry word3")
}
