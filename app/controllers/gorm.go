package controllers

import (
	"database/sql"
	"devices-server/app/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	r "github.com/revel/revel"
)

// type: revel controller with `gorm.DB`
// c.Txn will keep `db gorm.DB`
type GormController struct {
	*r.Controller
	Txn *gorm.DB
}

// it can be used for jobs
var Gdb gorm.DB

// init db
func InitDB() {
	var err error
	// open db
	var db_user = revel.Config.String("db.user")
	var db_password = revel.Config.String("db.pasword")
	var db_name = revel.Config.String("db.name")

	var drivesources bytes.Buffer
	drivesources.WriteString(db_user)
	drivesources.WriteString(":")
	drivesources.WriteString(db_password)
	drivesources.WriteString("@")
	drivesources.WriteString("/")
	drivesources.WriteString(db_name)

	Gdb, err = gorm.Open("mysql", drivesources.toString())
	if err != nil {
		r.ERROR.Println("FATAL", err)
		panic(err)
	}
	Gdb.AutoMigrate(&models.User{})
	// uniquie index if need
	//Gdb.Model(&models.User{}).AddUniqueIndex("idx_user_name", "name")
}

// transactions

// This method fills the c.Txn before each transaction
func (c *GormController) Begin() r.Result {
	txn := Gdb.Begin()
	if txn.Error != nil {
		panic(txn.Error)
	}
	c.Txn = txn
	return nil
}

// This method clears the c.Txn after each transaction
func (c *GormController) Commit() r.Result {
	if c.Txn == nil {
		return nil
	}
	c.Txn.Commit()
	if err := c.Txn.Error; err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}

// This method clears the c.Txn after each transaction, too
func (c *GormController) Rollback() r.Result {
	if c.Txn == nil {
		return nil
	}
	c.Txn.Rollback()
	if err := c.Txn.Error; err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}
