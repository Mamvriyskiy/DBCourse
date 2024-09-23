package factory

import (
	method "github.com/Mamvriyskiy/database_course/main/testsdatabase/method"
	"github.com/jmoiron/sqlx"
)

type ObjectSystem interface {
	InsertObject(connDB *sqlx.DB) (int, error)
	DeleteObject()
}

func New(typeObject, keyСharacter string) ObjectSystem {
	switch typeObject {
	case "user":
		return method.NewUser(keyСharacter)
	default:
		return nil
	} 
}
