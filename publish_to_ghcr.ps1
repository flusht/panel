$ErrorActionPreference = "Stop"
Write-Host "=== Publish SunPanel to GitHub Container Registry (GHCR) ==="

# 1. Input
$username = Read-Host "Enter your GitHub Username"
if ([string]::IsNullOrWhiteSpace($username)) { Write-Error "Username required" }

# GHCR requires image name to be lowercase
$imageName = "sun-panel-custom"
$tag = "latest"
# Full format: ghcr.io/owner/image:tag
$fullImageName = "ghcr.io/$($username.ToLower())/$imageName:$tag"

Write-Host "Target Image: $fullImageName"

# 2. Build
Write-Host "Building..."
docker build -t $fullImageName .

# 3. Reminder
Write-Host "IMPORTANT: You must login to GHCR before pushing."
Write-Host "Command: docker login ghcr.io -u $username -p <YOUR_PAT_TOKEN>"
Write-Host "(Password must be a Personal Access Token with 'write:packages' scope)"
Write-Host "Press Enter to continue if you are already logged in..."
Pause

# 4. Push
Write-Host "Pushing..."
docker push $fullImageName

Write-Host "=== Success ==="
Write-Host "Image: $fullImageName"
Write-Host ""
Write-Host "Update docker-compose.yml to use:"
Write-Host "  image: $fullImageName"
