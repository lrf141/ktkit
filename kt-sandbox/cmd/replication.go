package cmd

import (
	"context"
	"fmt"
	"github.com/lrf141/ktkit/kt-sandbox/pkg"

	"github.com/spf13/cobra"
)

const (
	InstancePrefixName = "kt"
)

var replicationConfig struct {
	clusterName     string
	instances       int
	isSemiSync      bool
	imageRepository string
	imageTag        string
	mycnfPath       string
}

var replicationCmd = &cobra.Command{
	Use:   "replication",
	Short: "Creating a MySQL replication",
	Long:  "Creating a MySQL replication",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if len(replicationConfig.clusterName) == 0 {
			return fmt.Errorf("cluster name must be specified")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return runReplication(cmd, args)
	},
}

func runReplication(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	manager, err := pkg.NewMySQLManager(replicationConfig.imageRepository, replicationConfig.imageTag, replicationConfig.instances)
	if err != nil {
		return err
	}
	err = manager.CreateInstances(ctx, replicationConfig.clusterName)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	replicationFlags := replicationCmd.Flags()
	replicationFlags.StringVarP(&replicationConfig.clusterName, "name", "n", "", "replication cluster name")
	replicationFlags.IntVarP(&replicationConfig.instances, "instances", "i", 1, "replication instances. 1 source, N replicas.")
	replicationFlags.BoolVarP(&replicationConfig.isSemiSync, "semiSync", "s", false, "Configure semi-synchronous replication")
	replicationFlags.StringVarP(&replicationConfig.imageRepository, "image-repository", "r", pkg.DefaultImageRepository, "Image repository")
	replicationFlags.StringVarP(&replicationConfig.imageTag, "image-tag", "t", pkg.DefaultImageTag, "Image tag")
	replicationFlags.StringVarP(&replicationConfig.mycnfPath, "mycnf", "c", "resources/", "Base path of my.cnf.")
	rootCmd.AddCommand(replicationCmd)
}
