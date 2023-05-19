package main

import (
	"docker/cmd"
	"github.com/docker/docker/client"
)

func main() {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	// list all containers
	cmd.CheckContainers(cli)
	// run container
	cmd.RunImage(cli)
	cmd.CheckContainers(cli)
	// stop container
	cmd.StopAllContainer(cli)
	// inspect container
	cmd.Inspect(cli)
}
