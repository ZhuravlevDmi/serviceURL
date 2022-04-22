package storage

import (
	"errors"
	"github.com/ZhuravlevDmi/serviceURL/internal/util"
)

type StorageMapURL struct {
	MapURL map[string]string
}

type Storage interface {
	Read(miniURL string) string
	Record(bigURL string) (string, error)
}

func ReadStorage(s Storage, miniURL string) string {
	return s.Read(miniURL)
}

func RecordStorage(s Storage, bigURL string) (string, error) {
	return s.Record(bigURL)
}

func (s *StorageMapURL) Read(miniURL string) string {
	return s.MapURL[miniURL]
}

func (s *StorageMapURL) Record(bigURL string) (string, error) {
	for miniURL, URL := range s.MapURL {
		if URL == bigURL {
			return miniURL, errors.New("данный URL уже записан")
		}
	}
	for {
		miniURL := util.GenerateMiniUrl()
		if s.MapURL[miniURL] != "" {
			continue
		}
		s.MapURL[miniURL] = bigURL
		return miniURL, nil
	}
}
