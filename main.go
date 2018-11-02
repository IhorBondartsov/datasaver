package main

import (
	"github.com/IhorBondartsov/datasaver/database/customdb"
	"github.com/IhorBondartsov/datasaver/server/jsonrpcserver"
)

func main() {
	cfg := jsonrpcserver.CfgAPI{
		Addr: ":1812",
	}
	x := jsonrpcserver.NewAPI(cfg, customdb.NewDB())
	jsonrpcserver.Start(x)
}
