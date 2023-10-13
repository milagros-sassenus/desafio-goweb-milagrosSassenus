FROM golang:1.19.0

WORKDIR /usr/src/app

# Copy only necessary files and directories
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the application
RUN go build -o app ./cmd/main.go

CMD ["./app"]
