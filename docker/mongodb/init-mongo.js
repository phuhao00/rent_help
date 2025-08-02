// MongoDB initialization script for RentHelp
print('Initializing RentHelp database...');

// Switch to the renthelp database
db = db.getSiblingDB('renthelp');

// Create collections with validation
db.createCollection('users', {
  validator: {
    $jsonSchema: {
      bsonType: 'object',
      required: ['email', 'password', 'full_name', 'role'],
      properties: {
        email: {
          bsonType: 'string',
          description: 'must be a string and is required'
        },
        password: {
          bsonType: 'string',
          description: 'must be a string and is required'
        },
        full_name: {
          bsonType: 'string',
          description: 'must be a string and is required'
        },
        role: {
          bsonType: 'string',
          enum: ['tenant', 'landlord', 'admin'],
          description: 'must be one of: tenant, landlord, admin'
        }
      }
    }
  }
});

db.createCollection('properties', {
  validator: {
    $jsonSchema: {
      bsonType: 'object',
      required: ['title', 'description', 'type', 'price', 'address', 'owner_id'],
      properties: {
        title: {
          bsonType: 'string',
          description: 'must be a string and is required'
        },
        type: {
          bsonType: 'string',
          enum: ['apartment', 'house', 'condo', 'townhouse', 'studio'],
          description: 'must be one of the property types'
        },
        price: {
          bsonType: 'number',
          minimum: 0,
          description: 'must be a positive number'
        }
      }
    }
  }
});

db.createCollection('bookings');
db.createCollection('reviews');
db.createCollection('messages');
db.createCollection('notifications');

// Create indexes for better performance
print('Creating indexes...');

// User indexes
db.users.createIndex({ email: 1 }, { unique: true });
db.users.createIndex({ role: 1 });
db.users.createIndex({ created_at: -1 });

// Property indexes
db.properties.createIndex({ owner_id: 1 });
db.properties.createIndex({ type: 1, price: 1 });
db.properties.createIndex({ 'location.city': 1 });
db.properties.createIndex({ available: 1, created_at: -1 });
db.properties.createIndex({ 'location.coordinates': '2dsphere' });

// Booking indexes
db.bookings.createIndex({ property_id: 1 });
db.bookings.createIndex({ tenant_id: 1 });
db.bookings.createIndex({ landlord_id: 1 });
db.bookings.createIndex({ start_date: 1, end_date: 1 });
db.bookings.createIndex({ status: 1, created_at: -1 });

// Review indexes
db.reviews.createIndex({ property_id: 1 });
db.reviews.createIndex({ reviewer_id: 1 });
db.reviews.createIndex({ reviewee_id: 1 });

// Message indexes
db.messages.createIndex({ sender_id: 1, receiver_id: 1 });
db.messages.createIndex({ booking_id: 1 });
db.messages.createIndex({ created_at: -1 });

// Notification indexes
db.notifications.createIndex({ user_id: 1, is_read: 1 });
db.notifications.createIndex({ created_at: -1 });

print('RentHelp database initialization completed!');
print('Collections created: users, properties, bookings, reviews, messages, notifications');
print('Indexes created for optimal performance');
