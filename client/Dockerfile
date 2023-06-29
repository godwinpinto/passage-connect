# Use the Amazon Linux 2 image as the base
FROM amazonlinux:latest as builder

# Install any necessary build dependencies for Go and C development
RUN yum install -y gcc glibc-static tar gzip pam-devel

# Define any additional runtime dependencies for the C PAM module
#RUN yum install -y libpam

RUN mkdir /goinstall

WORKDIR /goinstall
# Install Go
RUN curl -O https://dl.google.com/go/go1.19.2.linux-amd64.tar.gz
RUN tar -xf go1.19.2.linux-amd64.tar.gz
ENV PATH=$PATH:/goinstall/go/bin


# Set the working directory inside the container
WORKDIR /app

# Copy your Go and C source files to the container
COPY . .

# Build the Go binary
RUN make


FROM httpd:latest

RUN mkdir /goinstall

WORKDIR /goinstall

COPY --from=builder /app/passage_connect.so .

# Expose the container port
EXPOSE 80


#RUN cp passage_connect.so /build/

# Use a lightweight base image for the final stage
#FROM amazonlinux:latest

# Set the working directory inside the container
#WORKDIR /app

# Copy the built Go binary from the builder stage
#COPY --from=builder /app/app .

# Copy the built C PAM module from the builder stage
#COPY --from=builder /app/pam_module.so .

# Define any additional runtime dependencies for the C PAM module
#RUN yum install -y libpam

# Set the entrypoint to your Go binary
#ENTRYPOINT ["./app"]
