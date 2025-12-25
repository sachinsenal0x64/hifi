#!/bin/sh

# Check if jq is installed
if ! command -v jq &> /dev/null
then
    echo "Error: jq is not installed. Please install jq and try again."
    exit 1
fi

JSON_URL="https://gist.githubusercontent.com/sachinsenal0x64/0b7945a4a0df4c77ecf3bc7b7a7ee2f5/raw/167abddc3923799aef143e2a863bcf3db8314959/hifi-managed.json"

JSON=$(curl -s "$JSON_URL")

if ! echo "$JSON" | jq empty 2> /dev/null; then
    echo "Error: JSON is invalid. Please fix the JSON file."
    exit 1
fi

# Export environment variables
export MANAGE_HOST=$(echo "$JSON" | jq -r '.MANAGE_HOST')
export MODE=$(echo "$JSON" | jq -r '.MODE')
export TIDAL_HOST=$(echo "$JSON" | jq -r '.TIDAL_HOST[0]')
export SCHEME=$(echo "$JSON" | jq -r '.SCHEME')
export CLIENT_ID=$(echo "$JSON" | jq -r '.CLIENT_ID[0]')

# Optional: print to verify
echo "MANAGE_HOST=$MANAGE_HOST"
echo "MODE=$MODE"
echo "TIDAL_HOST=$TIDAL_HOST"
echo "SCHEME=$SCHEME"
echo "CLIENT_ID=$CLIENT_ID"

