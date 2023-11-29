FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o /app/bin/server ./cmd/server/server.go
RUN go build -o /app/bin/worker ./cmd/worker/worker.go

# Install supervisor
RUN apt-get update && apt-get install -y supervisor

# Copy the supervisor configuration file into the container
COPY supervisord.conf /etc/supervisor/conf.d/supervisord.conf

EXPOSE 3002

CMD ["/usr/bin/supervisord"]
