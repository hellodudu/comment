package datastore

import (
	"context"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hellodudu/Ultimate/iface"
	"github.com/hellodudu/Ultimate/utils/global"
	"github.com/jinzhu/gorm"
	log "github.com/rs/zerolog/log"
)

type Datastore struct {
	db     *gorm.DB
	ctx    context.Context
	cancel context.CancelFunc
	chStop chan struct{}

	// table
	global *iface.TableGlobal
}

func NewDatastore() (iface.IDatastore, error) {
	datastore := &Datastore{
		chStop: make(chan struct{}, 1),
	}

	datastore.ctx, datastore.cancel = context.WithCancel(context.Background())

	// default use docker env value
	var mysqlAddr string
	if mysqlAddr = os.Getenv("MYSQL_ADDR"); len(mysqlAddr) == 0 {
		mysqlAddr = global.MysqlAddr
	}

	mysqlDSN := fmt.Sprintf("%s:%s@(%s:%s)/%s", global.MysqlUser, global.MysqlPwd, mysqlAddr, global.MysqlPort, global.MysqlDB)
	var err error
	datastore.db, err = gorm.Open("mysql", mysqlDSN)
	if err != nil {
		log.Fatal().Err(err).Send()
		return nil, err
	}

	// datastore.db.LogMode(true)

	datastore.initDatastore()
	return datastore, nil
}

func (m *Datastore) DB() *gorm.DB {
	return m.db
}

func (m *Datastore) TableGlobal() *iface.TableGlobal {
	return m.global
}

func (m *Datastore) Run() {
	for {
		select {
		case <-m.ctx.Done():
			log.Info().Msg("datastore context done...")
			m.chStop <- struct{}{}
			return
		}
	}

}

func (m *Datastore) Stop() {
	m.db.Close()
	m.cancel()
	<-m.chStop
	close(m.chStop)
}

func (m *Datastore) initDatastore() {
	m.loadGlobal()
}

func (m *Datastore) loadGlobal() {
	m.global = &iface.TableGlobal{
		Id:                 global.UltimateID,
		TimeStamp:          int(int32(time.Now().Unix())),
		ArenaSeason:        0,
		ArenaWeekEndTime:   0,
		ArenaSeasonEndTime: 0,
	}

	m.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").AutoMigrate(m.global)
	if m.db.FirstOrCreate(m.global, global.UltimateID).RecordNotFound() {
		m.db.Create(m.global)
	}

	log.Info().
		Interface("global", m.global).
		Msg("datastore loadGlobal success")
}
