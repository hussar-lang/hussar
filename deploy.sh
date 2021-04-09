#!/bin/sh

export GIT_COMMIT=$(git rev-list -1 HEAD)
export GIT_TAG=$(git describe --tags)
gox -osarch="windows/amd64 darwin/amd64 linux/amd64 linux/arm" -output "dist/hussar_{{.OS}}-{{.Arch}}" -ldflags "-X main.GitCommit=$GIT_COMMIT" -ldflags "-X main.VersionString=$GIT_TAG"
ghr --username kscarlett --token "$GITHUB_TOKEN" --replace --prerelease --debug pre-release dist/