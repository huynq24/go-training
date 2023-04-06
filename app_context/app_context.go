package app_context

import "gorm.io/gorm"

type AppContext interface {
	GetMainDBConnection() *gorm.DB
}

type appCtx struct {
	db *gorm.DB
}

func NewAppContext(db *gorm.DB) *appCtx {
	return &appCtx{db}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}
