#!/bin/sh

echo "Pulling latest code from Git..."
git pull

echo "Pulling latest Docker images..."
docker compose pull

echo "Stopping running containers..."
docker compose down

echo "Rebuilding and starting containers..."
docker compose up -d --build

echo "Cleaning up unused Docker images..."
docker image prune -f

echo "✅ Deployment complete!"
