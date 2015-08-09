package config

import (
	"io"
	"io/ioutil"
	"os"
)

func FromFile(filename string) (*Config, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	return FromReader(filename, f)
}

func FromReader(filename string, reader io.Reader) (*Config, error) {
	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return FromString(filename, string(bytes))
}
