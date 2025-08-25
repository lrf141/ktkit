package pkg

import "fmt"

type Image interface {
	Reference() string
}

type DockerImage struct {
	repository string
	tag        string
}

func NewDockerImage(repository, tag string) *DockerImage {
	return &DockerImage{
		repository: repository,
		tag:        tag,
	}
}

func (i *DockerImage) Reference() string {
	return fmt.Sprintf("%s:%s", i.repository, i.tag)
}
