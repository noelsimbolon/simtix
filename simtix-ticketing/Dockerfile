FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o /app/bin/server ./cmd/server/server.go
RUN go build -o /app/bin/worker ./cmd/worker/worker.go

EXPOSE 3002

CMD ["./run.sh"]