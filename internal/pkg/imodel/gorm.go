package imodel

import (
	"context"
	"log"

	"github.com/jinzhu/gorm"
)

// NewGromDB NewGromDB
// Open initialize a new db connection, need to import driver first, e.g:
//
//     import _ "github.com/go-sql-driver/mysql"
//     func main() {
//       db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
//     }
// GORM has wrapped some drivers, for easier to remember driver's import path, so you could import the mysql driver with
//    import _ "github.com/jinzhu/gorm/dialects/mysql"
//    // import _ "github.com/jinzhu/gorm/dialects/postgres"
//    // import _ "github.com/jinzhu/gorm/dialects/sqlite"
//    // import _ "github.com/jinzhu/gorm/dialects/mssql"
func NewGromDB(driver, dsn string) (context.CancelFunc, *gorm.DB, error) {
	log.Println("NewGromDB", driver, dsn)
	db, err := gorm.Open(driver, dsn)
	if err != nil {
		return nil, nil, err
	}

	cctx, cancel := context.WithCancel(context.Background())

	go func() {
		select {
		case <-cctx.Done():
			db.Close()
		}
	}()

	return cancel, db, err
}
