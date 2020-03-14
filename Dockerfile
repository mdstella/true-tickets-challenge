FROM golang:1.13-alpine

# This Dockerfile is optimized for go binaries, change it as much as necessary
# for your language of choice.
# Set the Current Working Directory inside the container
WORKDIR /true-tickets-challenge

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

EXPOSE 9091

CMD ["./main"]