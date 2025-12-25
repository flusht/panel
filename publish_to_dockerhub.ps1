$ErrorActionPreference = "Stop"

Write-Host "=== Publish SunPanel to Docker Hub ==="

# 1. Get User Input
$username = Read-Host "Enter your Docker Hub Username"
if ([string]::IsNullOrWhiteSpace($username)) {
    Write-Error "Username is required."
}

$imageName = "sun-panel-custom"
$tag = "latest"
$fullImageName = "$username/$imageName:$tag"

Write-Host "Target Image: $fullImageName"

# 2. Build
Write-Host "Building image..."
docker build -t $fullImageName .

# 3. Login check (optional, user might be logged in)
Write-Host "Please ensure you are logged in to Docker Hub."
Write-Host "Run 'docker login' if you haven't already."
Pause

# 4. Push
Write-Host "Pushing image to Docker Hub..."
docker push $fullImageName

Write-Host "=== Success ==="
Write-Host "Your image is now available at: $fullImageName"
Write-Host ""
Write-Host "To use this image, update your docker-compose.yml:"
Write-Host "  image: $fullImageName"
Write-Host "  # build: .  <-- Comment or remove this line"
