FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o /app/bin/server ./bin/server/main.go
RUN go build -o /app/bin/worker ./bin/worker/main.go

RUN chmod +x /app/bin/server
RUN chmod +x /app/bin/worker

EXPOSE 3003

CMD ["/app/bin/server", "/app/bin/worker"]
