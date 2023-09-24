package main

import (
	"context"
	"fmt"
	"os"

	"dagger.io/dagger"
)

func main() {
	if err := build(context.Background()); err != nil {
		fmt.Println(err)
	}
}

func build(ctx context.Context) error {
	fmt.Println("Building with Dagger")

	// define build matrix
	oses := []string{"linux", "darwin"}
	arches := []string{"amd64", "arm64"}

	// initialize Dagger client
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
	if err != nil {
		return err
	}
	defer client.Close()

	// get reference to the local project
	src := client.Host().Directory(".")

	// create empty directory to put build outputs
	outputs := client.Directory()

	// get `golang` image
	golang := client.
		Container().
		From("golang:latest").
		WithExec([]string{"go", "install", "github.com/caddyserver/xcaddy/cmd/xcaddy@latest"})

	// mount cloned repository into `golang` image
	golang = golang.WithDirectory("./", src).WithWorkdir("./")

	for _, goos := range oses {
		for _, goarch := range arches {
			// create a directory for each os and arch
			path := fmt.Sprintf("build/%s/%s/", goos, goarch)

			// set GOARCH and GOOS in the build environment
			build := golang.WithEnvVariable("GOOS", goos)
			build = build.WithEnvVariable("GOARCH", goarch)

			// build application caddy with extensions
			build = build.WithExec([]string{
				"xcaddy",
				"build",
				"--with",
				"github.com/ttnesby/slack-block-builder/caddy-ext/azalertslacknotification",
				"--with",
				"github.com/corazawaf/coraza-caddy/v2",
				"--with",
				"github.com/mholt/caddy-ratelimit",
				// "-o",
				// path,
			})

			// get reference to build output directory in container
			outputs = outputs.WithDirectory(path, build.Directory(path))
		}
	}
	// write build artifacts to host
	_, err = outputs.Export(ctx, ".")
	if err != nil {
		return err
	}

	return nil
}
