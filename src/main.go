package dfl

import (
	"DFL/src/pkg/config"
	"DFL/src/pkg/labrpc"
	"DFL/src/pkg/raft"
	"DFL/src/pkg/shardctrler"
	"log"
	"net"
	"net/rpc"
)
var me = 1
func main(){
	servers := make(map[int]*labrpc.ClientEnd)
	persister:=raft.MakePersister()
	server:=shardctrler.StartServer(servers, me, persister)
	err:=rpc.RegisterName("ShardCtrler", server)
	if err != nil {
		log.Fatalf("Failed to register ShardCtrler")
	}
	err=rpc.RegisterName("Raft",server.Raft())
	listener, err := net.Listen(config.Configurations5[me].Protocal,config.Configurations5[me].Ip)
	defer listener.Close()

	rpc.Accept(listener)
}