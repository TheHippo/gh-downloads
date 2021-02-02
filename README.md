# gh-downloads
List download number for Github release assets

## Building

Install in GOBIN:

    go install -ldflags="-s -w" -trimpath github.com/TheHippo/gh-downloads

Build locally:

	go install -ldflags="-s -w" -trimpath github.com/TheHippo/gh-downloads

## Usage

    gh-downloads USERNAME REPO

or:

	gh-downloads ORGANISATION REPO