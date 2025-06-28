package commands

//go:generate mockgen -source=file.go -destination=mocks/file.go

import "os"

type FileCommander interface {
	ReadFile(path string) (string, error)
}

func NewFileCommander() FileCommander {
	return &fileCommander{}
}

type fileCommander struct{}

func (*fileCommander) ReadFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	return string(data), err
}
