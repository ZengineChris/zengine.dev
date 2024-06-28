// A generated module for Zengine functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"fmt"
)

type Zengine struct{}

func (m *Zengine) Hello() string {
	return "Hello from ghere"
}

// Returns a container that echoes whatever string argument is provided
func (m *Zengine) ContainerEcho(stringArg string) *Container {
	return dag.Container().From("alpine:latest").WithExec([]string{"echo", stringArg})
}

// Returns lines that match a pattern in the files of the provided Directory
func (m *Zengine) GrepDir(ctx context.Context, directoryArg *Directory, pattern string) (string, error) {
	return dag.Container().
		From("alpine:latest").
		WithMountedDirectory("/mnt", directoryArg).
		WithWorkdir("/mnt").
		WithExec([]string{"grep", "-R", pattern, "."}).
		Stdout(ctx)
}

func (m *Zengine) Release(ctx context.Context, dir *Directory) string {

	version, err := dag.Nsv(dir).Next(ctx)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(version)

	foo, err := dag.
		Buildx().
		Build(ctx, dir, "ghcr.io/zenginedev/website:0.0.1")

	if err != nil {
		fmt.Println(err)
	}
	// build a docker container

	fmt.Println(foo)

	// store the next version in a var
	return version

}

func (m *Zengine) BuildImage(ctx context.Context, dir *Directory, token *Secret) (string, error) {
	ops := DockerBuildPublishOpts{Tags: []string{"0.0.4"}}

	return dag.Docker(DockerOpts{
		Registry: "ghcr.io",
		Username: "ZengineChris",
		Password: token,
	}).
		Build(dir).
		Publish(ctx, "ghcr.io/zenginedev/website", ops)
}
