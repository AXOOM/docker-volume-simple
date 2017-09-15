Param ([string]$Version = "0.1-dev")
$ErrorActionPreference = "Stop"
Push-Location $(Split-Path -Path $MyInvocation.MyCommand.Definition -Parent)

$env:DOCKER_REGISTRY = if ($Version.Contains("-")) {"docker-ci.axoom.cloud"} else {"docker.axoom.cloud"}
$env:VERSION = $Version
docker-compose up --build

Pop-Location
