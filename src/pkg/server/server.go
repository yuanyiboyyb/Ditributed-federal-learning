package server

/*
import (
	"DFL/src/pkg/config"
	"DFL/src/pkg/labrpc"
	"DFL/src/pkg/raft"
	"net/rpc"
)
type Server struct{
	peers []*labrpc.ClientEnd
	raft  *raft.Raft
	me 		int
	persister *raft.Persister
	applyCh chan raft.ApplyMsg
}
func MakeServer(me int)(server *Server){
	server = &Server{}
	server.me = me
	server.persister = raft.MakePersister()
	rpcserver:=rpc.NewServer()
	raft:=raft.Make([]*labrpc.ClientEnd,server.me,server.persister,)
	for index,value:=range config.Configurations5{

	}
} */