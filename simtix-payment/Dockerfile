FROM golang:latest

WORKDIR /go/src/app

COPY . .

RUN go build -o /go/src/app/bin/server ./bin/server/main.go
RUN go build -o /go/src/app/bin/worker ./bin/worker/main.go

RUN chmod +x /go/src/app/bin/server
RUN chmod +x /go/src/app/bin/worker

EXPOSE 3003

CMD ["./bin/server", "./bin/worker"]