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
loose_repl_semi_sync_master_timeout=10000
`
	SemiSyncReplicationReplicaTemplate = `
loose_rpl_semi_sync_slave_enabled=1
`
)

const (
	SourceDirName        = "source"
	ReplicaDirNameFormat = "replica-%d"
	MyCnfName            = "my.cnf"
)

type MyCnf struct {
	path       string
	index      int // if index is 0, it is source instance
	isSemiSync bool
}

func NewMyCnf(basePath string, index int, isSemiSync bool) *MyCnf {
	return &MyCnf{
		path:       basePath,
		index:      index,
		isSemiSync: isSemiSync,
	}
}

func NewAsyncReplMyCnf(basePath string, index int) *MyCnf {
	return &MyCnf{
		path:       basePath,
		index:      index,
		isSemiSync: false,
	}
}

func NewSemiSyncReplMyCnf(basePath string, index int) *MyCnf {
	return &MyCnf{
		path:       basePath,
		index:      index,
		isSemiSync: true,
	}
}

func (m *MyCnf) GenerateReplicationConfig() error {
	if m.index == 0 {
		return m.GenerateSourceConfig()
	}
	return m.GenerateReplicaConfig()
}

func (m *MyCnf) GenerateSourceConfig() error {
	return m.GenerateReplicationConfigWithTmpl(AsyncReplicationMyCnfTemplate, SemiSyncReplicationSourceTemplate)
}

func (m *MyCnf) GenerateReplicaConfig() error {
	return m.GenerateReplicationConfigWithTmpl(AsyncReplicationMyCnfTemplate, SemiSyncReplicationReplicaTemplate)
}

// GenerateReplicationConfigWithTmpl
// asyncTmpl includes server_id=%d in the first line of my.cnf
func (m *MyCnf) GenerateReplicationConfigWithTmpl(asyncTmpl string, semiSyncTmpl string) error {
	myCnfBody := fmt.Sprintf(asyncTmpl, m.index+1)
	if m.isSemiSync {
		myCnfBody += semiSyncTmpl
	}

	dirName := SourceDirName
	if m.index > 0 {
		dirName = fmt.Sprintf(ReplicaDirNameFormat, m.index)
	}

	confPath := filepath.Join(m.path, dirName)
	if _, err := os.Stat(confPath); err != nil {
		err := os.MkdirAll(confPath, 0777)
		if err != nil {
			return err
		}
	}

	err := os.WriteFile(filepath.Join(confPath, MyCnfName), []byte(myCnfBody), 0777)
	if err != nil {
		return err
	}
	return nil
}
