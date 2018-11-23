package main

import (
	"fmt"
	"github.com/IhorBondartsov/datasaver/cfg"
	"github.com/IhorBondartsov/datasaver/database/customdb"
	"github.com/IhorBondartsov/datasaver/web/myproto"
)

func main() {
	fmt.Println(greeating)
	cfg := myproto.GRPCServerCfg{
		Port: cfg.PORT,
		DB:   customdb.NewDB(),
	}

	x := myproto.NewServer(cfg)
	x.Start()
}

var greeating = `
HELLO! MY DEAR FRIEND!!
___________________________________
 ∧__∧
/ . .\
( >ω<)
(っ▄︻▇〓▄︻┻┳═*  
(     )    /\💥
...................................
Start server ...`

var parting = `

`




// MAIN FOR JSON RPC
// func main(){
// cfg := jsonrpcserver.CfgAPI{
// 	Addr: ":1812",
// }
// x := jsonrpcserver.NewAPI(cfg, customdb.NewDB())
// jsonrpcserver.Start(x)
// }
