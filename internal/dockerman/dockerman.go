package dockerman

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func CreateAdvocateContainer() error {
	// Create a new Docker client
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return err
	}

	// Pull the Docker image (optional)
	imageName := "golang:latest"
	err = pullImage(cli, imageName)
	if err != nil {
		return err
	}

	// Define container configuration
	containerConfig := &container.Config{
		Image: imageName
	}

	// Create the container
	resp, err := cli.ContainerCreate(context.Background(), containerConfig, nil, nil, nil, "")
	if err != nil {
		return err
	}

	// Start the container
	err = cli.ContainerStart(context.Background(), resp.ID, container.StartOptions{})
	if err != nil {
		return err
	}

	fmt.Printf("Container %s started successfully.\n", resp.ID)

    return nil
}

func pullImage(cli *client.Client, imageName string) error {
	out, err := cli.ImagePull(context.Background(), imageName, types.ImagePullOptions{})
	if err != nil {
		return err
	}
	defer out.Close()

	// Print the pull status
	buf := make([]byte, 1024)
	for {
		_, err := out.Read(buf)
		if err != nil {
			break
		}
	}

	return nil
}