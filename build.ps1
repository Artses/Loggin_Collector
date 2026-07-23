param (
    [string]$Action = "build"
)

$BinaryName = "log-collector.exe"
$CmdDir = ".\cmd\main.go"

switch ($Action) {
    "build" {
        Write-Host "Building native binary for Windows..." -ForegroundColor Cyan
        go build -o $BinaryName $CmdDir
        if ($LASTEXITCODE -eq 0) {
            Write-Host "Build complete: .\$BinaryName" -ForegroundColor Green
        } else {
            Write-Host "Build failed!" -ForegroundColor Red
        }
    }
    "run" {
        Write-Host "Building and running..." -ForegroundColor Cyan
        go build -o $BinaryName $CmdDir
        if ($LASTEXITCODE -eq 0) {
            Write-Host "Running .\$BinaryName..." -ForegroundColor Green
            .\$BinaryName
        } else {
            Write-Host "Build failed! Cannot run." -ForegroundColor Red
        }
    }
    "clean" {
        Write-Host "Cleaning up..." -ForegroundColor Cyan
        if (Test-Path $BinaryName) {
            Remove-Item $BinaryName
            Write-Host "Cleanup complete." -ForegroundColor Green
        } else {
            Write-Host "Nothing to clean." -ForegroundColor Yellow
        }
    }
    "test" {
        Write-Host "Running tests..." -ForegroundColor Cyan
        go test ./...
    }
    default {
        Write-Host "Unknown action: $Action" -ForegroundColor Red
        Write-Host "Usage: .\build.ps1 [build | run | clean | test]" -ForegroundColor Yellow
    }
}
