FROM golang:1.21-alpine as build

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/simtix-ticketing ./bin/server/.


FROM alpine:3.18.4

WORKDIR /app

COPY --from=build /app/simtix-ticketing /app/simtix-ticketing

# Update alpine
RUN apk update --no-cache && \
    apk add --no-cache bash && \
    # Set permission
    chmod +x /app/simtix-ticketing

# Run app
CMD ["/app/simtix-ticketing"]