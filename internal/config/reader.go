package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Reader struct {
	path string
}

func NewReader(path string) *Reader {
	return &Reader{
		path: path,
	}
}

func (r *Reader) ReadConfig() ([]*Model, error) {
	var ret []*Model
	f, err := os.Open(r.path)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file with models: %w", err)
	}

	err = json.NewDecoder(f).Decode(&ret)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file with models: %w", err)
	}

	return ret, nil
}
