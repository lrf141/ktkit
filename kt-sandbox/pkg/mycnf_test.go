package pkg_test

import (
	"github.com/lrf141/ktkit/kt-sandbox/pkg"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
	"path/filepath"
)

const (
	MyCnfBasePathForTest = "./resource/test"
)

var _ = Describe("MyCnf", func() {

	It("Generate MyCnf: Async Replication", func() {
		// Exercise
		sut := pkg.NewMyCnf(MyCnfBasePathForTest, 2, false)
		err := sut.GenerateReplicationMyCnf()
		Expect(err).NotTo(HaveOccurred())

		// Verify
		sourcePath := filepath.Join(MyCnfBasePathForTest, "source", "my.cnf")
		sourceByte, err := os.ReadFile(sourcePath)
		source := string(sourceByte)
		Expect(err).NotTo(HaveOccurred())
		Expect(source).To(ContainSubstring("server_id=1"))
		replicaPath := filepath.Join(MyCnfBasePathForTest, "replica-1", "my.cnf")
		replicaByte, err := os.ReadFile(replicaPath)
		replica := string(replicaByte)
		Expect(err).NotTo(HaveOccurred())
		Expect(replica).To(ContainSubstring("server_id=2"))

		// Cleanup
		err = os.Remove(sourcePath)
		Expect(err).NotTo(HaveOccurred())
		err = os.Remove(replicaPath)
		Expect(err).NotTo(HaveOccurred())
	})
})
