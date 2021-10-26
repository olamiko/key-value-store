package utils

import (
	"bytes"
	"encoding/gob"
	"errors"
	"os"
)

type Store struct {
	data     map[string]string
	Filename string
}

func NewStore(filename string) Store {
	newStore := Store{Filename: filename}
	newStore.Load()
	return newStore
	//return &Store{data: make(map[string]string), Filename: filename}
}

func (s Store) Get(key string) string {
	res, ok := s.data[key]
	if !ok {
		return "Error: Key not found"
	}
	return res
}

func (s *Store) Set(key string, value string) error {
	s.data[key] = value

	err := s.Flush()
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) Load() error {

	// Deal with this later... open file properly
	if _, err := os.Stat(s.Filename); errors.Is(err, os.ErrNotExist) {
		s.data = make(map[string]string)
		return nil
	}

	rawData, err := os.ReadFile(s.Filename)
	if err != nil {
		return err
	}

	buffer := bytes.NewBuffer(rawData)

	decoder := gob.NewDecoder(buffer)
	err = decoder.Decode(&s.data)
	if err != nil {
		return err
	}

	return nil
}

func (s Store) Flush() error {
	buffer := new(bytes.Buffer)

	encoder := gob.NewEncoder(buffer)

	err := encoder.Encode(s.data)
	if err != nil {
		return err
	}

	f, err := os.Create(s.Filename)
	if err != nil {
		return err
	}

	buffer.WriteTo(f)
	return nil
}
