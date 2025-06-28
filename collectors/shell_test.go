package collectors_test

import (
	"errors"
	"fmt"

	"github.com/ajaypanthagani/sugabred/collectors"
	commandmock "github.com/ajaypanthagani/sugabred/commands/mocks"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Shell Collector", func() {
	var (
		mockShellCommander *commandmock.MockShellCommander
		mockFileCommander  *commandmock.MockFileCommander
		ctrl               *gomock.Controller
	)

	var (
		shellCollector collectors.ShellCollector
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockShellCommander = commandmock.NewMockShellCommander(ctrl)
		mockFileCommander = commandmock.NewMockFileCommander(ctrl)
		shellCollector = collectors.NewShellCollector(mockShellCommander, mockFileCommander)
	})

	Describe("CollectShell", func() {
		When("error retrieving default shell", func() {
			It("returns error", func() {
				expectedErr := fmt.Errorf("error retrieving default shell")
				mockShellCommander.EXPECT().RunCommand("dscl", gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return("", expectedErr)

				shellConfig, err := shellCollector.CollectShell()

				Expect(err).To(HaveOccurred())
				Expect(errors.Is(err, expectedErr)).To(BeTrue())
				Expect(shellConfig).To(BeNil())
			})
		})

		When("successfully retrieved shell config", func() {
			It("returns shell config", func() {
				mockShellCommander.EXPECT().RunCommand("dscl", gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return("", nil)
				mockShellCommander.EXPECT().RunCommand(gomock.Any(), "-i", "-c", "alias").Return("", nil)
				mockFileCommander.EXPECT().ReadFile(gomock.Any()).Return("", nil).Times(6)

				shellConfig, err := shellCollector.CollectShell()

				Expect(err).ToNot(HaveOccurred())
				Expect(shellConfig).ToNot(BeNil())
			})
		})
	})
})
