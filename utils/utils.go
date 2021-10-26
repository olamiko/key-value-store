package utils

import (
	"bytes"
	"encoding/gob"
	"log"
	"os"
)

type Store struct {
	data     map[string]string
	filename string
}

func (s Store) Get(key string) string {
	res, ok := s.data[key]
	if !ok {
		return "Error: Key not found"
	}
	return res
}

func (s *Store) Set(key string, value string) {
	s.data[key] = value

	s.flush()
}

func (s *Store) Load() {

	rawData, err := os.ReadFile(s.filename)
	if err != nil {
		log.Fatal(err)
	}

	buffer := bytes.NewBuffer(rawData)

	decoder := gob.NewDecoder(buffer)
	err = decoder.Decode(&s.data)
	if err != nil {
		log.Fatal(err)
	}

}

func (s Store) Flush() {
	buffer := new(bytes.Buffer)

	encoder := gob.NewEncoder(buffer)

	err := encoder.Encode(s.data)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create(s.filename)
	if err != nil {
		log.Fatal(err)
	}

	buffer.WriteTo(f)

}
