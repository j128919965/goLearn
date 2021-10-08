package data

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo)

// NewData .
func NewData() (*gorm.DB , error) {
	db, err := gorm.Open("mysql", "root:o@S$9i3r6V@tcp(ssacgn.online:3306)/cloud?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}
	db.SingularTable(true)
	return db,err
}
