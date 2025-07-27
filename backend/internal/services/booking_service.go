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

type BookingService struct {
	collection *mongo.Collection
}

func NewBookingService(db *mongo.Database) *BookingService {
	return &BookingService{
		collection: db.Collection("bookings"),
	}
}

func (s *BookingService) CreateBooking(booking *models.Booking) error {
	booking.CreatedAt = time.Now()
	booking.UpdatedAt = time.Now()
	booking.Status = "pending"

	result, err := s.collection.InsertOne(context.Background(), booking)
	if err != nil {
		return err
	}

	booking.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (s *BookingService) GetBookings(filter bson.M, limit, skip int64) ([]*models.Booking, error) {
	var bookings []*models.Booking

	opts := options.Find().SetLimit(limit).SetSkip(skip).SetSort(bson.D{{Key: "created_at", Value: -1}})
	cursor, err := s.collection.Find(context.Background(), filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var booking models.Booking
		if err := cursor.Decode(&booking); err != nil {
			return nil, err
		}
		bookings = append(bookings, &booking)
	}

	return bookings, nil
}

func (s *BookingService) GetBookingByID(id primitive.ObjectID) (*models.Booking, error) {
	var booking models.Booking
	err := s.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&booking)
	if err != nil {
		return nil, err
	}
	return &booking, nil
}

func (s *BookingService) UpdateBooking(id primitive.ObjectID, updates bson.M) error {
	updates["updated_at"] = time.Now()
	_, err := s.collection.UpdateOne(
		context.Background(),
		bson.M{"_id": id},
		bson.M{"$set": updates},
	)
	return err
}

func (s *BookingService) DeleteBooking(id primitive.ObjectID) error {
	_, err := s.collection.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}

func (s *BookingService) GetBookingsByTenant(tenantID primitive.ObjectID) ([]*models.Booking, error) {
	return s.GetBookings(bson.M{"tenant_id": tenantID}, 100, 0)
}

func (s *BookingService) GetBookingsByProperty(propertyID primitive.ObjectID) ([]*models.Booking, error) {
	return s.GetBookings(bson.M{"property_id": propertyID}, 100, 0)
}
