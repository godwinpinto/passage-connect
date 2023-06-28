#!/bin/bash

# Create a folder "passage-connect" outside connect-server
mkdir ../../passage-connect

# Copy passage_connect_server into the new directory
cp -r -p ../server/passage_connect_server ../../passage-connect/

# Copy everything from build into the new directory and rename to "web"
cp -r -p ../web/build ../../passage-connect/web

# Prompt user for PASSAGE_CONNECT_URL
read -p "Enter PASSAGE_CONNECT_URL: " passage_connect_url

# Prompt user for PASSAGE_APP_ID
read -p "Enter PASSAGE_APP_ID: " passage_app_id

# Prompt user for PASSAGE_API_KEY
read -p "Enter PASSAGE_API_KEY: " passage_api_key

# Create a config.properties file and save the properties
echo "PASSAGE_CONNECT_URL=$passage_connect_url" > ../../passage-connect/config.properties
echo "PASSAGE_APP_ID=$passage_app_id" >> ../../passage-connect/config.properties
echo "PASSAGE_API_KEY=$passage_api_key" >> ../../passage-connect/config.properties

echo "Setup completed successfully!"