package output_test

import (
	"os"
	"testing"

	"github.com/ajaypanthagani/sugabred/output"
	"github.com/ajaypanthagani/sugabred/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gopkg.in/yaml.v3"
)

func TestOutput(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Output Suite")
}

type MockFileWriter struct {
	CalledPath string
	Written    []byte
	Perm       os.FileMode
	Err        error
}

func (m *MockFileWriter) WriteFile(name string, data []byte, perm os.FileMode) error {
	m.CalledPath = name
	m.Written = data
	m.Perm = perm
	return m.Err
}

var _ = Describe("WriteSnapshot", func() {
	var (
		mockWriter *MockFileWriter
		snapshot   *types.Snapshot
	)

	BeforeEach(func() {
		mockWriter = &MockFileWriter{}
		snapshot = &types.Snapshot{
			Homebrew: []types.BrewPackage{{Name: "go", Version: "1.22.1"}},
			Casks:    []types.BrewCask{{Name: "chrome", Version: "125.0"}},
			EnvVars:  map[string]string{"FOO": "bar"},
		}
	})

	It("marshals and writes the snapshot using the file writer", func() {
		err := output.WriteSnapshot(snapshot, "mock-path.yaml", mockWriter)
		Expect(err).To(BeNil())

		Expect(mockWriter.CalledPath).To(Equal("mock-path.yaml"))

		var decoded types.Snapshot
		err = yaml.Unmarshal(mockWriter.Written, &decoded)
		Expect(err).To(BeNil())
		Expect(decoded).To(Equal(*snapshot))
	})
})
