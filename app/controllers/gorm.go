package controllers

import (
	"bytes"
	"database/sql"
	m "devices-server/app/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	r "github.com/revel/revel"
)

func ensureOption(option string) string {
	value, found := r.Config.String(option)
	if !found {
		r.ERROR.Fatalf("Option %v not found", option)
	}
	return value
}

func driversourcesForMysql(username string, password string, name string) string {
	var drivesources bytes.Buffer
	drivesources.WriteString(username)
	drivesources.WriteString(":")
	drivesources.WriteString(password)
	drivesources.WriteString("@")
	drivesources.WriteString("/")
	drivesources.WriteString(name)
	drivesources.WriteString("?charset=utf8&parseTime=True")
	return drivesources.String()
}

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

	driver := ensureOption("db.driver")
	username := ensureOption("db.username")
	password := ensureOption("db.password")
	name := ensureOption("db.name")

	// 今はmysql前提
	drivesources := driversourcesForMysql(username, password, name)

	// open db
	Gdb, err = gorm.Open(driver, drivesources)
	if err != nil {
		r.ERROR.Println("FATAL", err)
		panic(err)
	}

	Gdb.SingularTable(true)

	//開発中なので常にdropしている
	Gdb.DropTableIfExists(&m.User{})
	Gdb.CreateTable(&m.User{})

	Gdb.DropTableIfExists(&m.Device{})
	Gdb.CreateTable(&m.Device{})
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
