package dockerman

import (
	"fmt"
    "context"
    "os"
	"archive/tar"
    "github.com/docker/docker/api/types"
    "github.com/docker/docker/api/types/container"
    "github.com/docker/docker/client"
    "github.com/docker/docker/pkg/stdcopy"
)

func CreateAdvocateDocker() error {
	fmt.Println("crash 1")
    ctx := context.Background()
    cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
    if err != nil {
        return err
    }
	fmt.Println("crash 2")

    // Path to Dockerfile
    dockerBuildContext, err := os.Open(".")
    if err != nil {
        return err
    }
    defer dockerBuildContext.Close()

	fmt.Println("crash 3a")

    // Build the Docker image
    buildOptions := types.ImageBuildOptions{
        Dockerfile: "Dockerfile", // assuming Dockerfile is in the current directory
        Remove:     true,         // Remove intermediate containers after a successful build
    }
    buildResponse, err := cli.ImageBuild(ctx, dockerBuildContext, buildOptions)
    if err != nil {
		fmt.Println("response error here")
        return err
    }
    defer buildResponse.Body.Close()

	fmt.Println("crash 3")

    // Optionally, output the build response to stdout
    _, err = stdcopy.StdCopy(os.Stdout, os.Stderr, buildResponse.Body)
    if err != nil {
        return err
    }

    // Create a container from the built image
    containerConfig := &container.Config{
        Image: "testing_advocate",
    }
    containerResp, err := cli.ContainerCreate(ctx, containerConfig, nil, nil, nil, "")
    if err != nil {
        return err
    }

    // Start the container
    if err := cli.ContainerStart(ctx, containerResp.ID, container.StartOptions{}); err != nil {
        return err
    }

    // The container is now running
    return nil
}
