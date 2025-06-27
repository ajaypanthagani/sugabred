package collectors_test

import (
	"encoding/json"
	"fmt"
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

	Describe("CollectPackages", func() {
		When("error retrieving formula list", func() {
			It("returns error", func() {
				expectedErr := fmt.Errorf("error retrieving brew formula list")
				mockBrewCommander.EXPECT().RunBrewListFormula().Return(nil, expectedErr)

				pkgs, err := brewCollector.CollectPackages()

				Expect(err).To(HaveOccurred())
				Expect(err).To(Equal(expectedErr))
				Expect(pkgs).To(BeNil())
			})
		})

		When("error retrieving brew info for a package", func() {
			It("returns remaining packages skipping erroneous package", func() {
				expectedErr := fmt.Errorf("error retrieving brew info for package")

				mockBrewCommander.EXPECT().RunBrewListFormula().Return(mockBrewFormulae, nil)
				mockBrewCommander.EXPECT().RunBrewInfoJSON("git", false).Return(nil, expectedErr)
				mockBrewCommander.EXPECT().RunBrewInfoJSON("go", false).Return(mockBrewInfoMap["go"], nil)

				pkgs, err := brewCollector.CollectPackages()

				Expect(err).ToNot(HaveOccurred())
				Expect(pkgs).To(ConsistOf(
					types.BrewPackage{Name: "go", Version: "1.22.1"},
				))
				Expect(pkgs).ToNot(ConsistOf(
					types.BrewPackage{Name: "git", Version: "2.44.0"},
				))
			})
		})

		When("successfully collects brew packages", func() {
			It("returns collected packages", func() {
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
		})
	})

	Describe("CollectCasks", func() {
		When("error retrieving cask list", func() {
			It("returns error", func() {
				expectedErr := fmt.Errorf("error retrieving brew formula list")
				mockBrewCommander.EXPECT().RunBrewListCask().Return(nil, expectedErr)

				casks, err := brewCollector.CollectCasks()

				Expect(err).To(HaveOccurred())
				Expect(err).To(Equal(expectedErr))
				Expect(casks).To(BeNil())
			})
		})

		When("error retrieving brew info for a cask", func() {
			It("skipps erroneous cask in result", func() {
				expectedErr := fmt.Errorf("error retrieving brew formula list")

				mockBrewCommander.EXPECT().RunBrewListCask().Return(mockBrewCasks, nil)
				mockBrewCommander.EXPECT().RunBrewInfoJSON("chrome", true).Return(nil, expectedErr)

				casks, err := brewCollector.CollectCasks()

				Expect(err).ToNot(HaveOccurred())
				Expect(casks).ToNot(ConsistOf(
					types.BrewCask{Name: "chrome", Version: "125.0"},
				))
			})
		})

		When("successfully collects brew packages", func() {
			It("returns collected packages", func() {
				mockBrewCommander.EXPECT().RunBrewListCask().Return(mockBrewCasks, nil)
				mockBrewCommander.EXPECT().RunBrewInfoJSON("chrome", true).Return(mockBrewInfoMap["chrome"], nil)

				casks, err := brewCollector.CollectCasks()
				Expect(err).ToNot(HaveOccurred())
				Expect(casks).To(ConsistOf(
					types.BrewCask{Name: "chrome", Version: "125.0"},
				))
			})
		})
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
