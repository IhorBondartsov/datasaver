package database

import "github.com/IhorBondartsov/datasaver/entity"

type DataBase interface {
	Save(pd entity.PersonData) error
	Print()
}
