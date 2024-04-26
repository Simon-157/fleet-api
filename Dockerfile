FROM golang:1.22.2 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o fleet-api -a -installsuffix cgo cmd/fleet-api/main.go

FROM scratch

WORKDIR /app

COPY --from=builder /app/fleet-api .

# Expose port 8080
EXPOSE 8080

# Command to run the executable
CMD ["./fleet-api"]
