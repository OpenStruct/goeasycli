# Ensure the script runs with elevated permissions
function Ensure-Elevated {
    if (-not ([Security.Principal.WindowsPrincipal][Security.Principal.WindowsIdentity]::GetCurrent()).IsInRole([Security.Principal.WindowsBuiltInRole] "Administrator")) {
        Start-Process powershell -ArgumentList "-NoProfile -ExecutionPolicy Bypass -File `"$PSCommandPath`"" -Verb RunAs
        exit
    }
}

Ensure-Elevated

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
        Write-Host $_.Exception.Message
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
$ARCH = if ([System.Environment]::Is64BitOperatingSystem) { "x86_64" } else { "i386" }

# Construct the download URL
$URL = "$BASE_URL/goeasycli_windows_${ARCH}.zip"

$TEMP_DIR = Join-Path ([System.IO.Path]::GetTempPath()) ([System.IO.Path]::GetRandomFileName())
New-Item -ItemType Directory -Path $TEMP_DIR -Force | Out-Null
Write-Host "Downloading $URL..."

try {
    $webClient = New-Object System.Net.WebClient
    $webClient.DownloadFile($URL, "$TEMP_DIR\goeasycli.zip")
} catch {
    Write-Host "Error downloading file: $_"
    exit 1
}

# Extract the zip file
try {
    Add-Type -AssemblyName System.IO.Compression.FileSystem
    [System.IO.Compression.ZipFile]::ExtractToDirectory("$TEMP_DIR\goeasycli.zip", $TEMP_DIR)
} catch {
    Write-Host "Error extracting zip file: $_"
    exit 1
}

# Create the destination directory if it doesn't exist
$destinationDir = "$env:ProgramFiles\goeasycli"
if (!(Test-Path -Path $destinationDir)) {
    New-Item -ItemType Directory -Path $destinationDir -Force | Out-Null
}

# Move the binary to a directory in PATH
try {
    Move-Item -Path "$TEMP_DIR\goeasycli.exe" -Destination "$destinationDir\goeasycli.exe" -Force
    Write-Host "Successfully moved goeasycli.exe to $destinationDir"
} catch {
    Write-Host "Error moving goeasycli.exe: $_"
    exit 1
}

# Add to PATH if not already there
$currentPath = [Environment]::GetEnvironmentVariable("Path", [System.EnvironmentVariableTarget]::Machine)
if ($currentPath -notlike "*$destinationDir*") {
    try {
        [Environment]::SetEnvironmentVariable("Path", "$currentPath;$destinationDir", [System.EnvironmentVariableTarget]::Machine)
        Write-Host "Added $destinationDir to system PATH"
    } catch {
        Write-Host "Error updating PATH: $_"
        Write-Host "You may need to manually add $destinationDir to your system PATH"
    }
} else {
    Write-Host "$destinationDir is already in system PATH"
}

# Clean up
Remove-Item -Path $TEMP_DIR -Recurse -Force -ErrorAction SilentlyContinue

# Verify installation
if (Get-Command goeasycli -ErrorAction SilentlyContinue) {
    Write-Host "goeasycli installed successfully!"
} else {
    Write-Host "Installation failed. Please check if $destinationDir\goeasycli.exe exists and is in your PATH."
    exit 1
}
