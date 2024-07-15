# Exit on any error
$ErrorActionPreference = "Stop"

$env:REPO = "OpenStruct/goeasycli"
# Function to fetch the latest tag using GitHub API
function Fetch-LatestTag {
    param (
        [string]$RepoUrl
    )
    
    $apiUrl = "https://api.github.com/repos/$RepoUrl/tags"
    
    try {
        $response = Invoke-RestMethod -Uri $apiUrl -ErrorAction Stop
        if ($response.Count -gt 0) {
            $script:CLI_VERSION = $response[0].name
            Write-Host "Latest tag found: $CLI_VERSION"
        } else {
            Write-Host "Error: No tags found in the repository." -ForegroundColor Red
            exit 1
        }
    } catch {
        Write-Host "Error: Unable to fetch tags from the repository." -ForegroundColor Red
        exit 1
    }
}

# Check if Git is available
if (Get-Command git -ErrorAction SilentlyContinue) {
    $repoUrl = "https://github.com/$env:REPO.git"
    $gitOutput = git ls-remote --tags --refs --sort="version:refname" $repoUrl 2>$null
    
    if ($LASTEXITCODE -eq 0 -and $gitOutput) {
        $script:CLI_VERSION = ($gitOutput -split '/')[-1]
        Write-Host "Latest tag found (via git): $CLI_VERSION"
    } else {
        Write-Host "Git command failed. Falling back to API method."
        Fetch-LatestTag $env:REPO
    }
} else {
    Write-Host "Git is not installed. Using API method to fetch latest tag."
    Fetch-LatestTag $env:REPO
}

# Use CLI_VERSION here...
Write-Host "Final CLI_VERSION: $CLI_VERSION"

$BASE_URL = "https://github.com/OpenStruct/goeasycli/releases/download/$CLI_VERSION"

# Determine the architecture
$ARCH = if ([System.Environment]::Is64BitOperatingSystem) { "amd64" } else { "386" }

# Construct the download URL
$URL = "$BASE_URL/goeasycli_${CLI_VERSION}_windows_${ARCH}.zip"


$TEMP_DIR = [System.IO.Path]::GetTempPath() + [System.IO.Path]::GetRandomFileName()
New-Item -ItemType Directory -Path $TEMP_DIR | Out-Null
Write-Host "Downloading $URL..."
$webClient = New-Object System.Net.WebClient
$webClient.DownloadFile($URL, "$TEMP_DIR\goeasycli.zip")

# Extract the zip file
Add-Type -AssemblyName System.IO.Compression.FileSystem
[System.IO.Compression.ZipFile]::ExtractToDirectory("$TEMP_DIR\goeasycli.zip", $TEMP_DIR)

# Move the binary to a directory in PATH
Move-Item -Path "$TEMP_DIR\goeasycli.exe" -Destination "$env:ProgramFiles\goeasycli\goeasycli.exe"
[Environment]::SetEnvironmentVariable("Path", [Environment]::GetEnvironmentVariable("Path", [System.EnvironmentVariableTarget]::Machine) + ";$env:ProgramFiles\goeasycli", [System.EnvironmentVariableTarget]::Machine)

# Clean up
Remove-Item -Path "$TEMP_DIR" -Recurse

# Verify installation
if (Get-Command goeasycli -ErrorAction SilentlyContinue) {
    Write-Host "goeasycli installed successfully!"
} else {
    Write-Host "Installation failed."
    exit 1
}