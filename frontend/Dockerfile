# Multi-stage Dockerfile for Next.js frontend

# Development stage
FROM node:18-alpine AS development

WORKDIR /app

# Copy package files
COPY package.json ./

# Install all dependencies (including dev dependencies)
RUN npm install

# Copy source code
COPY . .

# Expose port
EXPOSE 3000

# Default command for development
CMD ["npm", "run", "dev"]

# Install dependencies stage for production
FROM node:18-alpine AS deps
WORKDIR /app

# Copy package files
COPY package.json ./

# Install dependencies
RUN npm install --only=production

# Build stage
FROM node:18-alpine AS builder
WORKDIR /app

# Copy dependencies
COPY --from=deps /app/node_modules ./node_modules

# Copy source code
COPY . .

# Build the application
RUN npm run build

# Production stage
FROM node:18-alpine AS production
WORKDIR /app

ENV NODE_ENV production

RUN addgroup --system --gid 1001 nodejs
RUN adduser --system --uid 1001 nextjs

# Copy built application
COPY --from=builder /app/public ./public
COPY --from=builder --chown=nextjs:nodejs /app/.next/standalone ./
COPY --from=builder --chown=nextjs:nodejs /app/.next/static ./.next/static

USER nextjs

EXPOSE 3000

ENV PORT 3000

CMD ["node", "server.js"]

ENV PORT 3000

CMD ["node", "server.js"]
