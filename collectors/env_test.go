package collectors_test

import (
	"testing"

	"github.com/ajaypanthagani/sugabred/collectors"
	commandmock "github.com/ajaypanthagani/sugabred/commands/mocks"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type MockEnvReader struct {
	Vars []string
}

func (m MockEnvReader) Environ() []string {
	return m.Vars
}

func TestEnv(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Env Collector Suite")
}

var _ = Describe("Env Collector", func() {
	var (
		mockEnvCommander *commandmock.MockEnvCommander
		ctrl             *gomock.Controller
	)

	var (
		envCollector collectors.EnvCollector
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockEnvCommander = commandmock.NewMockEnvCommander(ctrl)
		envCollector = collectors.NewEnvCollector(mockEnvCommander)
	})

	It("should collect environment variables into a map", func() {
		mockEnvValues := []string{"FOO=bar", "HELLO=world"}

		mockEnvCommander.EXPECT().Environ().Return(mockEnvValues)
		env := envCollector.CollectEnvVars()
		Expect(env).To(HaveKeyWithValue("FOO", "bar"))
		Expect(env).To(HaveKeyWithValue("HELLO", "world"))
	})
})
