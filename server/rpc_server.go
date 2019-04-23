package ultimate

import (
	"context"
	"net"
	"sync"

	"github.com/hellodudu/Ultimate/global"
	"github.com/hellodudu/Ultimate/logger"
	world_message "github.com/hellodudu/Ultimate/proto"
	"google.golang.org/grpc"
)

type RpcServer struct {
	s  map[*grpc.Server]struct{}
	ln net.Listener
	mu sync.Mutex
}

type rpcResponser struct {
}

func NewRpcServer() (*RpcServer, error) {
	s := &RpcServer{
		s: make(map[*grpc.Server]struct{}),
	}

	addr, err := global.IniMgr.GetIniValue("config/ultimate.ini", "listen", "RpcListenAddr")
	if err != nil {
		logger.Error("cannot read ini RpcListenAddr!")
		return nil, err
	}

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		logger.Error("rpc server failed to listen: ", err)
		return nil, err
	}

	logger.Print("rpc server listening at ", addr)

	s.ln = ln
	return s, nil
}

// SayHello implements helloworld.GreeterServer
func (s *rpcResponser) SayHello(ctx context.Context, in *world_message.HelloRequest) (*world_message.HelloReply, error) {
	logger.Info("Received: ", in.Name)
	return &world_message.HelloReply{Message: "Reply " + in.Name}, nil
}

func (s *rpcResponser) GetScore(ctx context.Context, in *world_message.GetScoreRequest) (*world_message.GetScoreReply, error) {
	logger.Info("Received: ", in.Id)
	arena := Instance().GetGameMgr().GetArena()
	data, err := arena.GetDataByID(in.Id)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return &world_message.GetScoreReply{Score: data.Score}, nil
}

func (server *RpcServer) Run() {
	s := grpc.NewServer()

	server.mu.Lock()
	server.s[s] = struct{}{}
	server.mu.Unlock()

	world_message.RegisterGreeterServer(s, &rpcResponser{})
	world_message.RegisterInviterServer(s, &rpcResponser{})
	if err := s.Serve(server.ln); err != nil {
		logger.Error("failed to service rpc Greeter: ", err)
		return
	}
}

func (server *RpcServer) Stop() {
	server.mu.Lock()

	for s := range server.s {
		s.Stop()
		delete(server.s, s)
	}

	server.mu.Unlock()
	server.ln.Close()
}