package service

import (
	"context"
	"invoice/db"
	"invoice/models"

	"go.mongodb.org/mongo-driver/bson"
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

	_, err := collection.InsertOne(context.Background(), &models.Invoice{
		FirstName: details.FirstName,
		LastName:  details.LastName,
		Gender:    details.Gender,
		Email:     details.Email,
		Password:  details.Password,
	})

	if err != nil {
		return err
	}
	return nil
}
