package config

import (
	"github.com/hashicorp/hcl"
)

func FromString(filename string, contents string) (*Config, error) {
	config := Config{}
	err := hcl.Decode(&config, contents)
	if err != nil {
		return nil, err
	}

	config.SourceFilename = filename
	return &config, nil
}
