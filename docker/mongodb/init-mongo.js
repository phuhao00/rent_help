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

// Insert sample data if collections are empty
print('Checking if sample data insertion is needed...');

// Sample Users
if (db.users.countDocuments() === 0) {
  print('Inserting sample users...');
  
  const sampleUsers = [
    {
      _id: ObjectId(),
      email: 'admin@renthelp.com',
      password: '$2b$12$LQv3c1yqBwqQ4QxQ.3K.oeqJ5UQzQJ5P5QhP5QhP5QhP5QhP5QhP5Q', // password: admin123
      full_name: 'Administrator',
      role: 'admin',
      phone: '+1-555-0100',
      avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=admin',
      is_verified: true,
      is_active: true,
      created_at: new Date(),
      updated_at: new Date()
    },
    {
      _id: ObjectId(),
      email: 'john.landlord@email.com',
      password: '$2b$12$LQv3c1yqBwqQ4QxQ.3K.oeqJ5UQzQJ5P5QhP5QhP5QhP5QhP5QhP5Q', // password: landlord123
      full_name: 'John Smith',
      role: 'landlord',
      phone: '+1-555-0101',
      avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=john',
      is_verified: true,
      is_active: true,
      created_at: new Date(),
      updated_at: new Date()
    },
    {
      _id: ObjectId(),
      email: 'jane.tenant@email.com',
      password: '$2b$12$LQv3c1yqBwqQ4QxQ.3K.oeqJ5UQzQJ5P5QhP5QhP5QhP5QhP5QhP5Q', // password: tenant123
      full_name: 'Jane Doe',
      role: 'tenant',
      phone: '+1-555-0102',
      avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=jane',
      is_verified: true,
      is_active: true,
      created_at: new Date(),
      updated_at: new Date()
    },
    {
      _id: ObjectId(),
      email: 'mike.landlord@email.com',
      password: '$2b$12$LQv3c1yqBwqQ4QxQ.3K.oeqJ5UQzQJ5P5QhP5QhP5QhP5QhP5QhP5Q', // password: landlord123
      full_name: 'Mike Johnson',
      role: 'landlord',
      phone: '+1-555-0103',
      avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=mike',
      is_verified: true,
      is_active: true,
      created_at: new Date(),
      updated_at: new Date()
    }
  ];
  
  const userResult = db.users.insertMany(sampleUsers);
  print(`Inserted ${userResult.insertedIds.length} sample users`);
  
  // Store user IDs for reference
  const adminId = sampleUsers[0]._id;
  const johnId = sampleUsers[1]._id;
  const janeId = sampleUsers[2]._id;
  const mikeId = sampleUsers[3]._id;
  
  // Sample Properties
  if (db.properties.countDocuments() === 0) {
    print('Inserting sample properties...');
    
    const sampleProperties = [
      {
        _id: ObjectId(),
        title: 'Modern Downtown Apartment',
        description: 'Beautiful 2-bedroom apartment in the heart of downtown. Walking distance to restaurants, shops, and public transportation.',
        type: 'apartment',
        price: 1200,
        bedrooms: 2,
        bathrooms: 1,
        area: 850,
        address: {
          street: '123 Main Street',
          city: 'Downtown',
          state: 'CA',
          zip_code: '90210',
          country: 'USA'
        },
        location: {
          type: 'Point',
          coordinates: [-118.2437, 34.0522] // Los Angeles coordinates
        },
        amenities: ['parking', 'gym', 'laundry', 'balcony'],
        images: [
          'https://images.unsplash.com/photo-1522708323590-d24dbb6b0267?w=800',
          'https://images.unsplash.com/photo-1493809842364-78817add7ffb?w=800'
        ],
        owner_id: johnId,
        available: true,
        available_from: new Date(),
        lease_terms: ['6_months', '12_months'],
        pets_allowed: false,
        smoking_allowed: false,
        utilities_included: ['water', 'heating'],
        created_at: new Date(),
        updated_at: new Date()
      },
      {
        _id: ObjectId(),
        title: 'Cozy Studio Near University',
        description: 'Perfect for students! Fully furnished studio apartment near the university campus with all utilities included.',
        type: 'studio',
        price: 800,
        bedrooms: 0,
        bathrooms: 1,
        area: 450,
        address: {
          street: '456 University Ave',
          city: 'College Town',
          state: 'CA',
          zip_code: '90211',
          country: 'USA'
        },
        location: {
          type: 'Point',
          coordinates: [-118.2537, 34.0622]
        },
        amenities: ['wifi', 'furnished', 'laundry', 'study_area'],
        images: [
          'https://images.unsplash.com/photo-1586023492125-27b2c045efd7?w=800',
          'https://images.unsplash.com/photo-1560448204-e02f11c3d0e2?w=800'
        ],
        owner_id: mikeId,
        available: true,
        available_from: new Date(),
        lease_terms: ['3_months', '6_months', '12_months'],
        pets_allowed: false,
        smoking_allowed: false,
        utilities_included: ['water', 'electricity', 'internet', 'heating'],
        created_at: new Date(),
        updated_at: new Date()
      },
      {
        _id: ObjectId(),
        title: 'Spacious Family House',
        description: 'Large 4-bedroom house perfect for families. Features a big backyard, garage, and quiet neighborhood.',
        type: 'house',
        price: 2500,
        bedrooms: 4,
        bathrooms: 3,
        area: 2200,
        address: {
          street: '789 Oak Street',
          city: 'Suburbia',
          state: 'CA',
          zip_code: '90212',
          country: 'USA'
        },
        location: {
          type: 'Point',
          coordinates: [-118.2637, 34.0722]
        },
        amenities: ['garage', 'backyard', 'fireplace', 'dishwasher', 'laundry'],
        images: [
          'https://images.unsplash.com/photo-1568605114967-8130f3a36994?w=800',
          'https://images.unsplash.com/photo-1570129477492-45c003edd2be?w=800'
        ],
        owner_id: johnId,
        available: true,
        available_from: new Date(Date.now() + 30 * 24 * 60 * 60 * 1000), // Available in 30 days
        lease_terms: ['12_months', '24_months'],
        pets_allowed: true,
        smoking_allowed: false,
        utilities_included: ['water'],
        created_at: new Date(),
        updated_at: new Date()
      }
    ];
    
    const propertyResult = db.properties.insertMany(sampleProperties);
    print(`Inserted ${propertyResult.insertedIds.length} sample properties`);
    
    // Sample Bookings
    print('Inserting sample bookings...');
    
    const sampleBookings = [
      {
        _id: ObjectId(),
        property_id: sampleProperties[0]._id,
        tenant_id: janeId,
        landlord_id: johnId,
        start_date: new Date(Date.now() + 7 * 24 * 60 * 60 * 1000), // In 7 days
        end_date: new Date(Date.now() + 372 * 24 * 60 * 60 * 1000), // 1 year lease
        monthly_rent: 1200,
        security_deposit: 2400,
        status: 'confirmed',
        lease_terms: '12_months',
        special_requests: 'Pet-friendly if possible',
        created_at: new Date(),
        updated_at: new Date()
      }
    ];
    
    const bookingResult = db.bookings.insertMany(sampleBookings);
    print(`Inserted ${bookingResult.insertedIds.length} sample bookings`);
    
    // Sample Reviews
    print('Inserting sample reviews...');
    
    const sampleReviews = [
      {
        _id: ObjectId(),
        property_id: sampleProperties[0]._id,
        reviewer_id: janeId,
        reviewee_id: johnId,
        rating: 5,
        title: 'Excellent landlord and property!',
        comment: 'John is very responsive and the apartment is exactly as described. Great location and well-maintained.',
        review_type: 'property',
        created_at: new Date(Date.now() - 30 * 24 * 60 * 60 * 1000), // 30 days ago
        updated_at: new Date(Date.now() - 30 * 24 * 60 * 60 * 1000)
      }
    ];
    
    const reviewResult = db.reviews.insertMany(sampleReviews);
    print(`Inserted ${reviewResult.insertedIds.length} sample reviews`);
  }
} else {
  print('Sample users already exist, skipping data insertion');
}

print('Sample data initialization completed!');
print('---------------------------------------------');
print('Sample Login Credentials:');
print('Admin: admin@renthelp.com / admin123');
print('Landlord: john.landlord@email.com / landlord123');
print('Landlord: mike.landlord@email.com / landlord123');
print('Tenant: jane.tenant@email.com / tenant123');
print('---------------------------------------------');
