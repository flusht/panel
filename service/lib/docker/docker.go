package docker

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type ContainerInfo struct {
	ID      string   `json:"id"`
	Names   []string `json:"names"`
	Image   string   `json:"image"`
	State   string   `json:"state"`
	Status  string   `json:"status"`
	Created int64    `json:"created"`
}

type ContainerStats struct {
	CPUPercent    float64 `json:"cpuPercent"`
	MemoryUsage   float64 `json:"memoryUsage"`
	MemoryLimit   float64 `json:"memoryLimit"`
	MemoryPercent float64 `json:"memoryPercent"`
}

// NewClient creates a new docker client
func NewClient() (*client.Client, error) {
	return client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
}

// ListContainers returns a list of containers
func ListContainers() ([]ContainerInfo, error) {
	cli, err := NewClient()
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		return nil, err
	}

	var result []ContainerInfo
	for _, c := range containers {
		result = append(result, ContainerInfo{
			ID:      c.ID,
			Names:   c.Names,
			Image:   c.Image,
			State:   c.State,
			Status:  c.Status,
			Created: c.Created,
		})
	}
	return result, nil
}

// StartContainer starts a container
func StartContainer(containerID string) error {
	cli, err := NewClient()
	if err != nil {
		return err
	}
	defer cli.Close()

	return cli.ContainerStart(context.Background(), containerID, types.ContainerStartOptions{})
}

// StopContainer stops a container
func StopContainer(containerID string) error {
	cli, err := NewClient()
	if err != nil {
		return err
	}
	defer cli.Close()

	// Default timeout (nil)
	return cli.ContainerStop(context.Background(), containerID, nil)
}

// RestartContainer restarts a container
func RestartContainer(containerID string) error {
	cli, err := NewClient()
	if err != nil {
		return err
	}
	defer cli.Close()

	return cli.ContainerRestart(context.Background(), containerID, nil)
}

// GetContainerLogs returns the logs of a container
func GetContainerLogs(containerID string) (string, error) {
	cli, err := NewClient()
	if err != nil {
		return "", err
	}
	defer cli.Close()

	out, err := cli.ContainerLogs(context.Background(), containerID, types.ContainerLogsOptions{ShowStdout: true, ShowStderr: true, Tail: "100"})
	if err != nil {
		return "", err
	}
	defer out.Close()

	buf := new(strings.Builder)
	_, err = io.Copy(buf, out)
	return buf.String(), err
}
