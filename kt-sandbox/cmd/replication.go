package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	InstancePrefixName = "kt-"
)

var replicationConfig struct {
	clusterName string
	instances   int
	isSemiSync  bool
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
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("replication called")
	},
}

func init() {
	replicationFlags := replicationCmd.Flags()
	replicationFlags.StringVarP(&replicationConfig.clusterName, "name", "n", "", "replication cluster name")
	replicationFlags.IntVarP(&replicationConfig.instances, "instances", "i", 1, "replication instances. 1 source, N replicas.")
	replicationFlags.BoolVarP(&replicationConfig.isSemiSync, "semiSync", "s", false, "Configure semi-synchronous replication")
	rootCmd.AddCommand(replicationCmd)
}
