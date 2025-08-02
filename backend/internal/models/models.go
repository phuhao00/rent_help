package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Email       string             `bson:"email" json:"email" binding:"required,email"`
	Password    string             `bson:"password" json:"-"`
	FullName    string             `bson:"full_name" json:"full_name" binding:"required"`
	FirstName   string             `bson:"first_name" json:"first_name"`
	LastName    string             `bson:"last_name" json:"last_name"`
	Phone       string             `bson:"phone" json:"phone"`
	Avatar      string             `bson:"avatar" json:"avatar"`
	Role        string             `bson:"role" json:"role"` // "tenant", "landlord", "admin"
	IsVerified  bool               `bson:"is_verified" json:"is_verified"`
	Verified    bool               `bson:"verified" json:"verified"`
	DateOfBirth *time.Time         `bson:"date_of_birth" json:"date_of_birth,omitempty"`
	Gender      string             `bson:"gender" json:"gender,omitempty"` // "male", "female", "other"
	Occupation  string             `bson:"occupation" json:"occupation,omitempty"`
	Bio         string             `bson:"bio" json:"bio,omitempty"`
	Languages   []string           `bson:"languages" json:"languages,omitempty"`
	Preferences UserPreferences    `bson:"preferences" json:"preferences"`
	Address     Address            `bson:"address" json:"address,omitempty"`
	SocialLinks SocialLinks        `bson:"social_links" json:"social_links,omitempty"`
	Rating      UserRating         `bson:"rating" json:"rating"`
	IsActive    bool               `bson:"is_active" json:"is_active"`
	LastLoginAt *time.Time         `bson:"last_login_at" json:"last_login_at,omitempty"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}

type UserPreferences struct {
	Currency      string     `bson:"currency" json:"currency"`
	Language      string     `bson:"language" json:"language"`
	Timezone      string     `bson:"timezone" json:"timezone"`
	NotifyByEmail bool       `bson:"notify_by_email" json:"notify_by_email"`
	NotifyBySMS   bool       `bson:"notify_by_sms" json:"notify_by_sms"`
	NotifyByPush  bool       `bson:"notify_by_push" json:"notify_by_push"`
	SearchRadius  int        `bson:"search_radius" json:"search_radius"`
	PropertyTypes []string   `bson:"property_types" json:"property_types"`
	PriceRange    PriceRange `bson:"price_range" json:"price_range"`
}

type PriceRange struct {
	Min      float64 `bson:"min" json:"min"`
	Max      float64 `bson:"max" json:"max"`
	Currency string  `bson:"currency" json:"currency"`
}

type Address struct {
	Street     string  `bson:"street" json:"street"`
	City       string  `bson:"city" json:"city"`
	State      string  `bson:"state" json:"state"`
	Country    string  `bson:"country" json:"country"`
	ZipCode    string  `bson:"zip_code" json:"zip_code"`
	PostalCode string  `bson:"postal_code" json:"postal_code"`
	Latitude   float64 `bson:"latitude" json:"latitude"`
	Longitude  float64 `bson:"longitude" json:"longitude"`
}

type GeoLocation struct {
	Type        string    `bson:"type" json:"type"`               // "Point"
	Coordinates []float64 `bson:"coordinates" json:"coordinates"` // [longitude, latitude]
}

type SocialLinks struct {
	Facebook  string `bson:"facebook" json:"facebook,omitempty"`
	Twitter   string `bson:"twitter" json:"twitter,omitempty"`
	LinkedIn  string `bson:"linkedin" json:"linkedin,omitempty"`
	Instagram string `bson:"instagram" json:"instagram,omitempty"`
}

type UserRating struct {
	Average     float64 `bson:"average" json:"average"`
	TotalRating int     `bson:"total_rating" json:"total_rating"`
	AsLandlord  Rating  `bson:"as_landlord" json:"as_landlord"`
	AsTenant    Rating  `bson:"as_tenant" json:"as_tenant"`
}

type Rating struct {
	Average float64 `bson:"average" json:"average"`
	Count   int     `bson:"count" json:"count"`
}

type Property struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title             string             `bson:"title" json:"title" binding:"required"`
	Description       string             `bson:"description" json:"description" binding:"required"`
	Type              string             `bson:"type" json:"type" binding:"required"` // "apartment", "house", "condo", "townhouse", "studio"
	Price             float64            `bson:"price" json:"price" binding:"required"`
	Currency          string             `bson:"currency" json:"currency"`
	Address           Address            `bson:"address" json:"address"`
	Location          GeoLocation        `bson:"location" json:"location"`
	Bedrooms          int                `bson:"bedrooms" json:"bedrooms"`
	Bathrooms         int                `bson:"bathrooms" json:"bathrooms"`
	Area              int                `bson:"area" json:"area"`
	SquareFeet        int                `bson:"square_feet" json:"square_feet"`
	Features          PropertyFeatures   `bson:"features" json:"features"`
	Amenities         []string           `bson:"amenities" json:"amenities"`
	Images            []string           `bson:"images" json:"images"`
	PropertyImages    []PropertyImage    `bson:"property_images" json:"property_images"`
	Videos            []string           `bson:"videos" json:"videos,omitempty"`
	VirtualTour       string             `bson:"virtual_tour" json:"virtual_tour,omitempty"`
	FloorPlan         string             `bson:"floor_plan" json:"floor_plan,omitempty"`
	Available         bool               `bson:"available" json:"available"`
	AvailableFrom     time.Time          `bson:"available_from" json:"available_from"`
	LeaseTerms        []string           `bson:"lease_terms" json:"lease_terms"`
	PetsAllowed       bool               `bson:"pets_allowed" json:"pets_allowed"`
	SmokingAllowed    bool               `bson:"smoking_allowed" json:"smoking_allowed"`
	UtilitiesIncluded []string           `bson:"utilities_included" json:"utilities_included"`
	Rules             PropertyRules      `bson:"rules" json:"rules"`
	Utilities         PropertyUtilities  `bson:"utilities" json:"utilities"`
	Safety            PropertySafety     `bson:"safety" json:"safety"`
	Parking           PropertyParking    `bson:"parking" json:"parking"`
	OwnerID           primitive.ObjectID `bson:"owner_id" json:"owner_id"`
	ViewCount         int                `bson:"view_count" json:"view_count"`
	FavoriteCount     int                `bson:"favorite_count" json:"favorite_count"`
	Rating            PropertyRating     `bson:"rating" json:"rating"`
	Status            string             `bson:"status" json:"status"` // "draft", "published", "rented", "maintenance"
	Featured          bool               `bson:"featured" json:"featured"`
	Priority          int                `bson:"priority" json:"priority"`
	Tags              []string           `bson:"tags" json:"tags,omitempty"`
	CreatedAt         time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt         time.Time          `bson:"updated_at" json:"updated_at"`
}

type PropertyFeatures struct {
	Furnished       bool `bson:"furnished" json:"furnished"`
	PetsAllowed     bool `bson:"pets_allowed" json:"pets_allowed"`
	SmokingAllowed  bool `bson:"smoking_allowed" json:"smoking_allowed"`
	Balcony         bool `bson:"balcony" json:"balcony"`
	Garden          bool `bson:"garden" json:"garden"`
	Terrace         bool `bson:"terrace" json:"terrace"`
	Basement        bool `bson:"basement" json:"basement"`
	Attic           bool `bson:"attic" json:"attic"`
	Fireplace       bool `bson:"fireplace" json:"fireplace"`
	Pool            bool `bson:"pool" json:"pool"`
	Gym             bool `bson:"gym" json:"gym"`
	Elevator        bool `bson:"elevator" json:"elevator"`
	AccessibleEntry bool `bson:"accessible_entry" json:"accessible_entry"`
	StorageUnit     bool `bson:"storage_unit" json:"storage_unit"`
	LaundryRoom     bool `bson:"laundry_room" json:"laundry_room"`
}

type PropertyImage struct {
	URL        string    `bson:"url" json:"url"`
	Caption    string    `bson:"caption" json:"caption,omitempty"`
	IsPrimary  bool      `bson:"is_primary" json:"is_primary"`
	Order      int       `bson:"order" json:"order"`
	UploadedAt time.Time `bson:"uploaded_at" json:"uploaded_at"`
}

type LeaseTerms struct {
	MinLeaseDuration int     `bson:"min_lease_duration" json:"min_lease_duration"` // in months
	MaxLeaseDuration int     `bson:"max_lease_duration" json:"max_lease_duration"` // in months
	SecurityDeposit  float64 `bson:"security_deposit" json:"security_deposit"`
	BrokerFee        float64 `bson:"broker_fee" json:"broker_fee"`
	ApplicationFee   float64 `bson:"application_fee" json:"application_fee"`
	PetDeposit       float64 `bson:"pet_deposit" json:"pet_deposit,omitempty"`
	KeyDeposit       float64 `bson:"key_deposit" json:"key_deposit,omitempty"`
	RentInAdvance    int     `bson:"rent_in_advance" json:"rent_in_advance"` // months
}

type PropertyRules struct {
	MaxOccupants         int        `bson:"max_occupants" json:"max_occupants"`
	MinAge               int        `bson:"min_age" json:"min_age,omitempty"`
	GenderPreference     string     `bson:"gender_preference" json:"gender_preference,omitempty"` // "male", "female", "any"
	ProfessionPreference string     `bson:"profession_preference" json:"profession_preference,omitempty"`
	QuietHours           QuietHours `bson:"quiet_hours" json:"quiet_hours,omitempty"`
	VisitorPolicy        string     `bson:"visitor_policy" json:"visitor_policy,omitempty"`
	PartyPolicy          string     `bson:"party_policy" json:"party_policy,omitempty"`
	CleaningSchedule     string     `bson:"cleaning_schedule" json:"cleaning_schedule,omitempty"`
}

type QuietHours struct {
	Start string `bson:"start" json:"start"` // "22:00"
	End   string `bson:"end" json:"end"`     // "08:00"
}

type PropertyUtilities struct {
	Electricity UtilityInfo `bson:"electricity" json:"electricity"`
	Gas         UtilityInfo `bson:"gas" json:"gas"`
	Water       UtilityInfo `bson:"water" json:"water"`
	Internet    UtilityInfo `bson:"internet" json:"internet"`
	CableTV     UtilityInfo `bson:"cable_tv" json:"cable_tv"`
	Heating     UtilityInfo `bson:"heating" json:"heating"`
	Cooling     UtilityInfo `bson:"cooling" json:"cooling"`
	Trash       UtilityInfo `bson:"trash" json:"trash"`
}

type UtilityInfo struct {
	Included    bool    `bson:"included" json:"included"`
	Description string  `bson:"description" json:"description,omitempty"`
	Cost        float64 `bson:"cost" json:"cost,omitempty"`
}

type PropertySafety struct {
	SmokeDetector    bool `bson:"smoke_detector" json:"smoke_detector"`
	CarbonMonoxide   bool `bson:"carbon_monoxide" json:"carbon_monoxide"`
	FireExtinguisher bool `bson:"fire_extinguisher" json:"fire_extinguisher"`
	SecuritySystem   bool `bson:"security_system" json:"security_system"`
	SecurityGuard    bool `bson:"security_guard" json:"security_guard"`
	CCTV             bool `bson:"cctv" json:"cctv"`
	GatedCommunity   bool `bson:"gated_community" json:"gated_community"`
	WellLit          bool `bson:"well_lit" json:"well_lit"`
}

type PropertyParking struct {
	Available   bool    `bson:"available" json:"available"`
	Type        string  `bson:"type" json:"type,omitempty"` // "garage", "covered", "open", "street"
	Spaces      int     `bson:"spaces" json:"spaces,omitempty"`
	Cost        float64 `bson:"cost" json:"cost,omitempty"`
	Description string  `bson:"description" json:"description,omitempty"`
}

type PropertyRating struct {
	Average       float64 `bson:"average" json:"average"`
	Count         int     `bson:"count" json:"count"`
	Location      float64 `bson:"location" json:"location"`
	Value         float64 `bson:"value" json:"value"`
	Cleanliness   float64 `bson:"cleanliness" json:"cleanliness"`
	Communication float64 `bson:"communication" json:"communication"`
	CheckIn       float64 `bson:"check_in" json:"check_in"`
	Accuracy      float64 `bson:"accuracy" json:"accuracy"`
}

type Location struct {
	Address         string          `bson:"address" json:"address"`
	City            string          `bson:"city" json:"city"`
	State           string          `bson:"state" json:"state"`
	Country         string          `bson:"country" json:"country"`
	PostalCode      string          `bson:"postal_code" json:"postal_code"`
	Latitude        float64         `bson:"latitude" json:"latitude"`
	Longitude       float64         `bson:"longitude" json:"longitude"`
	Neighborhood    string          `bson:"neighborhood" json:"neighborhood,omitempty"`
	PublicTransport []TransportInfo `bson:"public_transport" json:"public_transport,omitempty"`
	NearbyPlaces    []NearbyPlace   `bson:"nearby_places" json:"nearby_places,omitempty"`
}

type TransportInfo struct {
	Type        string  `bson:"type" json:"type"` // "bus", "metro", "train"
	Name        string  `bson:"name" json:"name"`
	Distance    float64 `bson:"distance" json:"distance"`         // in meters
	WalkingTime int     `bson:"walking_time" json:"walking_time"` // in minutes
}

type NearbyPlace struct {
	Type     string  `bson:"type" json:"type"` // "hospital", "school", "grocery", "restaurant", "gym"
	Name     string  `bson:"name" json:"name"`
	Distance float64 `bson:"distance" json:"distance"` // in meters
	Rating   float64 `bson:"rating" json:"rating,omitempty"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Name     string `json:"name" binding:"required"`
	Phone    string `json:"phone"`
	Role     string `json:"role" binding:"required"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
}

type Booking struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	PropertyID       primitive.ObjectID `bson:"property_id" json:"property_id" binding:"required"`
	TenantID         primitive.ObjectID `bson:"tenant_id" json:"tenant_id"`
	LandlordID       primitive.ObjectID `bson:"landlord_id" json:"landlord_id"`
	StartDate        time.Time          `bson:"start_date" json:"start_date" binding:"required"`
	EndDate          time.Time          `bson:"end_date" json:"end_date" binding:"required"`
	CheckInTime      string             `bson:"check_in_time" json:"check_in_time,omitempty"`
	CheckOutTime     string             `bson:"check_out_time" json:"check_out_time,omitempty"`
	Status           string             `bson:"status" json:"status"`                 // "pending", "confirmed", "checked_in", "checked_out", "cancelled", "completed"
	PaymentStatus    string             `bson:"payment_status" json:"payment_status"` // "pending", "partial", "paid", "refunded"
	RentAmount       float64            `bson:"rent_amount" json:"rent_amount"`
	SecurityDeposit  float64            `bson:"security_deposit" json:"security_deposit"`
	ServiceFee       float64            `bson:"service_fee" json:"service_fee"`
	CleaningFee      float64            `bson:"cleaning_fee" json:"cleaning_fee,omitempty"`
	TotalAmount      float64            `bson:"total_amount" json:"total_amount"`
	Currency         string             `bson:"currency" json:"currency"`
	PaymentMethod    string             `bson:"payment_method" json:"payment_method,omitempty"`
	PaymentIntentID  string             `bson:"payment_intent_id" json:"payment_intent_id,omitempty"`
	Message          string             `bson:"message" json:"message,omitempty"`
	SpecialRequests  []string           `bson:"special_requests" json:"special_requests,omitempty"`
	GuestInfo        GuestInfo          `bson:"guest_info" json:"guest_info,omitempty"`
	Agreements       []Agreement        `bson:"agreements" json:"agreements,omitempty"`
	CancellationInfo *CancellationInfo  `bson:"cancellation_info" json:"cancellation_info,omitempty"`
	Reviews          BookingReviews     `bson:"reviews" json:"reviews,omitempty"`
	CheckInDetails   *CheckInDetails    `bson:"check_in_details" json:"check_in_details,omitempty"`
	CheckOutDetails  *CheckOutDetails   `bson:"check_out_details" json:"check_out_details,omitempty"`
	CreatedAt        time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt        time.Time          `bson:"updated_at" json:"updated_at"`
}

type GuestInfo struct {
	Adults   int    `bson:"adults" json:"adults"`
	Children int    `bson:"children" json:"children,omitempty"`
	Infants  int    `bson:"infants" json:"infants,omitempty"`
	Purpose  string `bson:"purpose" json:"purpose,omitempty"` // "vacation", "business", "relocation", "other"
}

type Agreement struct {
	Type     string    `bson:"type" json:"type"` // "terms", "house_rules", "cancellation_policy"
	Content  string    `bson:"content" json:"content"`
	Agreed   bool      `bson:"agreed" json:"agreed"`
	AgreedAt time.Time `bson:"agreed_at" json:"agreed_at,omitempty"`
	Version  string    `bson:"version" json:"version,omitempty"`
}

type CancellationInfo struct {
	CancelledBy  primitive.ObjectID `bson:"cancelled_by" json:"cancelled_by"`
	CancelledAt  time.Time          `bson:"cancelled_at" json:"cancelled_at"`
	Reason       string             `bson:"reason" json:"reason"`
	RefundAmount float64            `bson:"refund_amount" json:"refund_amount"`
	RefundStatus string             `bson:"refund_status" json:"refund_status"` // "pending", "processed", "failed"
	RefundMethod string             `bson:"refund_method" json:"refund_method,omitempty"`
}

type BookingReviews struct {
	TenantReview   *Review `bson:"tenant_review" json:"tenant_review,omitempty"`
	LandlordReview *Review `bson:"landlord_review" json:"landlord_review,omitempty"`
}

type Review struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	BookingID      primitive.ObjectID `bson:"booking_id" json:"booking_id"`
	ReviewerID     primitive.ObjectID `bson:"reviewer_id" json:"reviewer_id"`
	RevieweeID     primitive.ObjectID `bson:"reviewee_id" json:"reviewee_id"`
	PropertyID     primitive.ObjectID `bson:"property_id" json:"property_id,omitempty"`
	Type           string             `bson:"type" json:"type"` // "tenant_to_landlord", "landlord_to_tenant", "property"
	Rating         ReviewRating       `bson:"rating" json:"rating"`
	Comment        string             `bson:"comment" json:"comment"`
	Pros           []string           `bson:"pros" json:"pros,omitempty"`
	Cons           []string           `bson:"cons" json:"cons,omitempty"`
	WouldRecommend bool               `bson:"would_recommend" json:"would_recommend"`
	IsPublic       bool               `bson:"is_public" json:"is_public"`
	Response       *ReviewResponse    `bson:"response" json:"response,omitempty"`
	CreatedAt      time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt      time.Time          `bson:"updated_at" json:"updated_at"`
}

type ReviewRating struct {
	Overall       float64 `bson:"overall" json:"overall"`
	Cleanliness   float64 `bson:"cleanliness" json:"cleanliness,omitempty"`
	Communication float64 `bson:"communication" json:"communication,omitempty"`
	CheckIn       float64 `bson:"check_in" json:"check_in,omitempty"`
	Accuracy      float64 `bson:"accuracy" json:"accuracy,omitempty"`
	Location      float64 `bson:"location" json:"location,omitempty"`
	Value         float64 `bson:"value" json:"value,omitempty"`
}

type ReviewResponse struct {
	Content   string    `bson:"content" json:"content"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
}

type CheckInDetails struct {
	ActualTime          time.Time `bson:"actual_time" json:"actual_time"`
	KeyHandover         bool      `bson:"key_handover" json:"key_handover"`
	PropertyWalkthrough bool      `bson:"property_walkthrough" json:"property_walkthrough"`
	InventoryCheck      bool      `bson:"inventory_check" json:"inventory_check"`
	Notes               string    `bson:"notes" json:"notes,omitempty"`
	Photos              []string  `bson:"photos" json:"photos,omitempty"`
	Signature           string    `bson:"signature" json:"signature,omitempty"`
}

type CheckOutDetails struct {
	ActualTime         time.Time `bson:"actual_time" json:"actual_time"`
	KeyReturn          bool      `bson:"key_return" json:"key_return"`
	PropertyInspection bool      `bson:"property_inspection" json:"property_inspection"`
	DamageAssessment   []Damage  `bson:"damage_assessment" json:"damage_assessment,omitempty"`
	CleaningStatus     string    `bson:"cleaning_status" json:"cleaning_status"` // "satisfactory", "needs_cleaning", "damage"
	DepositReturn      float64   `bson:"deposit_return" json:"deposit_return"`
	Notes              string    `bson:"notes" json:"notes,omitempty"`
	Photos             []string  `bson:"photos" json:"photos,omitempty"`
	Signature          string    `bson:"signature" json:"signature,omitempty"`
}

type Damage struct {
	Description string   `bson:"description" json:"description"`
	Cost        float64  `bson:"cost" json:"cost,omitempty"`
	Photos      []string `bson:"photos" json:"photos,omitempty"`
	Severity    string   `bson:"severity" json:"severity"` // "minor", "major", "severe"
}

// Message/Chat models
type Message struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	BookingID  primitive.ObjectID `bson:"booking_id" json:"booking_id,omitempty"`
	SenderID   primitive.ObjectID `bson:"sender_id" json:"sender_id"`
	ReceiverID primitive.ObjectID `bson:"receiver_id" json:"receiver_id"`
	Content    string             `bson:"content" json:"content"`
	Type       string             `bson:"type" json:"type"` // "text", "image", "file", "location"
	FileURL    string             `bson:"file_url" json:"file_url,omitempty"`
	IsRead     bool               `bson:"is_read" json:"is_read"`
	ReadAt     *time.Time         `bson:"read_at" json:"read_at,omitempty"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
}

// Notification models
type Notification struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID    primitive.ObjectID `bson:"user_id" json:"user_id"`
	Type      string             `bson:"type" json:"type"` // "booking", "message", "review", "payment", "system"
	Title     string             `bson:"title" json:"title"`
	Content   string             `bson:"content" json:"content"`
	Data      interface{}        `bson:"data" json:"data,omitempty"`
	IsRead    bool               `bson:"is_read" json:"is_read"`
	ActionURL string             `bson:"action_url" json:"action_url,omitempty"`
	Priority  string             `bson:"priority" json:"priority"` // "low", "medium", "high", "urgent"
	ExpiresAt *time.Time         `bson:"expires_at" json:"expires_at,omitempty"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}

// Payment models
type Payment struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	BookingID     primitive.ObjectID `bson:"booking_id" json:"booking_id"`
	PayerID       primitive.ObjectID `bson:"payer_id" json:"payer_id"`
	ReceiverID    primitive.ObjectID `bson:"receiver_id" json:"receiver_id"`
	Amount        float64            `bson:"amount" json:"amount"`
	Currency      string             `bson:"currency" json:"currency"`
	Type          string             `bson:"type" json:"type"`     // "rent", "deposit", "fee", "refund"
	Method        string             `bson:"method" json:"method"` // "credit_card", "bank_transfer", "paypal", "stripe"
	Status        string             `bson:"status" json:"status"` // "pending", "processing", "completed", "failed", "cancelled"
	ExternalID    string             `bson:"external_id" json:"external_id,omitempty"`
	TransactionID string             `bson:"transaction_id" json:"transaction_id,omitempty"`
	Description   string             `bson:"description" json:"description,omitempty"`
	ProcessedAt   *time.Time         `bson:"processed_at" json:"processed_at,omitempty"`
	CreatedAt     time.Time          `bson:"created_at" json:"created_at"`
}

// Favorite/Wishlist models
type Favorite struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID     primitive.ObjectID `bson:"user_id" json:"user_id"`
	PropertyID primitive.ObjectID `bson:"property_id" json:"property_id"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
}
