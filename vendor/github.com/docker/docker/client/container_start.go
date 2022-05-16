package client // import "github.com/docker/docker/client"

import (
	"context"
	"net/url"

	"github.com/docker/docker/api/types"
	"github.com/docker/go-connections/nat"
)

type portConfigWrapper struct {
	ExposedPorts	*map[nat.Port]struct{}
	PortBindings    *map[nat.Port][]nat.PortBinding
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
			ExposedPorts:	&options.ExposedPorts,
			PortBindings:	&options.PortBindings,
		}

		//Debug statements
    	//fmt.Println("\n\n\n(container_start.go)Sending the address for ExposedPorts, but this is ExposedPorts: ", options.ExposedPorts)
    	//fmt.Println("\n\n\n(container_start.go)Sending the address for PortBindings, but this is PortBindings: ", options.PortBindings)

		resp, err := cli.post(ctx, "/containers/"+containerID+"/start", query, body, nil)
		ensureReaderClosed(resp)
		return err
	}

	//Debug statement
	//fmt.Println("\n\n\n(in container_start.go)Sending 'nil' body as part of POST request for ContainerStart: ", )

	resp, err := cli.post(ctx, "/containers/"+containerID+"/start", query, nil, nil)
	ensureReaderClosed(resp)
	return err
}
