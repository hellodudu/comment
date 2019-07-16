package server

import (
	"fmt"
	"log"
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
	"github.com/nsqio/go-nsq"
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

// NewUltimate return IUltimate
func NewUltimate() (iface.IUltimate, error) {
	umt := &ultimate{}

	if err := umt.initDatastore(); err != nil {
		return nil, err
	}

	if err := umt.initTask(); err != nil {
		return nil, err
	}

	if err := umt.initWorldMgr(); err != nil {
		return nil, err
	}

	if err := umt.initGameMgr(); err != nil {
		return nil, err
	}

	if err := umt.initGameService(); err != nil {
		return nil, err
	}

	if err := umt.initMsgParser(); err != nil {
		return nil, err
	}

	if err := umt.initTCPServer(); err != nil {
		return nil, err
	}

	if err := umt.initHTTPServer(); err != nil {
		return nil, err
	}

	if err := umt.initNsq(); err != nil {
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
func (umt *ultimate) initTask() error {
	var err error
	if umt.td, err = task.NewDispatcher(); err != nil {
		return err
	}

	logger.Info("task init ok!")

	return nil
}

// init datastore
func (umt *ultimate) initDatastore() error {
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

func (umt *ultimate) initMsgParser() error {
	umt.mp = NewMsgParser(umt.gm, umt.wm)
	logger.Info("msg parser init ok!")
	return nil
}

// InitTCPServer init
func (umt *ultimate) initTCPServer() error {
	var err error
	if umt.tcpServ, err = NewTcpServer(umt.mp, umt.td); err != nil {
		return err
	}

	logger.Info("tcp_server init ok!")
	return nil
}

// init http server
func (umt *ultimate) initHTTPServer() error {
	umt.httpServ = NewHttpServer(umt.gm)
	logger.Info("http_server init ok!")
	return nil
}

// init world session
func (umt *ultimate) initWorldMgr() error {
	var err error
	if umt.wm, err = world.NewWorldMgr(umt.ds); err != nil {
		return err
	}

	logger.Info("world_mgr init ok!")
	return nil
}

func (umt *ultimate) initGameMgr() error {
	var err error
	if umt.gm, err = game.NewGameMgr(umt.wm); err != nil {
		return err
	}

	logger.Info("game_mgr init ok!")
	return nil
}

func (umt *ultimate) initGameService() error {

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

func (umt *ultimate) initNsq() error {
	config := nsq.NewConfig()
	w, err := nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		return err
	}

	w.SetLogger(log.New(os.Stderr, "", log.LstdFlags), LogLevelInfo)

	err := w.Publish("write_test", []byte("test"))
	if err != nil {
		t.Fatalf("should lazily connect - %s", err)
	}

	w.Stop()

	err = w.Publish("write_test", []byte("fail test"))
	if err != ErrStopped {
		t.Fatalf("should not be able to write after Stop()")
	}

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
