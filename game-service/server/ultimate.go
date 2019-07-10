package server

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/go-redis/redis"
	datastore "github.com/hellodudu/Ultimate/game-service/db"
	"github.com/hellodudu/Ultimate/game-service/game"
	"github.com/hellodudu/Ultimate/game-service/handler"
	"github.com/hellodudu/Ultimate/game-service/world"
	"github.com/hellodudu/Ultimate/iface"
	"github.com/hellodudu/Ultimate/logger"
	pbGame "github.com/hellodudu/Ultimate/proto/game"
	"github.com/hellodudu/Ultimate/utils/global"
	"github.com/hellodudu/Ultimate/utils/task"
	"github.com/micro/go-micro"
)

// ultimate define
type ultimate struct {
	td iface.IDispatcher // task dispatcher
	ds iface.IDatastore  // datastore
	wm iface.IWorldMgr   // world manager
	gm iface.IGameMgr    // game manager
	mp iface.IMsgParser  // msg parser

	gameSrv     micro.Service
	gameHandler *handler.GameHandler

	rds      *redis.Client // redis
	tcpServ  *TCPServer    // tcp server
	httpServ *HttpServer   // http server
	wg       sync.WaitGroup
}

func NewUltimate() (iface.IUltimate, error) {
	umt := &ultimate{}

	if err := umt.InitDatastore(); err != nil {
		return nil, err
	}

	if err := umt.InitTask(); err != nil {
		return nil, err
	}

	if err := umt.InitWorldMgr(); err != nil {
		return nil, err
	}

	if err := umt.InitGameMgr(); err != nil {
		return nil, err
	}

	if err := umt.InitGameService(); err != nil {
		return nil, err
	}

	if err := umt.InitMsgParser(); err != nil {
		return nil, err
	}

	if err := umt.InitTCPServer(); err != nil {
		return nil, err
	}

	if err := umt.InitHttpServer(); err != nil {
		return nil, err
	}

	logger.Info("all init ok!")

	return umt, nil
}

func (umt *ultimate) WorldMgr() iface.IWorldMgr {
	return umt.wm
}

func (umt *ultimate) GameMgr() iface.IGameMgr {
	return umt.gm
}

func (umt *ultimate) Datastore() iface.IDatastore {
	return umt.ds
}

// init task and taskdispatcher
func (umt *ultimate) InitTask() error {
	var err error
	if umt.td, err = task.NewDispatcher(); err != nil {
		return err
	}

	logger.Info("task init ok!")

	return nil
}

// init db
func (umt *ultimate) InitDatastore() error {
	var err error
	if umt.ds, err = datastore.NewDatastore(); err != nil {
		return err
	}

	logger.Info("datastore init ok!")
	return nil
}

func (umt *ultimate) InitRedis() {
	umt.rds = redis.NewClient(&redis.Options{
		Addr:     global.RedisAddr,
		Password: global.RedisPwd,
		DB:       global.RedisDB,
	})

	if _, err := umt.rds.Ping().Result(); err != nil {
		logger.Fatal(err)
		return
	}

	logger.Info("redis init ok")
}

func (umt *ultimate) InitMsgParser() error {
	umt.mp = NewMsgParser(umt.gm, umt.wm)
	logger.Info("msg parser init ok!")
	return nil
}

// InitTCPServer init
func (umt *ultimate) InitTCPServer() error {
	var err error
	if umt.tcpServ, err = NewTcpServer(umt.mp, umt.td); err != nil {
		return err
	}

	logger.Info("tcp_server init ok!")
	return nil
}

// init http server
func (umt *ultimate) InitHttpServer() error {
	umt.httpServ = NewHttpServer(umt.gm)
	logger.Info("http_server init ok!")
	return nil
}

// init world session
func (umt *ultimate) InitWorldMgr() error {
	var err error
	if umt.wm, err = world.NewWorldMgr(umt.ds); err != nil {
		return err
	}

	logger.Info("world_mgr init ok!")
	return nil
}

func (umt *ultimate) InitGameMgr() error {
	var err error
	if umt.gm, err = game.NewGameMgr(umt.wm); err != nil {
		return err
	}

	logger.Info("game_mgr init ok!")
	return nil
}

func (umt *ultimate) InitGameService() error {

	var err error
	if umt.gameHandler, err = handler.NewGameHandler(umt.gm, umt.wm); err != nil {
		return err
	}

	// New Service
	umt.gameSrv = micro.NewService(
		micro.Name("ultimate.service.game"),
		micro.Version("latest"),
	)

	// Initialise service
	umt.gameSrv.Init()

	// Register Handler
	pbGame.RegisterGameServiceHandler(umt.gameSrv.Server(), umt.gameHandler)

	logger.Info("game init ok!")
	return nil
}

// run
func (umt *ultimate) Run() {
	go umt.tcpServ.Run()
	go umt.httpServ.Run()
	go umt.wm.Run()
	go umt.gm.Run()
	go umt.ds.Run()

	// rpc service
	go func() {
		if err := umt.gameSrv.Run(); err != nil {
			logger.Fatal(err)
		}
	}()

	// server exit
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		sig := <-c
		logger.Info(fmt.Sprintf("ultimate server closing down (signal: %v)", sig))

		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT:
			umt.Stop()
			logger.Info("server exit safely")
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}

func (umt *ultimate) Stop() {
	umt.tcpServ.Stop()
	<-umt.ds.Stop()
	<-umt.wm.Stop()
}