package controllers

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	// short name for revel
	r "github.com/revel/revel"
	// YOUR APP NAME
	"database/sql"
	"goBackend/app/models"
)

// type: revel controller with `*gorm.DB`
// c.Txn will keep `Gdb *gorm.DB`
type BaseController struct {
	*r.Controller
	Txn *gorm.DB
}

type Jsend struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func (c BaseController) RenderJsend(s string, d interface{}) r.Result {
	jsend := Jsend{Status: s, Data: d}
	return c.RenderJson(jsend)
}

// it can be used for jobs
var Gdb *gorm.DB

// init db
func InitDB() {
	var err error
	// open db
	Gdb, err = gorm.Open("sqlite3", "test.db")
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
func (c *BaseController) Begin() r.Result {
	txn := Gdb.Begin()
	if txn.Error != nil {
		panic(txn.Error)
	}
	c.Txn = txn
	return nil
}

// This method clears the c.Txn after each transaction
func (c *BaseController) Commit() r.Result {
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
func (c *BaseController) Rollback() r.Result {
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
