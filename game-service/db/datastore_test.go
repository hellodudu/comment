package datastore

import (
	"testing"

	"github.com/hellodudu/Ultimate/utils/global"
)

func init() {
	global.Debugging = false
	global.MysqlUser = "root"
	global.MysqlPwd = ""
	global.MysqlAddr = "127.0.0.1"
	global.MysqlPort = "3306"
	global.MysqlDB = "db_ultimate"
	global.UltimateID = 110

}

func TestNewDatastore(t *testing.T) {

	_, err := NewDatastore()
	if err != nil {
		t.Error("NewDatastore error:", err)
	}

}
