package main

import (
	//"github.com/IhorBondartsov/datasaver/server/jsonrpcserver"
	"github.com/IhorBondartsov/datasaver/web/myproto"
)

func main() {
	// cfg := jsonrpcserver.CfgAPI{
	// 	Addr: ":1812",
	// }
	// x := jsonrpcserver.NewAPI(cfg, customdb.NewDB())
	// jsonrpcserver.Start(x)

	cfg := myproto.GRPCServerCfg{
		Port: "7777",
	}

	x := myproto.NewServer(cfg)
	x.Start()
}
