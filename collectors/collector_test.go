package collectors_test

import (
	"testing"

	"github.com/ajaypanthagani/sugabred/collectors"
	collectormock "github.com/ajaypanthagani/sugabred/collectors/mocks"
	"github.com/ajaypanthagani/sugabred/types"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCollector(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Collector Suite")
}

var _ = Describe("CollectAll", func() {
	var (
		mockBrewCollector *collectormock.MockBrewCollector
		mockEnvCollector  *collectormock.MockEnvCollector
		ctrl              *gomock.Controller
	)

	var (
		collector collectors.DevEnvCollector
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockBrewCollector = collectormock.NewMockBrewCollector(ctrl)
		mockEnvCollector = collectormock.NewMockEnvCollector(ctrl)
		collector = collectors.NewDevEnvCollector(mockBrewCollector, mockEnvCollector)
	})

	It("should return a valid snapshot", func() {
		brewPkgs := []types.BrewPackage{
			{
				Name:    "go",
				Version: "1.22.1",
			},
		}

		brewCasks := []types.BrewCask{
			{
				Name:    "chrome",
				Version: "125.0",
			},
		}

		envValues := map[string]string{
			"HELLO": "world",
		}

		mockBrewCollector.EXPECT().CollectPackages().Return(brewPkgs, nil)
		mockBrewCollector.EXPECT().CollectCasks().Return(brewCasks, nil)
		mockEnvCollector.EXPECT().CollectEnvVars().Return(envValues)

		snapshot, err := collector.CollectAll()
		Expect(err).To(BeNil())
		Expect(snapshot.Homebrew).To(ConsistOf(types.BrewPackage{Name: "go", Version: "1.22.1"}))
		Expect(snapshot.Casks).To(ConsistOf(types.BrewCask{Name: "chrome", Version: "125.0"}))
		Expect(snapshot.EnvVars).To(HaveKeyWithValue("HELLO", "world"))
	})
})
