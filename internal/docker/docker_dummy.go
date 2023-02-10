package docker

import (
	"simple-start-page/internal/entities"
)

type dockerDummy struct{}

func (d dockerDummy) GetLinks() ([]entities.Link, error) {
	return []entities.Link{}, nil
}

func (d dockerDummy) Close() error {
	return nil
}

func NewDockerDummyClient() Docker {
	return dockerDummy{}
}
