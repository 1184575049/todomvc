package model

import (
	"gorm.io/gorm"
)
type Todomvc struct {
	gorm.Model
	Item   string
	Status uint
}

type TodomvcAdd struct {
	Item string
}

type TodomvcDel struct {
	Id uint
}

type TodomvcUpdate struct {
	Item   string
	Id     uint
	Status uint
}

type TodomvcFind struct {
	Item   string
	Status int
}
