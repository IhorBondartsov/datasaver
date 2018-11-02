package customdb

import (
	"fmt"
	"github.com/IhorBondartsov/datasaver/database"
	"github.com/IhorBondartsov/datasaver/entity"
	"sync"
)

func NewDB() database.DataBase {
	m := map[int]entity.PersonData{}
	return &myDataBase{
		Data: m,
	}
}

type myDataBase struct {
	sync.Mutex
	Data map[int]entity.PersonData
}

func (db *myDataBase) Save(pd entity.PersonData) error {
	db.Lock()
	db.Data[pd.Id] = pd
	db.Unlock()
	return nil
}

func (db *myDataBase) Print() {
	db.Lock()
	for k, v := range db.Data {
		fmt.Println(k, v)
	}
	db.Unlock()
}
