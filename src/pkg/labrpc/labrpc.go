package labrpc

import (
	"DFL/src/pkg/config"
	"net/rpc"
)
type ClientEnd struct{
	Target *rpc.Client
	Connectcf config.Configuration

}
func(cl *ClientEnd)Call(serviceMethod string, args any, reply any)bool{
	if cl.Target == nil{
		connect,err:=rpc.Dial(cl.Connectcf.Protocal,cl.Connectcf.Ip)
		if err == nil{
			cl.Target = connect
		}else{
			return false
		}
	}
	Err:=cl.Target.Call(serviceMethod,args,reply)
	if Err !=nil{
		cl.Target = nil
	}
	return Err==nil
}