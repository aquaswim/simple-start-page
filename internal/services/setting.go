package services

import (
	"simple-start-page/internal/entities"
	"simple-start-page/internal/repositories"
)

type setting struct {
	settingRepo repositories.Setting
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

func (s *setting) ListLink() (*[]entities.Link, error) {
	links, err := s.settingRepo.GetListUrls()
	if err != nil {
		return nil, err
	}
	return links, nil
}

type Setting interface {
	ListLink() (*[]entities.Link, error)
	GetSetting() (*entities.Setting, error)
	UpdateListLink(*[]entities.Link) error
	UpdateSetting(setting2 *entities.Setting) error
}

func NewSettingService(settingRepo repositories.Setting) Setting {
	return &setting{
		settingRepo: settingRepo,
	}
}
