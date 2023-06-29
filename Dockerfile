#WARNING: DO NOT RUN THIS FILE DIRECTORY. RATHER GO TO directory server and execute install-connect.sh "sudo bash install-connect.sh"
# Stage 1: Build the React application
FROM node:latest AS react-builder

WORKDIR /app

# Copy the React source code
COPY ./web .

# Install dependencies and build the React application
RUN npm install
RUN npm run build

# Stage 2: Build the Go application
FROM golang:latest AS go-builder

WORKDIR /go/src/app

# Copy the Go source code
COPY ./server .

# Build the Go application
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .

# Stage 3: Create the final production image
FROM alpine:18

# Set the working directory
WORKDIR /app

# Copy the built React application from the react-builder stage
COPY --from=react-builder /app/build ./web

# Copy the Go executable from the go-builder stage
COPY --from=go-builder /go/src/app/server .

RUN mkdir ./cert
# Copy the certificate and private key
COPY ./cert ./cert

# Expose a port if your Go application listens on a specific port
EXPOSE 443

# Define the command to run when the container starts
CMD ["./server"]
