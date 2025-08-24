package pkg

import (
	"context"
	"fmt"
	"time"
)

const (
	InstancePrefixName = "kt"
	MySQLActiveStatus  = "mysqld is alive"
)

type MySQLManager struct {
	docker    *DockerManager
	source    *Instance
	replicas  []*Instance
	count     int
	networkId string
}

type Instance struct {
	name        string
	containerId string
}

func (i *Instance) Name() string {
	return i.name
}

func (i *Instance) ContainerId() string {
	return i.containerId
}

func NewMySQLManager(repository string, tag string, count int) (*MySQLManager, error) {
	docker, err := NewDockerManager(repository, tag)
	if err != nil {
		return nil, err
	}
	return &MySQLManager{
		docker:   docker,
		source:   nil,
		replicas: []*Instance{},
		count:    count,
	}, nil
}

func (m *MySQLManager) CreateInstance(ctx context.Context, name string, index int) (*Instance, error) {
	instanceName := fmt.Sprintf("%s-%s-%d", InstancePrefixName, name, index)
	containerId, err := m.docker.CreateContainer(ctx, instanceName, index, m.networkId)
	if err != nil {
		return &Instance{}, err
	}
	instance := &Instance{
		name:        instanceName,
		containerId: containerId,
	}
	err = m.docker.StartContainer(ctx, containerId)
	if err != nil {
		return nil, err
	}
	return instance, nil
}

func (m *MySQLManager) WaitStartUp(ctx context.Context, containerId string) error {
	maxWait := 10 * time.Second
	interval := 1 * time.Second
	deadline := time.Now().Add(maxWait)
	for {
		result, err := m.docker.ExecAttachContainer(ctx, containerId, []string{"/bin/sh", "-c", "mysqladmin ping -uroot -proot"})
		if err != nil {
			return err
		}

		if result == MySQLActiveStatus {
			break
		}

		if time.Now().After(deadline) {
			return fmt.Errorf("timed out waiting for container to start")
		}
		time.Sleep(interval)
	}
	return nil
}

func (m *MySQLManager) CreateInstances(ctx context.Context, name string) error {
	err := m.docker.PullImage(ctx)
	if err != nil {
		return err
	}

	networkId, err := m.docker.CreateNetwork(ctx)
	if err != nil {
		return err
	}
	m.networkId = networkId

	// create source instance
	source, err := m.CreateInstance(ctx, name, 0)
	if err != nil {
		return err
	}
	err = m.WaitStartUp(ctx, source.ContainerId())
	if err != nil {
		return err
	}
	m.source = source

	// create replica instances
	for i := 1; i <= m.count; i++ {
		replica, err := m.CreateInstance(ctx, name, i)
		if err != nil {
			return err
		}
		err = m.WaitStartUp(ctx, replica.ContainerId())
		if err != nil {
			return err
		}
		m.replicas = append(m.replicas, replica)
	}

	return nil
}

func (m *MySQLManager) Source() *Instance {
	return m.source
}

func (m *MySQLManager) Replicas() []*Instance {
	return m.replicas
}

func (m *MySQLManager) CleanupAll(ctx context.Context) error {
	err := m.docker.StopContainer(ctx, m.source.containerId)
	if err != nil {
		return err
	}
	err = m.docker.RemoveContainer(ctx, m.source.containerId)
	if err != nil {
		return err
	}

	for _, replica := range m.replicas {
		err = m.docker.StartContainer(ctx, replica.containerId)
		if err != nil {
			return err
		}

		err = m.docker.RemoveContainer(ctx, replica.containerId)
		if err != nil {
			return err
		}
	}

	err = m.docker.RemoveNetwork(ctx, m.networkId)
	if err != nil {
		return err
	}

	return nil
}
