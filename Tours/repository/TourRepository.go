package repository

import (
	"context"
	"fmt"
	"time"
	"tours_service/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TourRepository struct {
	TourClient *mongo.Client
}

func (rep *TourRepository) getCollection() *mongo.Collection {
	tourDatabase := rep.TourClient.Database("mongodb")
	tourCollection := tourDatabase.Collection("tours")
	return tourCollection
}
func (pr *TourRepository) GetById(id int) (*model.Tour, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	toursCollection := pr.getCollection()

	var tour model.Tour
	err := toursCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&tour)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &tour, nil
}

func (rep *TourRepository) Insert(tour *model.Tour) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	tourCollection := rep.getCollection()

	result, err := tourCollection.InsertOne(ctx, &tour)
	if err != nil {
		fmt.Print(err)
		return err
	}
	fmt.Printf("Documents ID: %v\n", result.InsertedID)
	return nil
}

func (rep *TourRepository) Update(tour *model.Tour) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	tourCollection := rep.getCollection()
	t, _ := rep.GetById(tour.ID)
	if t == nil {
		fmt.Print("No id was found to update")
		return nil
	}
	filter := bson.M{"_id": tour.ID}
	updateData := bson.M{
		"$set": tour,
	}
	_, err := tourCollection.UpdateOne(ctx, filter, updateData)
	if err != nil {
		fmt.Print(err)
		return err
	}
	fmt.Printf("Documents ID: %v\n", tour.ID)
	return nil
}

// func (repo *TourRepository) GetAll() (*[]model.Tour, error) {
// 	var tours []model.Tour
// 	dbResult := repo.DatabaseConnection.Table(`tours."Tour"`).Find(&tours)
// 	if dbResult != nil {
// 		return &tours, dbResult.Error
// 	}

// 	return &tours, nil
// }
