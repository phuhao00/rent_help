package handlers

import (
	"net/http"
	"strconv"

	"rent-help-backend/internal/models"
	"rent-help-backend/internal/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PropertyHandler struct {
	propertyService *services.PropertyService
}

func NewPropertyHandler(propertyService *services.PropertyService) *PropertyHandler {
	return &PropertyHandler{
		propertyService: propertyService,
	}
}

func (h *PropertyHandler) CreateProperty(c *gin.Context) {
	userIDStr, _ := c.Get("user_id")
	ownerID, err := primitive.ObjectIDFromHex(userIDStr.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var property models.Property
	if err := c.ShouldBindJSON(&property); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	property.OwnerID = ownerID

	if err := h.propertyService.CreateProperty(&property); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create property"})
		return
	}

	c.JSON(http.StatusCreated, property)
}

func (h *PropertyHandler) GetProperties(c *gin.Context) {
	// Parse query parameters
	limitStr := c.DefaultQuery("limit", "10")
	skipStr := c.DefaultQuery("skip", "0")
	city := c.Query("city")
	propertyType := c.Query("type")

	limit, _ := strconv.ParseInt(limitStr, 10, 64)
	skip, _ := strconv.ParseInt(skipStr, 10, 64)

	// Build filter
	filter := bson.M{"available": true}
	if city != "" {
		filter["location.city"] = bson.M{"$regex": city, "$options": "i"}
	}
	if propertyType != "" {
		filter["type"] = propertyType
	}

	properties, err := h.propertyService.GetProperties(filter, limit, skip)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get properties"})
		return
	}

	c.JSON(http.StatusOK, properties)
}

func (h *PropertyHandler) GetProperty(c *gin.Context) {
	idStr := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid property ID"})
		return
	}

	property, err := h.propertyService.GetPropertyByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Property not found"})
		return
	}

	c.JSON(http.StatusOK, property)
}

func (h *PropertyHandler) UpdateProperty(c *gin.Context) {
	idStr := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid property ID"})
		return
	}

	userIDStr, _ := c.Get("user_id")
	ownerID, err := primitive.ObjectIDFromHex(userIDStr.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Check if user owns the property
	property, err := h.propertyService.GetPropertyByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Property not found"})
		return
	}

	if property.OwnerID != ownerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized to update this property"})
		return
	}

	var updates bson.M
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Remove sensitive fields
	delete(updates, "_id")
	delete(updates, "owner_id")

	if err := h.propertyService.UpdateProperty(id, updates); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update property"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Property updated successfully"})
}

func (h *PropertyHandler) DeleteProperty(c *gin.Context) {
	idStr := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid property ID"})
		return
	}

	userIDStr, _ := c.Get("user_id")
	ownerID, err := primitive.ObjectIDFromHex(userIDStr.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Check if user owns the property
	property, err := h.propertyService.GetPropertyByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Property not found"})
		return
	}

	if property.OwnerID != ownerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized to delete this property"})
		return
	}

	if err := h.propertyService.DeleteProperty(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete property"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Property deleted successfully"})
}
