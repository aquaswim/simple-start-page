package services

import (
	"log"
	"simple-start-page/internal/docker"
	"simple-start-page/internal/entities"
	"simple-start-page/internal/repositories"
)

type setting struct {
	settingRepo  repositories.Setting
	dockerClient docker.Docker
}

func (s *setting) UpdateListLink(urls *[]entities.Link) error {
	return s.settingRepo.SaveListUrls(urls)
}

func (s *setting) UpdateSetting(st *entities.Setting) error {
	return s.settingRepo.SaveSetting(st)
}

func (s *setting) GetSetting() (*entities.Setting, error) {
	st, err := s.settingRepo.GetSetting()
	if err != nil {
		return nil, err
	}
	return st, nil
}

func (s *setting) ListLink(isSkipGetDockerLinks bool) (*[]entities.Link, error) {
	links, err := s.settingRepo.GetListUrls()
	if err != nil {
		return nil, err
	}
	if isSkipGetDockerLinks {
		return links, nil
	}
	dockerLinks, err := s.dockerClient.GetLinks()
	if err != nil {
		log.Println("Error on docker.GetLinks", err)
	} else if dockerLinks != nil {
		dockerLinks = append(dockerLinks, *links...)
		return &dockerLinks, nil
	}
	return links, nil
}

type Setting interface {
	ListLink(isSkipGetDockerLinks bool) (*[]entities.Link, error)
	GetSetting() (*entities.Setting, error)
	UpdateListLink(*[]entities.Link) error
	UpdateSetting(setting2 *entities.Setting) error
}

func NewSettingService(settingRepo repositories.Setting, dockerClient docker.Docker) Setting {
	return &setting{
		settingRepo:  settingRepo,
		dockerClient: dockerClient,
	}
}
