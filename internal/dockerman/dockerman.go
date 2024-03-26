package dockerman

import (
    "fmt"
    "context"
    "github.com/docker/docker/api/types"
    // "github.com/docker/docker/api/types/container"
    "github.com/docker/docker/client"
    "github.com/docker/docker/pkg/archive"
)

// build image from docker file
func buildAdvocateImage(cli *client.Client) (types.ImageBuildResponse, error) {
    ctx := context.Background()

    buildContext, err := archive.Tar("./", archive.Uncompressed)

    if err != nil {
        fmt.Println("Error building tar for build context")
        return types.ImageBuildResponse{}, err
    }

    buildOptions := types.ImageBuildOptions{
        Dockerfile: "internal/dockerman/Dockerfile",
        Tags:       []string{"gotesting1"},
    }
    
    resp, err := cli.ImageBuild(ctx, buildContext, buildOptions)

    if err != nil {
        return types.ImageBuildResponse{}, err
    }

    return resp, nil
}

// create container from image
// func createAdvocateContainer() error {

// }

// run image
// func startAdvocateContainer() error {

// }

func StartAdvocate() error {
    // Initialize our docker client
    cli, err := client.NewClientWithOpts(client.FromEnv)
    if err != nil {
        return err
    }
    defer cli.Close()

    resp, err := buildAdvocateImage(cli)

    if err != nil {
        fmt.Println("Error building image")
        return err
    }

    fmt.Println(resp.Body)

    return nil
}