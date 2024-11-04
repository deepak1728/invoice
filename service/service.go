package service

import (
	"context"
	"fmt"
	"invoice/db"
	"invoice/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetInvoiceByEmail(email string) (models.Invoice, error) {
	var invoice models.Invoice

	collection := db.GetCollection("invoice")

	err := collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&invoice)
	if err != nil {
		return invoice, err
	}

	return invoice, nil
}

func AddInvoice(details *models.Invoice) error {
	collection := db.GetCollection("invoice")

	indexModel := mongo.IndexModel{
		Keys: bson.M{
			"email": 1,
		},
		Options: options.Index().SetUnique(true),
	}

	_, err := collection.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		return fmt.Errorf("failed to create index: %v", err)
	}

	bsonData, err := bson.Marshal(details)
	if err != nil {
		return err
	}

	var doc bson.M
	err = bson.Unmarshal(bsonData, &doc)
	if err != nil {
		return err
	}

	_, err = collection.InsertOne(context.Background(), doc)
	if err != nil {
		return err
	}
	return nil
}
