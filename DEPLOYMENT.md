# Deployment Configuration

This document describes deployment configurations for different environments.

## Docker Deployment

### Development Environment

```bash
# Start development environment
docker-compose -f docker-compose.dev.yml up -d

# View logs
docker-compose logs -f

# Stop services
docker-compose down
```

### Production Environment

```bash
# Build production images
docker-compose -f docker-compose.prod.yml build

# Start production environment
docker-compose -f docker-compose.prod.yml up -d

# Scale services
docker-compose -f docker-compose.prod.yml up -d --scale backend=3
```

## Kubernetes Deployment

### Prerequisites
- Kubernetes cluster (1.18+)
- kubectl configured
- Ingress controller (nginx-ingress)

### Deploy to Kubernetes

```bash
# Create namespace
kubectl create namespace renthelp

# Apply configurations
kubectl apply -f k8s/

# Check deployment status
kubectl get pods -n renthelp

# Get service URLs
kubectl get ingress -n renthelp
```

## Environment Variables

### Backend (.env)
```bash
# Server Configuration
PORT=8080
GIN_MODE=release

# Database Configuration
MONGODB_URI=mongodb://mongodb:27017/renthelp
MONGODB_DATABASE=renthelp

# JWT Configuration
JWT_SECRET=your-super-secret-jwt-key-here
JWT_EXPIRES_IN=24h

# File Upload Configuration
UPLOAD_PATH=/app/uploads
MAX_FILE_SIZE=10MB

# Email Configuration
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USERNAME=your-email@gmail.com
SMTP_PASSWORD=your-app-password

# Redis Configuration (optional)
REDIS_URL=redis://redis:6379

# External Services
STRIPE_SECRET_KEY=sk_test_...
STRIPE_WEBHOOK_SECRET=whsec_...
CLOUDINARY_URL=cloudinary://...

# Monitoring
SENTRY_DSN=https://...
```

### Frontend (.env.local)
```bash
# API Configuration
NEXT_PUBLIC_API_URL=http://localhost:8080
NEXT_PUBLIC_WEBSOCKET_URL=ws://localhost:8080

# Authentication
NEXTAUTH_URL=http://localhost:3000
NEXTAUTH_SECRET=your-nextauth-secret

# External Services
NEXT_PUBLIC_GOOGLE_MAPS_API_KEY=AIza...
NEXT_PUBLIC_STRIPE_PUBLISHABLE_KEY=pk_test_...

# Analytics
NEXT_PUBLIC_GA_ID=G-...
NEXT_PUBLIC_HOTJAR_ID=...

# Feature Flags
NEXT_PUBLIC_ENABLE_CHAT=true
NEXT_PUBLIC_ENABLE_PAYMENTS=true
```

## CI/CD Pipeline

### GitHub Actions

```yaml
# .github/workflows/deploy.yml
name: Deploy to Production

on:
  push:
    branches: [main]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      
      - name: Setup Go
        uses: actions/setup-go@v3
        with:
          go-version: 1.21
          
      - name: Run backend tests
        run: |
          cd backend
          go test ./...
          
      - name: Setup Node.js
        uses: actions/setup-node@v3
        with:
          node-version: 18
          
      - name: Run frontend tests
        run: |
          cd frontend
          npm ci
          npm test

  deploy:
    needs: test
    runs-on: ubuntu-latest
    if: github.ref == 'refs/heads/main'
    
    steps:
      - uses: actions/checkout@v3
      
      - name: Deploy to production
        run: |
          echo "Deploying to production..."
          # Add deployment commands here
```

## Monitoring and Logging

### Prometheus Configuration

```yaml
# prometheus.yml
global:
  scrape_interval: 15s

scrape_configs:
  - job_name: 'renthelp-backend'
    static_configs:
      - targets: ['backend:8080']
    metrics_path: /metrics
    
  - job_name: 'renthelp-frontend'
    static_configs:
      - targets: ['frontend:3000']
    metrics_path: /api/metrics
```

### Grafana Dashboards

1. **Application Metrics**
   - Request rate and response times
   - Error rates and status codes
   - Database connection pool status

2. **Business Metrics**
   - User registrations
   - Property listings
   - Booking conversions

3. **Infrastructure Metrics**
   - CPU and memory usage
   - Disk and network I/O
   - Container health status

## Security Configuration

### SSL/TLS Setup

```nginx
# nginx.conf
server {
    listen 80;
    server_name renthelp.com www.renthelp.com;
    return 301 https://$server_name$request_uri;
}

server {
    listen 443 ssl http2;
    server_name renthelp.com www.renthelp.com;

    ssl_certificate /etc/ssl/certs/renthelp.crt;
    ssl_certificate_key /etc/ssl/private/renthelp.key;
    
    ssl_protocols TLSv1.2 TLSv1.3;
    ssl_ciphers ECDHE-RSA-AES256-GCM-SHA512:DHE-RSA-AES256-GCM-SHA512;
    ssl_prefer_server_ciphers off;
    
    location / {
        proxy_pass http://frontend:3000;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
    
    location /api {
        proxy_pass http://backend:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
}
```

### Firewall Rules

```bash
# UFW configuration
sudo ufw allow ssh
sudo ufw allow 80/tcp
sudo ufw allow 443/tcp
sudo ufw deny 3000/tcp  # Block direct frontend access
sudo ufw deny 8080/tcp  # Block direct backend access
sudo ufw enable
```

## Backup Strategy

### Database Backup

```bash
#!/bin/bash
# backup.sh

DATE=$(date +%Y%m%d_%H%M%S)
BACKUP_DIR="/backups"
MONGO_CONTAINER="renthelp_mongodb_1"

# Create backup
docker exec $MONGO_CONTAINER mongodump --db renthelp --out /tmp/backup_$DATE

# Copy backup from container
docker cp $MONGO_CONTAINER:/tmp/backup_$DATE $BACKUP_DIR/

# Upload to cloud storage (AWS S3)
aws s3 cp $BACKUP_DIR/backup_$DATE s3://renthelp-backups/ --recursive

# Clean old backups (keep last 30 days)
find $BACKUP_DIR -name "backup_*" -mtime +30 -delete
```

### File Storage Backup

```bash
#!/bin/bash
# files-backup.sh

DATE=$(date +%Y%m%d_%H%M%S)
UPLOAD_DIR="/app/uploads"
BACKUP_DIR="/backups/files"

# Sync files to backup directory
rsync -av --delete $UPLOAD_DIR/ $BACKUP_DIR/

# Upload to cloud storage
aws s3 sync $BACKUP_DIR s3://renthelp-files-backup/
```

## Performance Optimization

### Backend Optimizations

1. **Database Indexing**
```javascript
// MongoDB indexes
db.properties.createIndex({ location: "2dsphere" })
db.properties.createIndex({ type: 1, price: 1 })
db.properties.createIndex({ owner_id: 1, created_at: -1 })
db.bookings.createIndex({ property_id: 1, start_date: 1, end_date: 1 })
db.users.createIndex({ email: 1 }, { unique: true })
```

2. **Redis Caching**
```go
// Cache frequently accessed data
func (s *PropertyService) GetProperty(id string) (*models.Property, error) {
    // Try cache first
    if cached := s.cache.Get("property:" + id); cached != nil {
        return cached.(*models.Property), nil
    }
    
    // Fallback to database
    property, err := s.db.GetProperty(id)
    if err != nil {
        return nil, err
    }
    
    // Cache for 1 hour
    s.cache.Set("property:"+id, property, time.Hour)
    return property, nil
}
```

### Frontend Optimizations

1. **Code Splitting**
```typescript
// Dynamic imports for large components
const PropertyMap = dynamic(() => import('./PropertyMap'), {
  loading: () => <MapSkeleton />,
  ssr: false
})
```

2. **Image Optimization**
```typescript
// Next.js Image component with optimization
import Image from 'next/image'

<Image
  src={property.images[0]}
  alt={property.title}
  width={800}
  height={600}
  placeholder="blur"
  blurDataURL="data:image/jpeg;base64,..."
  priority={index < 3}
/>
```

## Troubleshooting

### Common Issues

1. **Container Memory Issues**
```bash
# Check container resources
docker stats

# Increase memory limits
# In docker-compose.yml
services:
  backend:
    mem_limit: 512m
    mem_reservation: 256m
```

2. **Database Connection Issues**
```bash
# Check MongoDB status
docker exec renthelp_mongodb_1 mongo --eval "db.runCommand({connectionStatus: 1})"

# Check logs
docker logs renthelp_mongodb_1
```

3. **SSL Certificate Issues**
```bash
# Check certificate expiry
openssl x509 -in /etc/ssl/certs/renthelp.crt -text -noout | grep "Not After"

# Renew Let's Encrypt certificate
certbot renew --nginx
```

### Health Checks

```yaml
# docker-compose.yml health checks
services:
  backend:
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost:8080/health"]
      interval: 30s
      timeout: 10s
      retries: 3
      
  frontend:
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost:3000/api/health"]
      interval: 30s
      timeout: 10s
      retries: 3
      
  mongodb:
    healthcheck:
      test: ["CMD", "mongo", "--eval", "db.runCommand({ping: 1})"]
      interval: 30s
      timeout: 10s
      retries: 3
```

## Scaling Considerations

### Horizontal Scaling

1. **Load Balancing**
```nginx
# nginx upstream configuration
upstream backend_servers {
    server backend1:8080;
    server backend2:8080;
    server backend3:8080;
}

upstream frontend_servers {
    server frontend1:3000;
    server frontend2:3000;
}
```

2. **Database Sharding**
```javascript
// MongoDB sharding configuration
sh.enableSharding("renthelp")
sh.shardCollection("renthelp.properties", { "location.city": 1 })
sh.shardCollection("renthelp.bookings", { "created_at": 1 })
```

### Vertical Scaling

```yaml
# docker-compose.yml resource limits
services:
  backend:
    deploy:
      resources:
        limits:
          cpus: '2.0'
          memory: 2G
        reservations:
          cpus: '1.0'
          memory: 1G
```
