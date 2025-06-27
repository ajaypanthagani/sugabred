package output

import (
	"os"

	"github.com/ajaypanthagani/sugabred/types"
	"gopkg.in/yaml.v3"
)

type FileWriter interface {
	WriteFile(name string, data []byte, perm os.FileMode) error
}

type OSFileWriter struct{}

func (OSFileWriter) WriteFile(name string, data []byte, perm os.FileMode) error {
	return os.WriteFile(name, data, perm)
}

func WriteSnapshot(snapshot *types.Snapshot, path string, writer FileWriter) error {
	data, err := yaml.Marshal(snapshot)
	if err != nil {
		return err
	}
	return writer.WriteFile(path, data, 0644)
}

func WriteSnapshotToFile(snapshot *types.Snapshot, path string) error {
	return WriteSnapshot(snapshot, path, OSFileWriter{})
}
