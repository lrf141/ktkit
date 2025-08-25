package pkg

import (
	"bytes"
	"context"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/docker/go-connections/nat"
	"io"
	"os"
	"strings"
	"time"
)

const (
	DefaultImageRepository      = "mysql"
	DefaultImageTag             = "8.4.6"
	InternalNetworkName         = "kt-sandbox-internal"
	InternalNetworkCidr         = "172.30.0.0/24"
	InternalNetworkIpAddrFormat = "172.30.0.%d"
)

type ImageConfig struct {
	repository string
	tag        string
}

func NewImageConfig(repository string, tag string) *ImageConfig {
	return &ImageConfig{
		repository: repository,
		tag:        tag,
	}
}

func (c *ImageConfig) ImageUri() string {
	return c.repository + ":" + c.tag
}

type DockerManager struct {
	client      *client.Client
	imageConfig *ImageConfig
}

func NewDockerManager(imageConfig *ImageConfig) (*DockerManager, error) {
	apiClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}

	return &DockerManager{
		client:      apiClient,
		imageConfig: imageConfig,
	}, nil
}

func (d *DockerManager) ExistImage(ctx context.Context) (bool, error) {
	images, err := d.client.ImageList(ctx, image.ListOptions{})
	if err != nil {
		return false, err
	}
	for _, i := range images {
		for _, imageTag := range i.RepoTags {
			if imageTag == d.imageConfig.ImageUri() {
				return true, nil
			}
		}
	}
	return false, nil
}

func (d *DockerManager) PullImage(ctx context.Context) error {
	existImage, err := d.ExistImage(ctx)
	if err != nil {
		return err
	}
	if existImage {
		fmt.Printf("Already pulled image %s\n", d.imageConfig.ImageUri())
		return nil
	}

	reader, err := d.client.ImagePull(ctx, d.imageConfig.ImageUri(), image.PullOptions{})
	if err != nil {
		return err
	}
	defer reader.Close()
	io.Copy(os.Stdout, reader)
	return nil
}

func (d *DockerManager) CreateContainer(ctx context.Context, name string, index int, networkId string) (string, error) {
	config := &container.Config{
		Image: d.imageConfig.ImageUri(),
		Tty:   true,
		Env: []string{
			"MYSQL_ROOT_PASSWORD=root",
			"MYSQL_DATABASE=test",
		},
	}
	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			"3306/tcp": []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: fmt.Sprintf("3306%d", index),
				},
			},
		},
	}
	networkConfig := &network.NetworkingConfig{
		EndpointsConfig: map[string]*network.EndpointSettings{
			InternalNetworkName: {
				NetworkID: networkId,
				// the value of `index` starts at 0.
				IPAddress: fmt.Sprintf(InternalNetworkIpAddrFormat, index+1),
			},
		},
	}

	c, err := d.client.ContainerCreate(ctx, config, hostConfig, networkConfig, nil, name)
	if err != nil {
		return "", err
	}
	return c.ID, nil
}

func (d *DockerManager) StartContainer(ctx context.Context, id string) error {
	err := d.client.ContainerStart(ctx, id, container.StartOptions{})
	if err != nil {
		return err
	}
	return nil
}

func (d *DockerManager) CreateNetwork(ctx context.Context) (string, error) {
	networkConfig := network.CreateOptions{
		Driver: "bridge",
		Options: map[string]string{
			"subnet": InternalNetworkCidr,
		},
	}
	n, err := d.client.NetworkCreate(ctx, InternalNetworkName, networkConfig)
	if err != nil {
		return "", err
	}
	return n.ID, nil
}

func (d *DockerManager) ExecAttachContainer(ctx context.Context, id string, cmd []string) (string, error) {
	resp, err := d.client.ContainerExecCreate(ctx, id, container.ExecOptions{
		Cmd:          cmd,
		AttachStdout: true,
		AttachStderr: true,
	})
	if err != nil {
		return "", err
	}

	attachResp, err := d.client.ContainerExecAttach(ctx, resp.ID, container.ExecAttachOptions{})
	if err != nil {
		return "", err
	}
	defer attachResp.Close()

	var buf bytes.Buffer
	// ignore stderr
	_, err = stdcopy.StdCopy(&buf, io.Discard, attachResp.Reader)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(buf.String()), nil
}

func (d *DockerManager) WaitContainerUp(ctx context.Context, id string, expected string, cmd []string) error {
	maxWait := 10 * time.Second
	interval := 1 * time.Second
	deadline := time.Now().Add(maxWait)
	for {
		result, err := d.ExecAttachContainer(ctx, id, cmd)
		if err != nil {
			return err
		}

		if result == expected {
			break
		}

		if time.Now().After(deadline) {
			return fmt.Errorf("timed out waiting for container to start %s", expected)
		}
		time.Sleep(interval)
	}
	return nil
}

func (d *DockerManager) StopContainer(ctx context.Context, id string) error {
	return d.client.ContainerStop(ctx, id, container.StopOptions{})
}

func (d *DockerManager) RemoveContainer(ctx context.Context, id string) error {
	return d.client.ContainerRemove(ctx, id, container.RemoveOptions{
		Force: true,
	})
}

func (d *DockerManager) RemoveNetwork(ctx context.Context, id string) error {
	return d.client.NetworkRemove(ctx, id)
}
