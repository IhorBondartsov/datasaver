package jsonrpcserver

import (
	"fmt"
	"github.com/IhorBondartsov/datasaver/entity"
	"github.com/IhorBondartsov/datasaver/server/api"

	"github.com/IhorBondartsov/datasaver/database"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json2"
	"net/http"
)

type CfgAPI struct {
	Addr string
}

func NewAPI(cfg CfgAPI, db database.DataBase) *API {
	return &API{
		Addr: cfg.Addr,
		db:   db,
	}
}

type API struct {
	Addr string
	db   database.DataBase
}

func (a *API) Echo(r *http.Request, args *api.EchoReq, res *api.EchoResp) error {
	fmt.Println("Echo")
	res.Message = "asd"
	return nil
}

func (a *API) Save(r *http.Request, args *entity.PersonData, res *api.SaveResp) error {
	fmt.Println("Save")
	err := a.db.Save(*args)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func (a *API) Print(r *http.Request, args *api.EchoReq, res *api.SaveResp) error {
	fmt.Println("Print")
	a.db.Print()
	return nil
}

func Start(c *API) {
	s := rpc.NewServer()
	s.RegisterCodec(json2.NewCodec(), "application/json")
	s.RegisterCodec(json2.NewCodec(), "application/json;charset=UTF-8")
	s.RegisterService(c, "API")

	r := mux.NewRouter()
	r.Handle("/rpc", s)
	http.ListenAndServe(":1812", r)
}
