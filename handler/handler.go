package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"todomvc/db"
	"todomvc/model"
)
//增
func Add(c *gin.Context) {
	var p model.TodomvcAdd
	c.ShouldBindJSON(&p)
	var m = &model.Todomvc{Item: p.Item}
	db.DB.Create(&m)
	c.JSON(http.StatusOK, nil)
}
//删
func Del(c *gin.Context) {
	var p model.TodomvcDel
	c.ShouldBindJSON(&p)
	db.DB.Where("id", p.Id).Delete(&model.Todomvc{})
	c.JSON(http.StatusOK, nil)
}
//改
func Update(c *gin.Context) {
	var p []model.TodomvcUpdate
	c.ShouldBindJSON(&p)
	for _, t := range p {
		db.DB.Model(&model.Todomvc{}).Where("id", t.Id).Select("Name", "Age").Updates(model.Todomvc{
			Item:   t.Item,
			Status: t.Status,
		})
	}
	c.JSON(http.StatusOK, nil)
}

//查
func Find(c *gin.Context) {
	var p model.TodomvcFind
	c.ShouldBindJSON(&p)
	var m []model.Todomvc
	var tx = db.DB
	if p.Status != -1 {
		tx = tx.Where("status", p.Status)
	}
	if p.Item != "" {
		tx = tx.Where("item LIKE?", "%"+p.Item+"p")
	}
	tx.Find(&m)
	c.JSON(http.StatusOK, &m)
}


