package cmd

import (
	"context"
	"fmt"
	"github.com/docker/docker/client"
)

func Inspect(cli *client.Client) {
	ins, err := cli.ContainerInspect(context.Background(), "5c0b1c1da7864917836a806886475166aa0a2ceefad7735394238af8f616392e")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", ins.Config.ExposedPorts)
	fmt.Printf("%+v\n", ins.HostConfig.PortBindings)

	// ok
	fmt.Printf("%+v\n", ins.NetworkSettings.Ports)
}
