package repositories

import (
	"simple-start-page/internal/entities"
	"simple-start-page/pkg"
)

const (
	keySetting = "setting"
	keyLink    = "links"
)

type setting struct {
	db pkg.Database
}

func (s *setting) SaveListUrls(links *[]entities.Link) error {
	return s.db.SaveObject(keyLink, links)
}

func (s *setting) SaveSetting(st *entities.Setting) error {
	return s.db.SaveObject(keySetting, st)
}

func (s *setting) GetListUrls() (*[]entities.Link, error) {
	links := new([]entities.Link)
	err := s.db.GetObject(keyLink, links)
	if err != nil {
		return nil, err
	}
	return links, nil
}

func (s *setting) GetSetting() (*entities.Setting, error) {
	st := new(entities.Setting)
	err := s.db.GetObject(keySetting, st)
	if err != nil {
		return nil, err
	}
	return st, nil
}

type Setting interface {
	GetListUrls() (*[]entities.Link, error)
	GetSetting() (*entities.Setting, error)
	SaveListUrls(links *[]entities.Link) error
	SaveSetting(setting *entities.Setting) error
}

func NewSettingRepo(db pkg.Database) Setting {
	return &setting{
		db: db,
	}
}
