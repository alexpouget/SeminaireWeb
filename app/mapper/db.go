package mapper

import (
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/revel/revel"
)

var db *gorm.DB

const dsnType = "mysql"

func InitDB() error {
	//dbm, err := gorm.Open(dsnType, "root:toto42@/seminaireweb?charset=utf8&parseTime=True&loc=Local")
  var err error
  db, err = gorm.Open(dsnType, "root:root@tcp(localhost:8889)/seminaireweb?charset=utf8&parseTime=True")
	if err != nil {
    panic("failed to connect database")
    revel.ERROR.Fatal(err)
  }
  db.LogMode(true)

  return err
}
  	
// Movie return Movie database object
func Movie() *gorm.DB {
	return db
}
