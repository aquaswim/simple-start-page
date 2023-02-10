package internal

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"simple-start-page/internal/entities"
)

type Docker interface {
	GetLinks() ([]entities.Link, error)
	Close() error
}

type docker struct {
	client *client.Client
}

func (d *docker) Close() error {
	return d.client.Close()
}

func (d *docker) GetLinks() ([]entities.Link, error) {
	containers, err := d.client.ContainerList(context.Background(), types.ContainerListOptions{
		Filters: filters.NewArgs(filters.Arg("status", "running")),
	})
	if err != nil {
		return nil, err
	}
	var out []entities.Link
	for _, container := range containers {
		// skip container without label
		if len(container.Labels) < 1 {
			continue
		}
		// only accept container that have ssp.url, ssp.name, ssp.icon label
		url, urlKeyExists := container.Labels["com.ssp.url"]
		icon, iconKeyExists := container.Labels["com.ssp.icon"]
		name, nameKeyExists := container.Labels["com.ssp.url"]
		if !urlKeyExists || !iconKeyExists || !nameKeyExists {
			continue
		}
		out = append(out, entities.Link{
			Name: name,
			Url:  url,
			Icon: icon,
		})
	}
	return out, nil
}

func NewDockerIntegration() (Docker, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	return &docker{
		client: cli,
	}, nil
}
