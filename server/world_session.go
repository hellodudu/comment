package ultimate

import (
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"net"
	"reflect"
	"sync"

	"github.com/fatih/color"
	"github.com/golang/protobuf/proto"
	"github.com/hellodudu/Ultimate/config"
	"github.com/hellodudu/Ultimate/utils"
)

// base net message type define
type BaseNetMsg struct {
	Id   uint32 // message name crc32
	Size uint32 // message size
}

// world session msg register info
type regInfo struct {
	cb func(net.Conn, *WorldSession, proto.Message)
}

type WorldSession struct {
	mapWorld  map[uint32]*World // all connected world
	mapConn   map[net.Conn]*World
	protoReg  map[uint32]*regInfo
	mu        sync.Mutex
	wg        sync.WaitGroup
	ctx       context.Context
	cancel    context.CancelFunc
	cTimeOutW chan uint32
}

func NewWorldSession() (*WorldSession, error) {
	w := &WorldSession{
		mapWorld:  make(map[uint32]*World),
		mapConn:   make(map[net.Conn]*World),
		protoReg:  make(map[uint32]*regInfo),
		cTimeOutW: make(chan uint32, config.WorldConnectMax),
	}

	w.ctx, w.cancel = context.WithCancel(context.Background())

	w.registerAllMessage()
	return w, nil
}

func (ws *WorldSession) Stop() {
	for _, world := range ws.mapWorld {
		world.Stop()
	}

	ws.cancel()
}

func (ws *WorldSession) registerAllMessage() {
	// cb -- callback function
	ws.registerProto(utils.Crc32("world_message.MWU_WorldLogon"), &regInfo{cb: HandleWorldLogon})

	ws.registerProto(utils.Crc32("world_message.MWU_TestConnect"), &regInfo{cb: HandleTestConnect})

	ws.registerProto(utils.Crc32("world_message.MWU_HeartBeat"), &regInfo{cb: HandleHeartBeat})

	ws.registerProto(utils.Crc32("world_message.MWU_WorldConnected"), &regInfo{cb: HandleWorldConnected})

	ws.registerProto(utils.Crc32("world_message.MWU_RequestPlayerInfo"), &regInfo{cb: HandleRequestPlayerInfo})

	ws.registerProto(utils.Crc32("world_message.MWU_RequestGuildInfo"), &regInfo{cb: HandleRequestGuildInfo})
}

func (ws *WorldSession) getRegisterProto(msgID uint32) (*regInfo, error) {
	v, ok := ws.protoReg[msgID]
	if ok {
		return v, nil
	}

	return v, errors.New("cannot find proto type registed in world_session!")
}

func (ws *WorldSession) registerProto(msgID uint32, info *regInfo) {
	if v, ok := ws.protoReg[msgID]; ok {
		log.Printf("register proto msg_id<%d> existed! protobuf type:%v\n", msgID, v)
		return
	}

	ws.protoReg[msgID] = info
}

// func binaryUnmarshal(data []byte) {
// 	msg := &MSG_MWU_WorldLogon{}
// 	byData := make([]byte, binary.Size(msg))

// 	// discard top 4 bytes(message size)
// 	copy(byData, data[4:])

// 	buf := &bytes.Buffer{}
// 	if _, err := buf.Write(byData); err != nil {
// 		log.Fatal(err)
// 	}

// 	// get top 4 bytes messageid
// 	msgID := binary.LittleEndian.Uint32(buf.Bytes()[:4])
// 	if msgID == utils.Crc32(string("MWU_WorldLogon")) {
// 		if err := binary.Read(buf, binary.LittleEndian, msg); err != nil {
// 			log.Fatal(err)
// 		}
// 		log.Printf("world<id:%d, name:%s> logon!\n", msg.WorldID, msg.WorldName)
// 	}

// 	log.Printf("translate msg:%+v\n", msg)
// }

func protoUnmarshal(data []byte, m proto.Message) {
	if err := proto.Unmarshal(data, m); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}

	log.Printf("translate msg to proto:%T\n", m)
}

// decode binarys to proto message
func (ws *WorldSession) decodeToProto(data []byte) (proto.Message, error) {
	byProto := data[12:]
	protoNameLen := binary.LittleEndian.Uint16(byProto[:2])

	if uint16(len(byProto)) < 2+protoNameLen {
		return nil, errors.New("recv proto msg length < 2+protoNameLen:" + string(byProto))
	}

	protoTypeName := string(byProto[2 : 2+protoNameLen])
	protoData := byProto[2+protoNameLen:]
	pType := proto.MessageType(protoTypeName)
	if pType == nil {
		return nil, errors.New(fmt.Sprintf("invalid message<%s>, won't deal with it!" + protoTypeName))
	}

	// unmarshal
	newProto, ok := reflect.New(pType.Elem()).Interface().(proto.Message)
	if !ok {
		return nil, errors.New(fmt.Sprintf("invalid message<%s>, won't deal with it!" + protoTypeName))
	}

	protoUnmarshal(protoData, newProto)
	return newProto, nil
}

// top 4 bytes are msgSize, next 8 bytes are BaseNetMsg
// if it is protobuf msg, then next 2 bytes are proto name length, the next is proto name, final is proto data.
// if it is transfer msg(transfer binarys to other world), then next are binarys to be transferd
func (ws *WorldSession) HandleMessage(con net.Conn, data []byte) {
	if len(data) <= 12 {
		log.Println(color.YellowString("tcp recv data length <= 12:%s", string(data)))
		return
	}

	baseMsg := &BaseNetMsg{}
	byBaseMsg := make([]byte, binary.Size(baseMsg))

	// discard top 4 bytes(message size)
	copy(byBaseMsg, data[4:4+8])
	buf := &bytes.Buffer{}
	if _, err := buf.Write(byBaseMsg); err != nil {
		log.Println(color.YellowString("cannot read message:", byBaseMsg, " from connection:", con, " err:", err.Error()))
		return
	}

	// get top 4 bytes messageid
	if err := binary.Read(buf, binary.LittleEndian, baseMsg); err != nil {
		log.Println(color.YellowString("cannot read message:", byBaseMsg, " from connection:", con, " err:", err.Error()))
		return
	}

	// proto message
	if baseMsg.Id == utils.Crc32(string("MWU_DirectProtoMsg")) {
		newProto, err := ws.decodeToProto(data)
		if err != nil {
			log.Println(color.YellowString(err.Error()))
			return
		}

		protoMsgID := utils.Crc32(proto.MessageName(newProto))
		r, err := ws.getRegisterProto(protoMsgID)
		if err != nil {
			log.Println(color.YellowString(fmt.Sprintf("unregisted protoMsgID<%d> received!", protoMsgID)))
		}

		// callback
		r.cb(con, ws, newProto)
	} else {

		// transfer message
	}

}

func (ws *WorldSession) AddWorld(id uint32, name string, con net.Conn) (*World, error) {
	var invalidID int32 = -1
	if id == uint32(invalidID) {
		return nil, errors.New("add world id invalid!")
	}

	if _, ok := ws.mapWorld[id]; ok {
		ws.KickWorld(id)
	}

	if _, ok := ws.mapConn[con]; ok {
		ws.KickWorld(id)
	}

	if uint32(len(ws.mapConn)) >= config.WorldConnectMax {
		return nil, errors.New("world connected num full!")
	}

	w := NewWorld(id, name, con, ws.cTimeOutW)
	ws.mu.Lock()
	ws.mapWorld[w.Id] = w
	ws.mapConn[w.Con] = w
	ws.mu.Unlock()
	log.Println(color.CyanString("add world <id:%d, name:%s, con:%v> success!", w.Id, w.Name, w.Con))
	go w.Run()
	return w, nil
}

func (ws *WorldSession) GetWorldByID(id uint32) *World {
	w, ok := ws.mapWorld[id]
	if !ok {
		return nil
	}
	return w
}

func (ws *WorldSession) GetWorldByCon(con net.Conn) *World {
	w, ok := ws.mapConn[con]
	if !ok {
		return nil
	}
	return w
}

func (ws *WorldSession) DisconnectWorld(con net.Conn) {
	w, ok := ws.mapConn[con]
	if !ok {
		return
	}

	log.Println(color.YellowString("World<id:%d> disconnected!", w.Id))
	w.Stop()

	ws.mu.Lock()
	delete(ws.mapWorld, w.Id)
	delete(ws.mapConn, con)
	ws.mu.Unlock()
}

func (ws *WorldSession) KickWorld(id uint32) {
	w, ok := ws.mapWorld[id]
	if !ok {
		return
	}

	if _, ok := ws.mapConn[w.Con]; !ok {
		return
	}

	log.Println(color.RedString("World<id:%d> was kicked by timeout reason!", w.Id))
	w.Stop()

	ws.mu.Lock()
	delete(ws.mapConn, w.Con)
	delete(ws.mapWorld, w.Id)
	ws.mu.Unlock()
}

func (ws *WorldSession) Run() {
	for {
		select {
		case <-ws.ctx.Done():
			log.Println(color.RedString("world session context done!"))
			return
		case wid := <-ws.cTimeOutW:
			ws.KickWorld(wid)
		}
	}
}
