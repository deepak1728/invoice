package service

import (
	"context"
	"invoice/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetInvoiceByEmail(ctx context.Context, filter primitive.M) ([]bson.M, error) {
	collection := db.GetCollection("invoice")

	result, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	defer result.Close(ctx)
	var results []bson.M

	if err = result.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func AddInvoice(ctx context.Context, details primitive.M) error {
	collection := db.GetCollection("invoice")

	_, err := collection.InsertOne(ctx, details)
	if err != nil {
		return err
	}
	return nil
}
