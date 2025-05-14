#!/bin/bash

echo "Vibe Dating App - Postman Collection Generator"
echo "==========================================="

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "Go is required but not installed. Please install Go and try again."
    exit 1
fi

# Go to the scripts directory
cd "$(dirname "$0")"

# Build and run the Go script
echo "Building and running the Postman collection generator..."
go run generate_postman_collection.go

# Check if the collection was generated
if [ -f "../vibe_dating_app.postman_collection.json" ]; then
    echo "Success! The Postman collection has been generated at:"
    echo "$(dirname $(dirname "$0"))/vibe_dating_app.postman_collection.json"
    echo ""
    echo "To import this collection into Postman:"
    echo "1. Open Postman"
    echo "2. Click 'Import' in the top left"
    echo "3. Choose 'File' and select the generated JSON file"
    echo "4. Click 'Import'"
    echo ""
    echo "Don't forget to set up the following environment variables in Postman:"
    echo "- base_url: http://localhost:8082/api/v1"
    echo "- user_id: (a valid user ID from your system)"
    echo "- profile_id: (a valid profile ID from your system)"
    echo "- match_id: (a valid match ID from your system)"
else
    echo "Error: Failed to generate the Postman collection."
    exit 1
fi

echo "===========================================" 