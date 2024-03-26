package dockerman

import (
    "archive/tar"
    "bytes"
    "context"
    "io"
    "os"
    "path/filepath"

    "github.com/docker/docker/api/types"
    "github.com/docker/docker/api/types/container"
    "github.com/docker/docker/client"
)

func CreateAdvocateContainer() error {
    ctx := context.Background()
    cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
    if err != nil {
        return err
    }

    // Create a buf to hold the tarred data to copy
    buf := new(bytes.Buffer)
    tw := tar.NewWriter(buf)
    defer tw.Close()

    // Walk the directory and tar each file
    dirPath := "./advocate-2"
    filepath.Walk(dirPath, func(file string, fi os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        // create a new dir/file header
        header, err := tar.FileInfoHeader(fi, file)
        if err != nil {
            return err
        }

        header.Name = filepath.ToSlash(file)

        // write the header
        if err := tw.WriteHeader(header); err != nil {
            return err
        }

        // if not a dir, write file content
        if !fi.IsDir() {
            data, err := os.Open(file)
            if err != nil {
                return err
            }
            if _, err := io.Copy(tw, data); err != nil {
                return err
            }
            data.Close()
        }
        return nil
    })

    tw.Close()

    // Use the Docker SDK to create and start a container
    resp, err := cli.ContainerCreate(ctx, &container.Config{
        Image: "golang:latest",
        WorkingDir: "/app",
        // Entrypoint:   []string{"go", "run", "/app/main.go"},
    }, nil, nil, nil, "")
    
    if err != nil {
        return err
    }

    // Copy the buffer to the container
    err = cli.CopyToContainer(ctx, resp.ID, "/app", buf, types.CopyToContainerOptions{})
        
    if err != nil {
        return err
    }

    // Start the container
    if err := cli.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
        return err
    }

    return nil
}