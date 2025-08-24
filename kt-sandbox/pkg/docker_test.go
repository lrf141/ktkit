package pkg_test

import (
	"context"
	"fmt"
	"github.com/lrf141/ktkit/kt-sandbox/pkg"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
	"os/exec"
)

const (
	testContainerTag = "8.4.6-test"
)

var _ = Describe("Docker suite", func() {

	var sut *pkg.DockerManager

	var ctx context.Context

	var containerId string
	var networkId string

	It("Init Docker suite", func() {
		ctx = context.Background()
		err := exec.Command("which", "docker").Run()
		if err != nil {
			fmt.Printf("Failed to setup pkg test, need docker: %v", err)
			os.Exit(1)
		}
	})

	BeforeEach(func() {
		manager, err := pkg.NewDockerManager(pkg.DefaultImageRepository, testContainerTag)
		if err != nil {
			fmt.Printf("Failed to setup pkg test: %v", err)
			os.Exit(1)
		}
		sut = manager
	})

	It("ExistImage", func() {
		// Setup
		_ = exec.Command("docker", "pull", pkg.DefaultImageRepository+":"+testContainerTag).Run()

		// Exercise
		exist, err := sut.ExistImage(ctx)

		// Verify
		Expect(err).NotTo(HaveOccurred())
		Expect(exist).To(BeTrue())
	})

	It("PullImage", func() {
		// Exercise
		err := sut.PullImage(ctx)

		// Verify
		Expect(err).NotTo(HaveOccurred())
		exist, err := sut.ExistImage(ctx)
		Expect(err).NotTo(HaveOccurred())
		Expect(exist).To(BeTrue())
	})

	It("CreateNetwork", func() {
		// Exercise
		id, err := sut.CreateNetwork(ctx)

		// Verify
		Expect(err).NotTo(HaveOccurred())
		Expect(id).NotTo(BeNil())

		networkId = id
	})

	It("CreateContainer", func() {
		// Exercise
		id, err := sut.CreateContainer(ctx, "test", 0, networkId)

		// Verify
		Expect(err).NotTo(HaveOccurred())
		Expect(id).NotTo(BeEmpty())

		containerId = id
	})

	It("StartContainer", func() {
		// Exercise
		err := sut.StartContainer(ctx, containerId)
		Expect(err).NotTo(HaveOccurred())
	})

	It("Cleanup docker resource", func() {
		// Exercise
		err := sut.StopContainer(ctx, containerId)
		Expect(err).NotTo(HaveOccurred())
		err = sut.RemoveContainer(ctx, containerId)
		Expect(err).NotTo(HaveOccurred())
		err = sut.RemoveNetwork(ctx, networkId)
		Expect(err).NotTo(HaveOccurred())
	})
})
