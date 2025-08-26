package pkg_test

import (
	"github.com/lrf141/ktkit/kt-sandbox/pkg"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
	"path/filepath"
)

const (
	MyCnfBasePathForTest = "../resource/test"
)

var _ = Describe("MyCnf", func() {

	It("Generate MyCnf: Async Replication Source", func() {
		// Exercise
		sut := pkg.NewAsyncReplMyCnf(MyCnfBasePathForTest, 0)
		err := sut.GenerateReplicationConfig()
		Expect(err).NotTo(HaveOccurred())

		// Verify
		path := filepath.Join(MyCnfBasePathForTest, "source", "my.cnf")
		bodyByte, err := os.ReadFile(path)
		body := string(bodyByte)
		Expect(err).NotTo(HaveOccurred())
		Expect(body).To(ContainSubstring("server_id=1"))
		Expect(body).NotTo(ContainSubstring("loose_repl_semi_sync"))

		// Cleanup
		err = os.Remove(path)
		Expect(err).NotTo(HaveOccurred())
	})

	It("Generate MyCnf: Async Replication Replica", func() {
		// Exercise
		sut := pkg.NewAsyncReplMyCnf(MyCnfBasePathForTest, 1) // if index is greater than 0, instance is replica.
		err := sut.GenerateReplicationConfig()
		Expect(err).NotTo(HaveOccurred())

		// Verify
		path := filepath.Join(MyCnfBasePathForTest, "replica-1", "my.cnf")
		bodyByte, err := os.ReadFile(path)
		body := string(bodyByte)
		Expect(err).NotTo(HaveOccurred())
		Expect(body).To(ContainSubstring("server_id=2"))
		Expect(body).NotTo(ContainSubstring("loose_repl_semi_sync"))

		// Cleanup
		err = os.Remove(path)
		Expect(err).NotTo(HaveOccurred())
	})

	It("Generate MyCnf: Semi Sync Replication Source", func() {
		// Exercise
		sut := pkg.NewSemiSyncReplMyCnf(MyCnfBasePathForTest, 0)
		err := sut.GenerateReplicationConfig()
		Expect(err).NotTo(HaveOccurred())

		// Verify
		path := filepath.Join(MyCnfBasePathForTest, "source", "my.cnf")
		bodyByte, err := os.ReadFile(path)
		body := string(bodyByte)
		Expect(err).NotTo(HaveOccurred())
		Expect(body).To(ContainSubstring("server_id=1"))
		Expect(body).To(ContainSubstring("loose_repl_semi_sync_master_enable=1"))
		Expect(body).NotTo(ContainSubstring("loose_rpl_semi_sync_slave_enabled=1"))

		// Cleanup
		err = os.Remove(path)
		Expect(err).NotTo(HaveOccurred())
	})

	It("Generate MyCnf: Semi Sync Replication Replica", func() {
		// Exercise
		sut := pkg.NewSemiSyncReplMyCnf(MyCnfBasePathForTest, 1)
		err := sut.GenerateReplicationConfig()
		Expect(err).NotTo(HaveOccurred())

		// Verify
		path := filepath.Join(MyCnfBasePathForTest, "replica-1", "my.cnf")
		bodyByte, err := os.ReadFile(path)
		body := string(bodyByte)
		Expect(err).NotTo(HaveOccurred())
		Expect(body).To(ContainSubstring("server_id=2"))
		Expect(body).NotTo(ContainSubstring("loose_repl_semi_sync_master_enable=1"))
		Expect(body).To(ContainSubstring("loose_rpl_semi_sync_slave_enabled=1"))

		// Cleanup
		err = os.Remove(path)
		Expect(err).NotTo(HaveOccurred())
	})
})
