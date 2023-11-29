# Use the latest golang base image
FROM golang:latest

# Set the working directory in the container
WORKDIR /app

# Copy the current directory contents into the container
COPY . .

# Build the Go app
RUN go build -o /app/bin/server ./cmd/server/server.go
RUN go build -o /app/bin/worker ./cmd/worker/worker.go

# Install supervisord
RUN apt-get update && apt-get install -y supervisor

# Copy the supervisor configuration file into the container
COPY supervisord.conf /etc/supervisor/conf.d/supervisord.conf

# Expose port 3003
EXPOSE 3003

# Run supervisord
CMD ["/usr/bin/supervisord"]
