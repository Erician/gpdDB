package relog

import (
	"os"
	"testing"

	"github.com/erician/gpdDB/common/gpdconst"
)

func TestNewRecoveryLog(t *testing.T) {
	dbName := "aaa"
	reLog, err := NewRecoveryLog(dbName)
	if err != nil {
		t.Error("expected ", nil, "not ", err)
	}
	reLog.Close()
	os.Remove(dbName + RecoveryLogDefaultSuffix)
}

func TestDisplay(t *testing.T) {

	dbName := "aaa"
	key := "bbb"
	value := "ccc"
	reLog, _ := NewRecoveryLog(dbName)
	reLog.WriteLogRecord(NewLogRecordCheckpoint())
	reLog.WriteLogRecord(NewLogRecordUserOp(gpdconst.PUT, 0, key, value))
	reLog.WriteLogRecord(NewLogRecordUserOp(gpdconst.DELETE, 1, key, value))
	reLog.WriteLogRecord(NewLogRecordAllocate(3))
	reLog.WriteLogRecord(NewLogRecordSetField(4, 40, value))

	reLog.Display()
	reLog.Close()

	os.Remove(dbName + RecoveryLogDefaultSuffix)

}
