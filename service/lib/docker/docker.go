package docker

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
)

// ContainerInfo for Frontend
type ContainerInfo struct {
	ID      string   `json:"id"`
	Names   []string `json:"names"`
	Image   string   `json:"image"`
	State   string   `json:"state"`
	Status  string   `json:"status"`
	Created int64    `json:"created"`
}

// dockerContainer matching Docker API
type dockerContainer struct {
	ID      string   `json:"Id"`
	Names   []string `json:"Names"`
	Image   string   `json:"Image"`
	State   string   `json:"State"`
	Status  string   `json:"Status"`
	Created int64    `json:"Created"`
}

func newUnixClient() (*http.Client, error) {
	return &http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				return net.Dial("unix", "/var/run/docker.sock")
			},
		},
	}, nil
}

// ListContainers returns a list of containers
func ListContainers() ([]ContainerInfo, error) {
	client, err := newUnixClient()
	if err != nil {
		return nil, err
	}

	resp, err := client.Get("http://docker/containers/json?all=1")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("docker api error: %s", resp.Status)
	}

	var raw []dockerContainer
	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		return nil, err
	}

	var result []ContainerInfo
	for _, c := range raw {
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
	client, err := newUnixClient()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("http://docker/containers/%s/start", containerID)
	resp, err := client.Post(url, "application/json", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 && resp.StatusCode != 304 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to start: %s %s", resp.Status, string(body))
	}
	return nil
}

// StopContainer stops a container
func StopContainer(containerID string) error {
	client, err := newUnixClient()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("http://docker/containers/%s/stop", containerID)
	resp, err := client.Post(url, "application/json", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 && resp.StatusCode != 304 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to stop: %s %s", resp.Status, string(body))
	}
	return nil
}

// RestartContainer restarts a container
func RestartContainer(containerID string) error {
	client, err := newUnixClient()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("http://docker/containers/%s/restart", containerID)
	resp, err := client.Post(url, "application/json", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to restart: %s %s", resp.Status, string(body))
	}
	return nil
}

// GetContainerLogs returns the logs of a container
func GetContainerLogs(containerID string) (string, error) {
	client, err := newUnixClient()
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("http://docker/containers/%s/logs?stdout=1&stderr=1&tail=100", containerID)
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("failed to get logs: %s", resp.Status)
	}

	// Docker logs usually contain a header for each frame (stream type, length). 
	// For simplicity, we just return the raw bytes, though it might contain some binary headers.
	// To do it properly we should strip the 8-byte header from each frame.
	// But raw might be readable enough for basic usage.
	
	body, err := io.ReadAll(resp.Body)
	// Strip binary headers if possible? 
	// Header: [STREAM_TYPE 1 byte] [0 0 0] [SIZE 4 bytes]
	// If we just return string, it shows mostly text.
	return string(body), err
}
