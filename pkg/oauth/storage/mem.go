package storage

import (
	"log"
	"github.com/RangelReale/osin"
)

type MemStorage struct {
	clients   map[string]osin.Client
	authorize map[string]*osin.AuthorizeData
	access    map[string]*osin.AccessData
	refresh   map[string]string
}

func NewMemStorage() *MemStorage {
	r := &MemStorage{
		clients:   make(map[string]osin.Client),
		authorize: make(map[string]*osin.AuthorizeData),
		access:    make(map[string]*osin.AccessData),
		refresh:   make(map[string]string),
	}

	return r
}

func (s *MemStorage) Clone() osin.Storage {
	return s
}

func (s *MemStorage) Close() {
}

func (s *MemStorage) GetClient(id string) (osin.Client, error) {
	log.Printf("GetClient: %s\n", id)
	if c, ok := s.clients[id]; ok {
		return c, nil
	}
	return nil, osin.ErrNotFound
}

func (s *MemStorage) SetClient(id string, client osin.Client) error {
	log.Printf("SetClient: %s\n", id)
	s.clients[id] = client
	return nil
}

func (s *MemStorage) SaveAuthorize(data *osin.AuthorizeData) error {
	log.Printf("SaveAuthorize: %s\n", data.Code)
	s.authorize[data.Code] = data
	return nil
}

func (s *MemStorage) LoadAuthorize(code string) (*osin.AuthorizeData, error) {
	log.Printf("LoadAuthorize: %s\n", code)
	if d, ok := s.authorize[code]; ok {
		return d, nil
	}
	return nil, osin.ErrNotFound
}

func (s *MemStorage) RemoveAuthorize(code string) error {
	log.Printf("RemoveAuthorize: %s\n", code)
	delete(s.authorize, code)
	return nil
}

func (s *MemStorage) SaveAccess(data *osin.AccessData) error {
	log.Printf("SaveAccess: %s\n", data.AccessToken)
	s.access[data.AccessToken] = data
	if data.RefreshToken != "" {
		log.Printf("SaveRefresh: %s\n", data.RefreshToken)
		s.refresh[data.RefreshToken] = data.AccessToken
	}
	return nil
}

func (s *MemStorage) LoadAccess(code string) (*osin.AccessData, error) {
	log.Printf("LoadAccess: %s\n", code)
	if d, ok := s.access[code]; ok {
		return d, nil
	}
	return nil, osin.ErrNotFound
}

func (s *MemStorage) RemoveAccess(code string) error {
	log.Printf("RemoveAccess: %s\n", code)
	delete(s.access, code)
	return nil
}

func (s *MemStorage) LoadRefresh(code string) (*osin.AccessData, error) {
	log.Printf("LoadRefresh: %s\n", code)
	if d, ok := s.refresh[code]; ok {
		return s.LoadAccess(d)
	}
	return nil, osin.ErrNotFound
}

func (s *MemStorage) RemoveRefresh(code string) error {
	log.Printf("RemoveRefresh: %s\n", code)
	delete(s.refresh, code)
	return nil
}
