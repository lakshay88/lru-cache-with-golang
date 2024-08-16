# Frontend build step
FROM node:18-alpine as build-frontend

WORKDIR /app/frontend
COPY lru-cache-frontend/ .
RUN npm install && npm run build

# Backedn Build step
FROM golang:1.20-alpine as build-backend

WORKDIR /app/backend
COPY go-backend/ .
RUN go build -o lru-cache-server main.go
FROM alpine:latest
WORKDIR /app

# Copy Go binary and React build from previous stages
COPY --from=build-backend /app/backend/lru-cache-server .
COPY --from=build-frontend /app/frontend/build ./frontend

EXPOSE 1234

# Set environment variables
ENV REACT_APP_API_URL=http://localhost:1234

# Start the Go server
CMD ["./lru-cache-server"]
