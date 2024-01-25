// dictionary.go
package dictionary

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Entry struct {
	Word       string `json:"word"`
	Definition string `json:"definition"`
}

type Dictionary struct {
	collection *mongo.Collection
}

func NewDictionary(connectionString, dbName, collectionName string) (*Dictionary, error) {
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	collection := client.Database(dbName).Collection(collectionName)

	return &Dictionary{
		collection: collection,
	}, nil
}

func (d *Dictionary) Add(word string, definition string) error {

	entry := Entry{Word: word, Definition: definition}

	ctx := context.Background()
	filter := bson.M{"word": word}
	update := bson.M{"$set": entry}

	_, err := d.collection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	return err
}
func (d *Dictionary) List() ([]Entry, error) {
	cursor, err := d.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var entryList []Entry
	for cursor.Next(context.Background()) {
		var entry Entry
		if err := cursor.Decode(&entry); err != nil {
			return nil, err
		}
		entryList = append(entryList, entry)
	}

	return entryList, nil
}

func (d *Dictionary) Remove(word string) error {
	filter := bson.M{"word": word}
	_, err := d.collection.DeleteOne(context.Background(), filter)

	return err
}

func (d *Dictionary) Get(word string) (Entry, error) {
	var entry Entry

	filter := bson.M{"word": word}
	err := d.collection.FindOne(context.Background(), filter).Decode(&entry)

	return entry, err
}
