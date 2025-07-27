package services

import (
	"context"
	"time"

	"rent-help-backend/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PropertyService struct {
	collection *mongo.Collection
}

func NewPropertyService(db *mongo.Database) *PropertyService {
	return &PropertyService{
		collection: db.Collection("properties"),
	}
}

func (s *PropertyService) CreateProperty(property *models.Property) error {
	property.CreatedAt = time.Now()
	property.UpdatedAt = time.Now()
	property.Available = true

	result, err := s.collection.InsertOne(context.Background(), property)
	if err != nil {
		return err
	}

	property.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (s *PropertyService) GetProperties(filter bson.M, limit, skip int64) ([]*models.Property, error) {
	var properties []*models.Property

	opts := options.Find().SetLimit(limit).SetSkip(skip).SetSort(bson.D{{Key: "created_at", Value: -1}})
	cursor, err := s.collection.Find(context.Background(), filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var property models.Property
		if err := cursor.Decode(&property); err != nil {
			return nil, err
		}
		properties = append(properties, &property)
	}

	return properties, nil
}

func (s *PropertyService) GetPropertyByID(id primitive.ObjectID) (*models.Property, error) {
	var property models.Property
	err := s.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&property)
	if err != nil {
		return nil, err
	}
	return &property, nil
}

func (s *PropertyService) UpdateProperty(id primitive.ObjectID, updates bson.M) error {
	updates["updated_at"] = time.Now()
	_, err := s.collection.UpdateOne(
		context.Background(),
		bson.M{"_id": id},
		bson.M{"$set": updates},
	)
	return err
}

func (s *PropertyService) DeleteProperty(id primitive.ObjectID) error {
	_, err := s.collection.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}

func (s *PropertyService) GetPropertiesByOwner(ownerID primitive.ObjectID) ([]*models.Property, error) {
	return s.GetProperties(bson.M{"owner_id": ownerID}, 100, 0)
}
