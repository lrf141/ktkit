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

const replicaCount = 1

var _ = Describe("MySQL", func() {
	var sut *pkg.MySQLManager
	var ctx context.Context

	It("Init MySQL suite", func() {
		ctx = context.Background()
		err := exec.Command("which", "docker").Run()
		if err != nil {
			fmt.Printf("Failed to setup mysql test, need docker: %v", err)
			os.Exit(1)
		}
		manager, err := pkg.NewMySQLManager(pkg.DefaultImageRepository, testContainerTag, replicaCount)
		if err != nil {
			fmt.Printf("Failed to setup mysql test: %v", err)
			os.Exit(1)
		}
		sut = manager
	})

	It("CreateInstances", func() {
		// Exercise
		err := sut.CreateInstances(ctx, "test")

		// Verify
		Expect(err).NotTo(HaveOccurred())
		Expect(sut.Source().Name()).To(Equal("kt-test-0"))
		Expect(sut.Source().ContainerId()).NotTo(BeEmpty())
		for index, replica := range sut.Replicas() {
			Expect(replica.Name()).To(Equal(fmt.Sprintf("kt-test-%d", index+1)))
			Expect(replica.ContainerId()).NotTo(BeEmpty())
		}
	})

	It("Cleanup", func() {
		// Exercise
		err := sut.CleanupAll(ctx)
		Expect(err).NotTo(HaveOccurred())
	})
})
