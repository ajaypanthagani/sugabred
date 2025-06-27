package collectors_test

import (
	"encoding/json"
	"testing"

	"github.com/ajaypanthagani/sugabred/collectors"
	commandmock "github.com/ajaypanthagani/sugabred/commands/mocks"
	"github.com/ajaypanthagani/sugabred/types"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestBrew(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Brew Collector Suite")
}

var _ = Describe("Brew Collector", func() {
	var (
		mockBrewCommander *commandmock.MockBrewCommander
		ctrl              *gomock.Controller
	)

	var (
		brewCollector collectors.BrewCollector
	)

	var (
		mockBrewFormulae = []string{"go", "git"}
		mockBrewCasks    = []string{"chrome"}
		mockBrewInfoMap  = map[string][]byte{
			"git":    mockFormulaJSON("git", "2.44.0"),
			"go":     mockFormulaJSON("go", "1.22.1"),
			"chrome": mockCaskJSON("chrome", "125.0"),
		}
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockBrewCommander = commandmock.NewMockBrewCommander(ctrl)
		brewCollector = collectors.NewBrewCollector(mockBrewCommander)
	})

	It("collects brew packages correctly", func() {
		mockBrewCommander.EXPECT().RunBrewListFormula().Return(mockBrewFormulae, nil)
		mockBrewCommander.EXPECT().RunBrewInfoJSON("git", false).Return(mockBrewInfoMap["git"], nil)
		mockBrewCommander.EXPECT().RunBrewInfoJSON("go", false).Return(mockBrewInfoMap["go"], nil)

		pkgs, err := brewCollector.CollectPackages()

		Expect(err).ToNot(HaveOccurred())
		Expect(pkgs).To(ConsistOf(
			types.BrewPackage{Name: "git", Version: "2.44.0"},
			types.BrewPackage{Name: "go", Version: "1.22.1"},
		))
	})

	It("collects brew casks correctly", func() {
		mockBrewCommander.EXPECT().RunBrewListCask().Return(mockBrewCasks, nil)
		mockBrewCommander.EXPECT().RunBrewInfoJSON("chrome", true).Return(mockBrewInfoMap["chrome"], nil)

		casks, err := brewCollector.CollectCasks()
		Expect(err).ToNot(HaveOccurred())
		Expect(casks).To(ConsistOf(
			types.BrewCask{Name: "chrome", Version: "125.0"},
		))
	})
})

func mockFormulaJSON(name, version string) []byte {
	data := types.BrewPackageInfo{
		Formulae: []types.BrewFormula{
			{
				Name: name,
				Versions: types.BrewVersions{
					Stable: version,
				},
			},
		},
	}

	res, _ := json.Marshal(data)
	return res
}

func mockCaskJSON(name, version string) []byte {
	data := types.BrewCaskInfo{
		Casks: []types.BrewCaskFormula{
			{
				Token:   name,
				Version: version,
			},
		},
	}
	res, _ := json.Marshal(data)
	return res
}
