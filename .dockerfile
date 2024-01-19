# Step 1: Use an official Go runtime as a parent image
FROM golang:1.18-alpine

# Step 2: Set the working directory inside the container
WORKDIR /app

# Step 3: Copy the local package files to the container's workspace
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Step 4: Copy the source code
COPY . .

# Step 5: Build the Go app
RUN go build -o myapp

# Step 6: Expose port 8080 to the outside world
EXPOSE 8080

# Step 7: Run the executable
CMD ["./myapp"]
