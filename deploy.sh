#!/bin/bash

echo "=== SunPanel Docker Deployment ==="
echo "1. Checking Docker socket permissions..."
# Ensure the user has access to docker socket, or at least warn
if [ ! -r /var/run/docker.sock ]; then
    echo "Warning: /var/run/docker.sock not found or not readable."
    echo "Make sure Docker is running."
fi

echo "2. Building and Starting Containers..."
# Force build to include new changes
docker-compose up -d --build

echo "=== Deployment Complete ==="
echo "Access the panel at http://localhost:3002"
