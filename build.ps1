Param ([string]$Version = "0.1-dev")
Push-Location $(Split-Path -Path $MyInvocation.MyCommand.Definition -Parent)

$DockerRegistry = if ($Version.Contains("-")) {"docker-ci.axoom.cloud"} else {"docker.axoom.cloud"}

echo "Building plugin image"
$ImageId = docker build -q .

echo "Removing old plugin version (if exists)"
docker plugin rm "$DockerRegistry/docker-volume-simple:$Version"

echo "Building & installing new plugin version"
docker run --rm --volume /var/run/docker.sock:/var/run/docker.sock $ImageId docker plugin create "$DockerRegistry/docker-volume-simple:$Version" /plugin

Pop-Location
