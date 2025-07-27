package validation

import (
	"errors"
	"regexp"
	"strings"
	"unicode"
)

// ValidationError represents a validation error
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// Error implements the error interface
func (ve ValidationError) Error() string {
	return ve.Message
}

// ValidationErrors represents multiple validation errors
type ValidationErrors []ValidationError

// Error implements the error interface
func (ves ValidationErrors) Error() string {
	var messages []string
	for _, ve := range ves {
		messages = append(messages, ve.Message)
	}
	return strings.Join(messages, ", ")
}

// HasErrors returns true if there are validation errors
func (ves ValidationErrors) HasErrors() bool {
	return len(ves) > 0
}

// Validator struct for validation rules
type Validator struct {
	errors ValidationErrors
}

// NewValidator creates a new validator instance
func NewValidator() *Validator {
	return &Validator{
		errors: make(ValidationErrors, 0),
	}
}

// AddError adds a validation error
func (v *Validator) AddError(field, message string) {
	v.errors = append(v.errors, ValidationError{
		Field:   field,
		Message: message,
	})
}

// GetErrors returns all validation errors
func (v *Validator) GetErrors() ValidationErrors {
	return v.errors
}

// HasErrors returns true if there are validation errors
func (v *Validator) HasErrors() bool {
	return len(v.errors) > 0
}

// Validate methods

// ValidateRequired validates that a field is not empty
func (v *Validator) ValidateRequired(field, value, fieldName string) {
	if strings.TrimSpace(value) == "" {
		v.AddError(field, fieldName+" is required")
	}
}

// ValidateEmail validates email format
func (v *Validator) ValidateEmail(field, email string) {
	if email == "" {
		return // Skip if empty, use ValidateRequired separately
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		v.AddError(field, "Invalid email format")
	}
}

// ValidatePassword validates password strength
func (v *Validator) ValidatePassword(field, password string) {
	if password == "" {
		return // Skip if empty, use ValidateRequired separately
	}

	if len(password) < 8 {
		v.AddError(field, "Password must be at least 8 characters long")
		return
	}

	var hasUpper, hasLower, hasNumber, hasSpecial bool
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if !hasUpper {
		v.AddError(field, "Password must contain at least one uppercase letter")
	}
	if !hasLower {
		v.AddError(field, "Password must contain at least one lowercase letter")
	}
	if !hasNumber {
		v.AddError(field, "Password must contain at least one number")
	}
	if !hasSpecial {
		v.AddError(field, "Password must contain at least one special character")
	}
}

// ValidateMinLength validates minimum length
func (v *Validator) ValidateMinLength(field, value string, minLength int, fieldName string) {
	if value == "" {
		return // Skip if empty, use ValidateRequired separately
	}

	if len(value) < minLength {
		v.AddError(field, fieldName+" must be at least "+string(rune(minLength))+" characters long")
	}
}

// ValidateMaxLength validates maximum length
func (v *Validator) ValidateMaxLength(field, value string, maxLength int, fieldName string) {
	if len(value) > maxLength {
		v.AddError(field, fieldName+" must be no more than "+string(rune(maxLength))+" characters long")
	}
}

// ValidateNumericRange validates that a number is within range
func (v *Validator) ValidateNumericRange(field string, value, min, max float64, fieldName string) {
	if value < min || value > max {
		v.AddError(field, fieldName+" must be between "+string(rune(int(min)))+" and "+string(rune(int(max))))
	}
}

// ValidatePhoneNumber validates phone number format
func (v *Validator) ValidatePhoneNumber(field, phone string) {
	if phone == "" {
		return // Skip if empty, use ValidateRequired separately
	}

	// Simple phone validation (adjust regex as needed for your region)
	phoneRegex := regexp.MustCompile(`^[+]?[1-9]\d{1,14}$`)
	if !phoneRegex.MatchString(phone) {
		v.AddError(field, "Invalid phone number format")
	}
}

// ValidateURL validates URL format
func (v *Validator) ValidateURL(field, url string) {
	if url == "" {
		return // Skip if empty, use ValidateRequired separately
	}

	urlRegex := regexp.MustCompile(`^https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)$`)
	if !urlRegex.MatchString(url) {
		v.AddError(field, "Invalid URL format")
	}
}

// ValidateOneOf validates that value is one of allowed values
func (v *Validator) ValidateOneOf(field, value string, allowed []string, fieldName string) {
	if value == "" {
		return // Skip if empty, use ValidateRequired separately
	}

	for _, allowedValue := range allowed {
		if value == allowedValue {
			return
		}
	}

	v.AddError(field, fieldName+" must be one of: "+strings.Join(allowed, ", "))
}

// User validation functions

// ValidateUserRegistration validates user registration data
func ValidateUserRegistration(email, password, fullName, phone string) error {
	validator := NewValidator()

	validator.ValidateRequired("email", email, "Email")
	validator.ValidateEmail("email", email)

	validator.ValidateRequired("password", password, "Password")
	validator.ValidatePassword("password", password)

	validator.ValidateRequired("full_name", fullName, "Full name")
	validator.ValidateMinLength("full_name", fullName, 2, "Full name")
	validator.ValidateMaxLength("full_name", fullName, 100, "Full name")

	if phone != "" {
		validator.ValidatePhoneNumber("phone", phone)
	}

	if validator.HasErrors() {
		return validator.GetErrors()
	}

	return nil
}

// ValidateUserLogin validates user login data
func ValidateUserLogin(email, password string) error {
	validator := NewValidator()

	validator.ValidateRequired("email", email, "Email")
	validator.ValidateEmail("email", email)

	validator.ValidateRequired("password", password, "Password")

	if validator.HasErrors() {
		return validator.GetErrors()
	}

	return nil
}

// Property validation functions

// ValidateProperty validates property data
func ValidateProperty(title, description, address, propertyType string, price float64) error {
	validator := NewValidator()

	validator.ValidateRequired("title", title, "Title")
	validator.ValidateMinLength("title", title, 5, "Title")
	validator.ValidateMaxLength("title", title, 200, "Title")

	validator.ValidateRequired("description", description, "Description")
	validator.ValidateMinLength("description", description, 20, "Description")
	validator.ValidateMaxLength("description", description, 2000, "Description")

	validator.ValidateRequired("address", address, "Address")
	validator.ValidateMinLength("address", address, 10, "Address")
	validator.ValidateMaxLength("address", address, 500, "Address")

	allowedTypes := []string{"apartment", "house", "condo", "townhouse", "studio"}
	validator.ValidateRequired("type", propertyType, "Property type")
	validator.ValidateOneOf("type", propertyType, allowedTypes, "Property type")

	if price <= 0 {
		validator.AddError("price", "Price must be greater than 0")
	}
	if price > 1000000 {
		validator.AddError("price", "Price cannot exceed 1,000,000")
	}

	if validator.HasErrors() {
		return validator.GetErrors()
	}

	return nil
}

// ValidateObjectID validates MongoDB ObjectID format
func ValidateObjectID(field, id string) error {
	if id == "" {
		return errors.New(field + " is required")
	}

	// Basic ObjectID validation (24 hex characters)
	objectIDRegex := regexp.MustCompile(`^[a-fA-F0-9]{24}$`)
	if !objectIDRegex.MatchString(id) {
		return errors.New("Invalid " + field + " format")
	}

	return nil
}
