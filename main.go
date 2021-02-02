package main

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v33/github"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Not enough arguments")
		os.Exit(1)
	}
	name := os.Args[1]
	repo := os.Args[2]

	client := github.NewClient(nil)
	ctx := context.Background()
	releases, _, err := client.Repositories.ListReleases(ctx, name, repo, &github.ListOptions{})
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}
	for _, release := range releases {
		fmt.Printf("Version: %s - %s\n", release.GetTagName(), release.GetName())
		for _, asset := range release.Assets {
			fmt.Printf("  %s - %d\n", asset.GetName(), asset.GetDownloadCount())
		}
	}
}
