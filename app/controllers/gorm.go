package controllers

import (
	"reflect"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// short name for revel
	r "github.com/revel/revel"
	// YOUR APP NAME
	"database/sql"
	"goBackend/app/models"
	"goBackend/app/utils"
)

// type: revel controller with `*gorm.DB`
// c.Db will keep `Gdb *gorm.DB`
type BaseController struct {
	*r.Controller
	Db         *gorm.DB
	authorized bool
	userId     uint
}

type Jsend struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func (c BaseController) RenderJsend(s string, d interface{}, m string) r.Result {
	jsend := Jsend{Status: s, Data: d, Message: m}
	return c.RenderJson(jsend)
}

func (c *BaseController) CheckToken() r.Result {
	var token string
	c.Params.Bind(&token, "token")

	if len(token) == 0 {
		return nil
	}

	claims := utils.ParseToken(token)

	floatUserId, ok := claims["id"].(float64)

	if !ok {
		panic(reflect.TypeOf(claims["id"]))
	}

	userId := uint(floatUserId)
	var user models.User
	c.Db.First(&user, userId)

	if len(user.Username) == 0 {
		panic(404)
	}

	c.userId = userId
	c.authorized = true

	return nil
}

// it can be used for jobs
var Gdb *gorm.DB

// init db
func InitDB() {
	var err error
	// open db
	Gdb, err = gorm.Open("postgres", "host=localhost user=revel dbname=revel sslmode=disable password=revel")
	if err != nil {
		r.ERROR.Println("FATAL", err)
		panic(err)
	}
	Gdb.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{}, &models.Fandom{}, &models.Blog{})
	Gdb.LogMode(true)
	// uniquie index if need
	//Gdb.Model(&models.User{}).AddUniqueIndex("idx_user_name", "name")
}

// transactions

// This method fills the c.Db before each transaction
func (c *BaseController) Begin() r.Result {
	Db := Gdb.Begin()
	if Db.Error != nil {
		panic(Db.Error)
	}
	c.Db = Db
	return nil
}

// This method clears the c.Db after each transaction
func (c *BaseController) Commit() r.Result {
	if c.Db == nil {
		return nil
	}
	c.Db.Commit()
	if err := c.Db.Error; err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Db = nil
	return nil
}

// This method clears the c.Db after each transaction, too
func (c *BaseController) Rollback() r.Result {
	if c.Db == nil {
		return nil
	}
	c.Db.Rollback()
	if err := c.Db.Error; err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Db = nil
	return nil
}
