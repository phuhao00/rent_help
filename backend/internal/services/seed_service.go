package services

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"

	"rent-help-backend/internal/models"
)

type SeedService struct {
	db *mongo.Database
}

func NewSeedService(db *mongo.Database) *SeedService {
	return &SeedService{db: db}
}

// SeedDatabase initializes the database with sample data if empty
func (s *SeedService) SeedDatabase(ctx context.Context) error {
	log.Println("Checking if database seeding is needed...")

	// Check if users collection is empty
	userCount, err := s.db.Collection("users").CountDocuments(ctx, bson.M{})
	if err != nil {
		return err
	}

	if userCount > 0 {
		log.Println("Database already contains data, skipping seeding")
		return nil
	}

	log.Println("Database is empty, starting seeding process...")

	// Seed users
	if err := s.seedUsers(ctx); err != nil {
		return err
	}

	// Seed properties
	if err := s.seedProperties(ctx); err != nil {
		return err
	}

	log.Println("Database seeding completed successfully!")
	return nil
}

func (s *SeedService) seedUsers(ctx context.Context) error {
	log.Println("Seeding users...")

	// Hash password for all sample users
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	users := []interface{}{
		models.User{
			ID:         primitive.NewObjectID(),
			FullName:   "Administrator",
			Email:      "admin@renthelp.com",
			Password:   string(hashedPassword),
			Phone:      "+1-555-0100",
			Role:       "admin",
			Avatar:     "https://api.dicebear.com/7.x/avataaars/svg?seed=admin",
			IsVerified: true,
			IsActive:   true,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		models.User{
			ID:         primitive.NewObjectID(),
			FullName:   "John Smith",
			Email:      "john.landlord@email.com",
			Password:   string(hashedPassword),
			Phone:      "+1-555-0101",
			Role:       "landlord",
			Avatar:     "https://api.dicebear.com/7.x/avataaars/svg?seed=john",
			IsVerified: true,
			IsActive:   true,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		models.User{
			ID:         primitive.NewObjectID(),
			FullName:   "Jane Doe",
			Email:      "jane.tenant@email.com",
			Password:   string(hashedPassword),
			Phone:      "+1-555-0102",
			Role:       "tenant",
			Avatar:     "https://api.dicebear.com/7.x/avataaars/svg?seed=jane",
			IsVerified: true,
			IsActive:   true,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		models.User{
			ID:         primitive.NewObjectID(),
			FullName:   "Mike Johnson",
			Email:      "mike.landlord@email.com",
			Password:   string(hashedPassword),
			Phone:      "+1-555-0103",
			Role:       "landlord",
			Avatar:     "https://api.dicebear.com/7.x/avataaars/svg?seed=mike",
			IsVerified: true,
			IsActive:   true,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
	}

	result, err := s.db.Collection("users").InsertMany(ctx, users)
	if err != nil {
		return err
	}

	log.Printf("Inserted %d sample users", len(result.InsertedIDs))
	return nil
}

func (s *SeedService) seedProperties(ctx context.Context) error {
	log.Println("Seeding properties...")

	// Get landlord IDs for property ownership
	var landlords []models.User
	cursor, err := s.db.Collection("users").Find(ctx, bson.M{"role": "landlord"})
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &landlords); err != nil {
		return err
	}

	if len(landlords) < 2 {
		log.Println("Not enough landlords found, skipping property seeding")
		return nil
	}

	properties := []interface{}{
		models.Property{
			ID:          primitive.NewObjectID(),
			Title:       "Modern Downtown Apartment",
			Description: "Beautiful 2-bedroom apartment in the heart of downtown. Walking distance to restaurants, shops, and public transportation.",
			Type:        "apartment",
			Price:       1200,
			Bedrooms:    2,
			Bathrooms:   1,
			Area:        850,
			Address: models.Address{
				Street:  "123 Main Street",
				City:    "Downtown",
				State:   "CA",
				ZipCode: "90210",
				Country: "USA",
			},
			Location: models.GeoLocation{
				Type:        "Point",
				Coordinates: []float64{-118.2437, 34.0522},
			},
			Amenities: []string{"parking", "gym", "laundry", "balcony"},
			Images: []string{
				"https://images.unsplash.com/photo-1522708323590-d24dbb6b0267?w=800",
				"https://images.unsplash.com/photo-1493809842364-78817add7ffb?w=800",
			},
			OwnerID:           landlords[0].ID,
			Available:         true,
			AvailableFrom:     time.Now(),
			LeaseTerms:        []string{"6_months", "12_months"},
			PetsAllowed:       false,
			SmokingAllowed:    false,
			UtilitiesIncluded: []string{"water", "heating"},
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
		},
		models.Property{
			ID:          primitive.NewObjectID(),
			Title:       "Cozy Studio Near University",
			Description: "Perfect for students! Fully furnished studio apartment near the university campus with all utilities included.",
			Type:        "studio",
			Price:       800,
			Bedrooms:    0,
			Bathrooms:   1,
			Area:        450,
			Address: models.Address{
				Street:  "456 University Ave",
				City:    "College Town",
				State:   "CA",
				ZipCode: "90211",
				Country: "USA",
			},
			Location: models.GeoLocation{
				Type:        "Point",
				Coordinates: []float64{-118.2537, 34.0622},
			},
			Amenities: []string{"wifi", "furnished", "laundry", "study_area"},
			Images: []string{
				"https://images.unsplash.com/photo-1586023492125-27b2c045efd7?w=800",
				"https://images.unsplash.com/photo-1560448204-e02f11c3d0e2?w=800",
			},
			OwnerID:           landlords[1].ID,
			Available:         true,
			AvailableFrom:     time.Now(),
			LeaseTerms:        []string{"3_months", "6_months", "12_months"},
			PetsAllowed:       false,
			SmokingAllowed:    false,
			UtilitiesIncluded: []string{"water", "electricity", "internet", "heating"},
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
		},
		models.Property{
			ID:          primitive.NewObjectID(),
			Title:       "Spacious Family House",
			Description: "Large 4-bedroom house perfect for families. Features a big backyard, garage, and quiet neighborhood.",
			Type:        "house",
			Price:       2500,
			Bedrooms:    4,
			Bathrooms:   3,
			Area:        2200,
			Address: models.Address{
				Street:  "789 Oak Street",
				City:    "Suburbia",
				State:   "CA",
				ZipCode: "90212",
				Country: "USA",
			},
			Location: models.GeoLocation{
				Type:        "Point",
				Coordinates: []float64{-118.2637, 34.0722},
			},
			Amenities: []string{"garage", "backyard", "fireplace", "dishwasher", "laundry"},
			Images: []string{
				"https://images.unsplash.com/photo-1568605114967-8130f3a36994?w=800",
				"https://images.unsplash.com/photo-1570129477492-45c003edd2be?w=800",
			},
			OwnerID:           landlords[0].ID,
			Available:         true,
			AvailableFrom:     time.Now().AddDate(0, 0, 30), // Available in 30 days
			LeaseTerms:        []string{"12_months", "24_months"},
			PetsAllowed:       true,
			SmokingAllowed:    false,
			UtilitiesIncluded: []string{"water"},
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
		},
		models.Property{
			ID:          primitive.NewObjectID(),
			Title:       "Luxury Penthouse with City View",
			Description: "Stunning penthouse apartment with panoramic city views. Features floor-to-ceiling windows, marble countertops, and premium amenities.",
			Type:        "apartment",
			Price:       3500,
			Bedrooms:    3,
			Bathrooms:   2,
			Area:        1800,
			Address: models.Address{
				Street:  "100 Skyline Drive",
				City:    "Downtown",
				State:   "CA",
				ZipCode: "90213",
				Country: "USA",
			},
			Location: models.GeoLocation{
				Type:        "Point",
				Coordinates: []float64{-118.2337, 34.0422},
			},
			Amenities: []string{"concierge", "rooftop_pool", "gym", "valet_parking", "wine_cellar"},
			Images: []string{
				"https://images.unsplash.com/photo-1545324418-cc1a3fa10c00?w=800",
				"https://images.unsplash.com/photo-1502672260266-1c1ef2d93688?w=800",
			},
			OwnerID:           landlords[1].ID,
			Available:         true,
			AvailableFrom:     time.Now(),
			LeaseTerms:        []string{"12_months", "24_months"},
			PetsAllowed:       false,
			SmokingAllowed:    false,
			UtilitiesIncluded: []string{"water", "heating", "cooling"},
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
		},
		models.Property{
			ID:          primitive.NewObjectID(),
			Title:       "Charming Victorian Townhouse",
			Description: "Historic Victorian townhouse with original details preserved. High ceilings, hardwood floors, and modern updates throughout.",
			Type:        "townhouse",
			Price:       2200,
			Bedrooms:    3,
			Bathrooms:   2,
			Area:        1600,
			Address: models.Address{
				Street:  "567 Heritage Lane",
				City:    "Historic District",
				State:   "CA",
				ZipCode: "90214",
				Country: "USA",
			},
			Location: models.GeoLocation{
				Type:        "Point",
				Coordinates: []float64{-118.2737, 34.0822},
			},
			Amenities: []string{"hardwood_floors", "fireplace", "garden", "parking", "storage"},
			Images: []string{
				"https://images.unsplash.com/photo-1449844908441-8829872d2607?w=800",
				"https://images.unsplash.com/photo-1523217582562-09d0def993a6?w=800",
			},
			OwnerID:           landlords[0].ID,
			Available:         true,
			AvailableFrom:     time.Now().AddDate(0, 0, 15), // Available in 15 days
			LeaseTerms:        []string{"6_months", "12_months", "24_months"},
			PetsAllowed:       true,
			SmokingAllowed:    false,
			UtilitiesIncluded: []string{"water"},
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
		},
		models.Property{
			ID:          primitive.NewObjectID(),
			Title:       "Modern Condo with Pool",
			Description: "Contemporary 2-bedroom condo in a luxury complex. Features modern appliances, granite countertops, and access to resort-style amenities.",
			Type:        "condo",
			Price:       1800,
			Bedrooms:    2,
			Bathrooms:   2,
			Area:        1200,
			Address: models.Address{
				Street:  "890 Luxury Blvd",
				City:    "Waterfront",
				State:   "CA",
				ZipCode: "90215",
				Country: "USA",
			},
			Location: models.GeoLocation{
				Type:        "Point",
				Coordinates: []float64{-118.2837, 34.0922},
			},
			Amenities: []string{"pool", "spa", "fitness_center", "tennis_court", "clubhouse"},
			Images: []string{
				"https://images.unsplash.com/photo-1560185007-cde436f6a4d0?w=800",
				"https://images.unsplash.com/photo-1512917774080-9991f1c4c750?w=800",
			},
			OwnerID:           landlords[1].ID,
			Available:         true,
			AvailableFrom:     time.Now(),
			LeaseTerms:        []string{"6_months", "12_months"},
			PetsAllowed:       true,
			SmokingAllowed:    false,
			UtilitiesIncluded: []string{"water", "heating", "cooling"},
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
		},
		models.Property{
			ID:          primitive.NewObjectID(),
			Title:       "Affordable Shared Living Space",
			Description: "Budget-friendly shared living space perfect for young professionals. Private bedroom with shared common areas including kitchen and living room.",
			Type:        "apartment",
			Price:       650,
			Bedrooms:    1,
			Bathrooms:   1,
			Area:        400,
			Address: models.Address{
				Street:  "321 Budget Street",
				City:    "Young Professional District",
				State:   "CA",
				ZipCode: "90216",
				Country: "USA",
			},
			Location: models.GeoLocation{
				Type:        "Point",
				Coordinates: []float64{-118.2937, 34.1022},
			},
			Amenities: []string{"shared_kitchen", "wifi", "laundry", "parking"},
			Images: []string{
				"https://images.unsplash.com/photo-1555636222-cae831e670b3?w=800",
				"https://images.unsplash.com/photo-1616486338812-3dadae4b4ace?w=800",
			},
			OwnerID:           landlords[0].ID,
			Available:         true,
			AvailableFrom:     time.Now(),
			LeaseTerms:        []string{"1_month", "3_months", "6_months"},
			PetsAllowed:       false,
			SmokingAllowed:    false,
			UtilitiesIncluded: []string{"water", "electricity", "internet"},
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
		},
	}

	result, err := s.db.Collection("properties").InsertMany(ctx, properties)
	if err != nil {
		return err
	}

	log.Printf("Inserted %d sample properties", len(result.InsertedIDs))
	return nil
}
