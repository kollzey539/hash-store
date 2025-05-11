FROM golang:1.24

WORKDIR /app

# Copy everything and build
COPY . .

RUN go mod tidy && go build -o hash-store .

EXPOSE 8080

CMD ["./hash-store"]
