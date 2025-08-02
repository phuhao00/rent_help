---
applyTo: '**'
---

# Rent Help Application - AI Coding Instructions

## Project Overview

**Rent Help** is a full-stack web application designed to help users manage rental properties, tenant information, and rental processes. The application provides a comprehensive platform for landlords and tenants to interact and manage rental-related activities.

### Technology Stack

**Frontend**:
- React 18.2.0 with TypeScript 4.9.5
- Next.js for server-side rendering and routing
- Material-UI v5 (@mui/material) for UI components
- React Router v6 for client-side routing
- Port: 3001 (development)

**Backend**:
- Go 1.21 with Gin framework
- Air v1.49.0 for hot reload during development
- JWT for authentication
- Port: 8081 (development), 8080 (internal)

**Database & Cache**:
- MongoDB 7.0 (Primary database)
- Redis 7-alpine (Caching and session management)
- MongoDB Port: 27018 (host) -> 27017 (container)
- Redis Port: 6380 (host) -> 6379 (container)

**Development Environment**:
- Docker & Docker Compose for containerization
- Air for Go hot reload
- Volume mounting for development

## Architecture Guidelines

### Project Structure
```
rent_help/
├── frontend/          # React/Next.js frontend
├── backend/           # Go backend API
├── docker/           # Docker configuration files
├── .github/          # GitHub workflows and instructions
└── docker-compose.dev.yml  # Development orchestration
```

### Backend Structure (Go)
```
backend/
├── cmd/server/       # Application entry point
├── internal/
│   ├── config/       # Configuration management
│   ├── handlers/     # HTTP handlers/controllers
│   ├── middleware/   # HTTP middleware
│   ├── models/       # Data models
│   └── services/     # Business logic
├── pkg/              # Shared packages
│   ├── database/     # Database connections
│   ├── errors/       # Error handling
│   ├── utils/        # Utility functions
│   └── validation/   # Input validation
└── uploads/          # File upload storage
```

### Frontend Structure (React)
```
frontend/
├── src/
│   ├── components/   # Reusable UI components
│   ├── pages/        # Next.js pages
│   ├── hooks/        # Custom React hooks
│   ├── services/     # API service calls
│   ├── utils/        # Utility functions
│   └── types/        # TypeScript type definitions
├── public/           # Static assets
└── styles/           # CSS/styling files
```

## Coding Standards

### Go Backend Guidelines

**Naming Conventions**:
- Use camelCase for variables and functions
- Use PascalCase for exported functions and types
- Use snake_case for database field names
- Package names should be lowercase, single words

**Code Organization**:
- Group related functionality in packages
- Keep handlers thin, move business logic to services
- Use dependency injection for testability
- Follow Clean Architecture principles

**Error Handling**:
- Always handle errors explicitly
- Use custom error types when appropriate
- Log errors with context
- Return meaningful HTTP status codes

**Database Interactions**:
- Use MongoDB aggregation pipelines for complex queries
- Implement proper connection pooling
- Use transactions for multi-document operations
- Validate input before database operations

**Example Go Code Style**:
```go
// Handler function
func (h *UserHandler) CreateUser(c *gin.Context) {
    var req CreateUserRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    user, err := h.userService.CreateUser(req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(http.StatusCreated, user)
}

// Model definition
type User struct {
    ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
    FullName string            `json:"fullName" bson:"full_name"`
    Email    string            `json:"email" bson:"email"`
    Password string            `json:"-" bson:"password"`
    CreatedAt time.Time        `json:"createdAt" bson:"created_at"`
}
```

### React Frontend Guidelines

**Component Structure**:
- Use functional components with hooks
- Implement TypeScript interfaces for props
- Use Material-UI components consistently
- Keep components small and focused

**State Management**:
- Use React hooks (useState, useEffect, etc.)
- Consider Context API for global state
- Implement custom hooks for reusable logic

**API Integration**:
- Create service functions for API calls
- Handle loading and error states
- Use proper TypeScript types for API responses

**Example React Component Style**:
```tsx
interface UserProfileProps {
  userId: string;
}

const UserProfile: React.FC<UserProfileProps> = ({ userId }) => {
  const [user, setUser] = useState<User | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchUser = async () => {
      try {
        const userData = await userService.getUser(userId);
        setUser(userData);
      } catch (err) {
        setError('Failed to load user');
      } finally {
        setLoading(false);
      }
    };

    fetchUser();
  }, [userId]);

  if (loading) return <CircularProgress />;
  if (error) return <Alert severity="error">{error}</Alert>;

  return (
    <Card>
      <CardContent>
        <Typography variant="h5">{user?.fullName}</Typography>
        <Typography variant="body2">{user?.email}</Typography>
      </CardContent>
    </Card>
  );
};
```

## Environment Configuration

### Environment Variables

**Backend (.env)**:
```
PORT=8080
MONGODB_URI=mongodb://admin:password123@mongodb:27017/renthelp?authSource=admin
JWT_SECRET=your-secret-key
DB_NAME=rent_help
REDIS_URL=redis://redis:6379
UPLOAD_PATH=/app/uploads
GIN_MODE=debug
```

**Frontend**:
```
NEXT_PUBLIC_API_URL=http://localhost:8081
NEXT_PUBLIC_APP_NAME=Rent Help
```

### Development Setup

**Required Tools**:
- Docker Desktop 27.5.1+
- Go 1.21+
- Node.js 18+
- Git

**Development Commands**:
```bash
# Start all services
docker-compose -f docker-compose.dev.yml up -d

# Stop all services
docker-compose -f docker-compose.dev.yml down

# View logs
docker-compose -f docker-compose.dev.yml logs -f [service-name]

# Rebuild services
docker-compose -f docker-compose.dev.yml up -d --build
```

## API Design Guidelines

### RESTful Endpoints
- Use proper HTTP methods (GET, POST, PUT, DELETE)
- Follow RESTful URL patterns: `/api/v1/users/{id}`
- Use consistent response formats
- Implement proper HTTP status codes

### Authentication
- Use JWT tokens for authentication
- Implement middleware for protected routes
- Include user context in requests

### Error Responses
```json
{
  "error": "User not found",
  "code": "USER_NOT_FOUND",
  "message": "The requested user could not be found",
  "timestamp": "2025-08-02T10:30:00Z"
}
```

### Success Responses
```json
{
  "data": {...},
  "message": "Operation completed successfully",
  "timestamp": "2025-08-02T10:30:00Z"
}
```

## Database Guidelines

### MongoDB Schema Design
- Use embedded documents for one-to-few relationships
- Use references for one-to-many relationships
- Implement proper indexing for frequently queried fields
- Use aggregation pipelines for complex queries

### Redis Usage
- Cache frequently accessed data
- Store session information
- Implement rate limiting data

## Testing Guidelines

### Backend Testing
- Write unit tests for services and handlers
- Use table-driven tests in Go
- Mock external dependencies
- Aim for 80%+ code coverage

### Frontend Testing
- Write component tests with React Testing Library
- Test user interactions and API integration
- Use Jest for unit tests
- Implement E2E tests for critical paths

## Security Guidelines

### Authentication & Authorization
- Hash passwords using bcrypt
- Validate JWT tokens on protected routes
- Implement role-based access control
- Use HTTPS in production

### Input Validation
- Validate all user inputs
- Sanitize data before database operations
- Implement rate limiting
- Use CORS properly

### Data Protection
- Never log sensitive information
- Encrypt sensitive data at rest
- Use environment variables for secrets
- Implement proper session management

## Performance Guidelines

### Backend Optimization
- Use connection pooling for database
- Implement caching strategies
- Use pagination for large datasets
- Optimize database queries

### Frontend Optimization
- Implement code splitting
- Use lazy loading for components
- Optimize images and assets
- Implement proper error boundaries

## Deployment Guidelines

### Production Considerations
- Use production-ready Docker images
- Implement health checks
- Set up monitoring and logging
- Use environment-specific configurations
- Implement graceful shutdown

### Docker Best Practices
- Use multi-stage builds
- Minimize image size
- Use specific version tags
- Implement proper health checks

## Common Issues & Solutions

### Known Fixes Applied
1. **MongoDB Health Check**: Use `mongosh` instead of deprecated `mongo`
2. **Environment Variables**: Use `MONGODB_URI` (not `MONGO_URI`)
3. **Field Mapping**: Use `FullName` field (not `Name`) in User model
4. **React Dependencies**: Use React Router v6 syntax with Material-UI v5

### Troubleshooting
- Check Docker container logs for errors
- Verify environment variable configuration
- Ensure proper network connectivity between services
- Check port availability and conflicts

## AI Assistant Guidelines

When working on this project:

1. **Always consider the full stack** - Changes in backend may require frontend updates
2. **Maintain consistency** - Follow established patterns and conventions
3. **Test thoroughly** - Verify changes work in the Docker environment
4. **Document changes** - Update relevant documentation and comments
5. **Security first** - Consider security implications of all changes
6. **Performance aware** - Consider performance impact of modifications
7. **Error handling** - Implement proper error handling and user feedback
8. **Type safety** - Use TypeScript types and Go interfaces properly

## Version Information

- **Last Updated**: August 2, 2025
- **Go Version**: 1.21
- **React Version**: 18.2.0
- **MongoDB Version**: 7.0
- **Redis Version**: 7-alpine
- **Docker Compose**: Development environment ready