// Demo module for the Dagger Cloud Namespace demo: one check that compiles a
// small Go program and returns the binary as the check's artifact.
package main

import (
	"context"
	"fmt"

	"dagger/demo/internal/dagger"
)

type Demo struct {
	// +private
	Source *dagger.Directory
}

func New(ws *dagger.Workspace) *Demo {
	return &Demo{
		Source: ws.Directory("/hello"),
	}
}

// Compile the hello program and return the binary.
func (m *Demo) Build() *dagger.File {
	return dag.Container().
		From("golang:1.24-alpine").
		WithDirectory("/src", m.Source).
		WithWorkdir("/src").
		WithEnvVariable("CGO_ENABLED", "0").
		WithExec([]string{"go", "build", "-o", "/out/hello", "."}).
		File("/out/hello")
}

// Compile the hello program and run it.
//
// +check
func (m *Demo) Hello(ctx context.Context) error {
	out, err := dag.Container().
		From("alpine:3.21").
		WithFile("/bin/hello", m.Build()).
		WithExec([]string{"/bin/hello"}).
		Stdout(ctx)
	if err != nil {
		return err
	}
	fmt.Println(out)
	return nil
}
