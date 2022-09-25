package models

import (
	"context"
	"fmt"

	"github.com/abdrakhmanovzh/notemaker/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var db *mongo.Database

type Note struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

func init() {
	db = config.Connect()
}

func CreateNote(l *Note) *Note {
	res, err := db.Collection("notes").InsertOne(context.TODO(), l)
	if err != nil {
		fmt.Println(err)
	}
	createdNote := &Note{}
	filter := bson.D{{Key: "_id", Value: res.InsertedID}}
	createdRec := db.Collection("notes").FindOne(context.Background(), filter)

	createdRec.Decode(createdNote)
	return createdNote
}

func ShowAllNotes() []Note {
	var Notes []Note
	query := bson.D{{}}

	cursor, err := db.Collection("notes").Find(context.Background(), query)
	if err != nil {
		panic(err)
	}

	if err = cursor.All(context.Background(), &Notes); err != nil {
		panic(err)
	}
	return Notes
}

func ShowNote(id string) Note {
	var Note []Note

	que := bson.D{{Key: "id", Value: id}}
	cursor, err := db.Collection("notes").Find(context.Background(), que)
	if err != nil {
		fmt.Println(err)
	}

	if err = cursor.All(context.Background(), &Note); err != nil {
		fmt.Println(err)
	}
	return Note[0]
}

func DeleteNote(id string) {

	query := bson.D{{Key: "id", Value: id}}

	res, err := db.Collection("notes").DeleteOne(context.Background(), query)
	if err != nil {
		panic(err)
	}
	if res.DeletedCount < 1 {
		panic(err)
	}
}
