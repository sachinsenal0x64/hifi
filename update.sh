#!/bin/sh

set -e  # Exit immediately if a command exits with a non-zero status.

echo "Pulling latest code from Git..."
sudo git pull

echo "Pulling latest Docker images..."
sudo docker compose pull

echo "Stopping running containers..."
sudo docker compose down

echo "Rebuilding and starting containers..."
sudo docker compose up -d --build

echo "Cleaning up unused Docker images..."
sudo docker image prune -f

echo "✅ Deployment complete!"
