package pkg

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	AsyncReplicationMyCnfTemplate = `
[mysqld]
server_id=%d
log_bin=ON
gtid_mode=ON
enforce-gtid-consistency=ON
`
	SemiSyncReplicationSourceTemplate = `
loose_repl_semi_sync_master_enable=1
loose_repl_semi_sync_master_timeout=%d
`
	SemiSyncReplicationReplicaTemplate = `
loose_rpl_semi_sync_slave_enabled=1
`
)

const (
	DefaultSemiSyncMasterTimeout = 10000
	SourceDirName                = "source"
	ReplicaDirNameFormat         = "replica-%d"
	MyCnfName                    = "my.cnf"
)

type MyCnf struct {
	path       string
	instance   int
	isSemiSync bool
}

func NewMyCnf(basePath string, instance int, isSemiSync bool) *MyCnf {
	return &MyCnf{
		path:       basePath,
		instance:   instance,
		isSemiSync: isSemiSync,
	}
}

func (m *MyCnf) GenerateReplicationMyCnf() error {
	for i := 0; i < m.instance; i++ {
		myCnfBody := fmt.Sprintf(AsyncReplicationMyCnfTemplate, i+1)
		if i == 0 && m.isSemiSync {
			myCnfBody = myCnfBody + fmt.Sprintf(SemiSyncReplicationSourceTemplate, DefaultSemiSyncMasterTimeout)
		}
		if i != 0 && m.isSemiSync {
			myCnfBody = myCnfBody + fmt.Sprintf(SemiSyncReplicationReplicaTemplate)
		}

		var dstPath string
		if i == 0 {
			dstPath = filepath.Join(m.path, fmt.Sprintf(SourceDirName))
		} else {
			dstPath = filepath.Join(m.path, fmt.Sprintf(ReplicaDirNameFormat, i))
		}

		if _, err := os.Stat(dstPath); err != nil {
			err := os.MkdirAll(dstPath, 0777)
			if err != nil {
				return err
			}
		}

		err := os.WriteFile(filepath.Join(dstPath, MyCnfName), []byte(myCnfBody), 0777)
		if err != nil {
			return err
		}
	}
	return nil
}
