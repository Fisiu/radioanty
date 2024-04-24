FROM golang:1.22-alpine
WORKDIR /app

COPY go.mod ./
RUN go mod download && go mod verify

COPY main.go ./
RUN go build -o radioanty

# Expose port for the application
EXPOSE 8088

# Run the compiled Go binary
CMD ["/app/radioanty"]