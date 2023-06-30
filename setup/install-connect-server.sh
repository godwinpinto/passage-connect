#!/bin/bash

# Define paths
outside_dir="../../cert"
inside_dir="../cert"
dockerfile_dir="../"

# Copy the cert directory to the Docker build context
cp -r "$outside_dir" "$inside_dir"

cp $outside_dir/.env ../web/.env

# Navigate to the directory containing the Dockerfile
cd "$dockerfile_dir"

# Build the Docker image
docker build -t passage-connect .

# Remove the copied cert directory
rm -rf cert

rm web/.env

echo "Setup completed successfully!"