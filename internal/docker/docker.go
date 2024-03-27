package docker

import (
    "fmt"
    "io"
    "context"
    "github.com/docker/docker/api/types"
    "github.com/docker/docker/api/types/container"
    "github.com/docker/docker/client"
    "github.com/docker/docker/pkg/archive"
)

const (
    // golang:latest packaged with advocate-2 app
    // created from internal/dockerman/Dockerfile
    imageName = "advocate_app"

    // preview app container name
    containerName = "advocate_preview"

    // context dir for docker daemon to create our image 
    buildCtxDir = "./"

    // path for Dockerfile in build context dir
    dockerfilePath = "internal/docker/Dockerfile"

    // exposed port on advocate container
    port = ":8090"
)

type Docker struct {
    ctx context.Context
    cli *client.Client
}

func NewDockerClient() (*Docker, error) {
    cli, err := client.NewClientWithOpts(client.FromEnv)
    if err != nil {
        fmt.Println("There was an error creating docker client.")
        return err
    }

    return &Docker {
        ctx: context.Background(),
        cli: cli
    }
}

func (d *Docker) CloseDockerClient() {
    if d.cli != nil {
        d.cli.Close()
    }
}

func (d *Docker) getContainerIpAddr(containerID string) (string, error) {
    containerJSON, err := d.cli.ContainerInspect(d.ctx, containerID)
    if err != nil {
        fmt.Println("Error with inspecting container")
        return "", err
    }
    
    // Assuming the container is using the default "bridge" network.
    ipAddr := containerJSON.NetworkSettings.Networks["bridge"].IPAddress 
    return ipAddr, nil
}

// Removes preview container if it exists
func (d *Docker) removeAdvocateContainer() error {
    containers, err := d.cli.ContainerList(d.ctx, container.ListOptions{All: true})

    if err != nil {
        fmt.Println("Error trying to get container list")
        return err
    }

    removeOpts := container.RemoveOptions{Force: true}

    for _, container := range containers {
        for _, name := range container.Names {
            if name == "/"+containerName {
                err = d.cli.ContainerRemove(d.ctx, container.ID, removeOpts)    
                if err != nil {
                    fmt.Println("issue trying to remove container")
                    return err
                }

                fmt.Println("Advocate container removed.")
            } 
        }
    }

    return nil
}

// build preview advocate image
func (d *Docker) buildAdvocateImage() (types.ImageBuildResponse, error) {
    buildContext, err := archive.Tar(buildCtxDir, archive.Uncompressed)
    if err != nil {
        fmt.Println("Error building tar for build context")
        return types.ImageBuildResponse{}, err
    }

    buildOptions := types.ImageBuildOptions{
        Dockerfile: dockerfilePath,
        Tags:       []string{imageName},
    }
    
    resp, err := d.cli.ImageBuild(d.ctx, buildContext, buildOptions)

    if err != nil {
        fmt.Println("Error creating advocate image")
        return types.ImageBuildResponse{}, err
    }

    return resp, nil
}

// create container from image
func (d *Docker) createAdvocateContainer() (container.CreateResponse, error) {
    resp, err := d,cli.ContainerCreate(
        d.ctx,
        &container.Config{
            Image: imageName,
        },
        nil, nil, nil, containerName,
    )

    if err != nil {
        fmt.Println("Error creating container.")
        return container.CreateResponse{}, err
    }

    return resp, nil
}

// run image
func (d *Docker) startAdvocateContainer(containerID string) error {
    return d.cli.ContainerStart(d.ctx, containerID,  container.StartOptions{})
}

func (d *Docker) StartAdvocatePreview() error {

    // package latest go with advocate-2
    fmt.Println("Building advocate image...")
    resp, err := d.buildAdvocateImage()

    if err != nil {
        fmt.Println("Error building image.")
        return err
    }

    defer resp.Body.Close()

    // get response as string to print
    // bodyContent, err := io.ReadAll(resp.Body) // debugging
    _, err = io.ReadAll(resp.Body)

    if err != nil {
        fmt.Println("Error getting image build response.")
        return err
    }
    // fmt.Println(string(bodyContent))

    fmt.Println("Advocate image created!")

    // remove any existing advocate-2 container apps
    err = d.removeAdvocateContainer()
    if  err != nil {
        fmt.Println("Error trying to remove container")
    }

    // create app container from advocate image
    fmt.Println("Creating advocate container.")
    createResp, err := d.createAdvocateContainer()

    if err != nil {
        fmt.Println("Create advocate container fail.")
        return err
    }

    fmt.Println("Container created!")

    fmt.Println("Starting preview container...")

    err = d.startAdvocateContainer(createResp.ID)

    if err != nil {
        fmt.Println("Error starting advocate container")
        return err
    }

    fmt.Println("Container started! id: ", createResp.ID)

    ipAddr, err := d.getContainerIpAddr(createResp.ID)

    if err != nil {
        fmt.Println("Error trying to get container ip address")
        return err
    }

    fmt.Println("Container running @", ipAddr, port)

    return nil
}