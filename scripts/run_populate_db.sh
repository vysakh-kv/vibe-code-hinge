#!/bin/bash

echo "Vibe Dating App - Database Population Script"
echo "==========================================="

# Check if python3 is installed
if ! command -v python3 &> /dev/null; then
    echo "Python 3 is required but not installed. Please install Python 3 and try again."
    exit 1
fi

# Create virtual environment if it doesn't exist
if [ ! -d "venv" ]; then
    echo "Creating virtual environment..."
    python3 -m venv venv
fi

# Activate virtual environment
echo "Activating virtual environment..."
source venv/bin/activate

# Install required packages
echo "Installing required packages..."
pip install requests

# Run the script
echo "Running database population script..."
python3 scripts/populate_db.py

# Deactivate virtual environment
deactivate

echo "==========================================="
echo "Database population process complete!" 