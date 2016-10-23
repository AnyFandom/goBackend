package controllers

import "github.com/revel/revel"

func init() {
	revel.OnAppStart(InitDB) // invoke InitDB function before
	revel.InterceptMethod((*BaseController).Begin, revel.BEFORE)
	revel.InterceptMethod((*BaseController).CheckToken, revel.BEFORE)
	revel.InterceptMethod((*BaseController).Commit, revel.AFTER)
	revel.InterceptMethod((*BaseController).Rollback, revel.FINALLY)
}
