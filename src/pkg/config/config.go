package config
type Configuration struct{
	Protocal string
	Ip 		 string
}

var Configurations5 = [5]Configuration{
	{"tcp", "127.0.0.1:9900"},
	{"tcp", "127.0.0.1:9901"},
	{"tcp", "127.0.0.1:9902"},
	{"tcp", "127.0.0.1:9903"},
	{"tcp", "127.0.0.1:9904"},
}