package store

import (
	"context"
	"errors"
	"sync"
)

const ModuleName = "store"

var ErrRecordNotFound = errors.New("record not found")

type Store struct {
	memory
}

func New() *Store {
	return &Store{}
}

func (s *Store) Name() string {
	return ModuleName
}

func (s *Store) Start(_ context.Context) error {
	if s == nil {
		return nil
	}

	s.memory = memory{sync.Map{}}
	return nil
}

func (s *Store) Stop() error {
	if s == nil {
		return nil
	}

	s.memory.Clear()
	return nil
}
