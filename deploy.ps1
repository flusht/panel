Write-Host "=== SunPanel Docker Deployment ==="
Write-Host "1. Building and Starting Containers..."

# Force build to include new changes
docker-compose up -d --build

Write-Host "=== Deployment Complete ==="
Write-Host "Access the panel at http://localhost:3002"
