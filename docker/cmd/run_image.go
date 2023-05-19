package cmd

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"io"
	"os"
)

// docker run -d -p 1880:1880 nodered/node-red

const imageName = "nodered/node-red"

func RunImage(cli *client.Client) {
	ctx := context.Background()
	out, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	defer out.Close()
	_, _ = io.Copy(os.Stdout, out)

	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			"1880/tcp": []nat.PortBinding{
				{
					HostIP:   "127.0.0.1",
					HostPort: "",
				},
			},
		},
	}

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: imageName,
		ExposedPorts: nat.PortSet{
			"1880/tcp": struct{}{},
		},
	}, hostConfig, nil, nil, "")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	fmt.Printf("created container: %s \n", resp.ID)
}
