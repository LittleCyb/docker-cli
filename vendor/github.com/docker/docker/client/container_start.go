package client // import "github.com/docker/docker/client"

import (
	"context"
	"net/url"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/go-connections/nat"
)

type portConfigWrapper struct {
	ExposedPorts	map[nat.Port]struct{}
	PortBindings    map[nat.Port][]nat.PortBinding
}

// ContainerStart sends a request to the docker daemon to start a container.
func (cli *Client) ContainerStart(ctx context.Context, containerID string, options types.ContainerStartOptions) error {
	query := url.Values{}
	if len(options.CheckpointID) != 0 {
		query.Set("checkpoint", options.CheckpointID)
	}
	if len(options.CheckpointDir) != 0 {
		query.Set("checkpoint-dir", options.CheckpointDir)
	}

	var body portConfigWrapper

	if len(options.ExposedPorts) != 0 && len(options.PortBindings) != 0 {
		body = portConfigWrapper{
			ExposedPorts:	options.ExposedPorts,
			PortBindings:	options.PortBindings,
		}
	}

	fmt.Println("\n\n\nThis is the body: %+v\n\n\n", body)

	resp, err := cli.post(ctx, "/containers/"+containerID+"/start", query, nil, nil)
	ensureReaderClosed(resp)
	return err
}
