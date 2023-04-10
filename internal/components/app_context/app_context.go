package app_context

import (
	"golang-training/internal/common"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
}

type appCtx struct {
	db     *gorm.DB
	config common.Config
}

func NewAppContext(db *gorm.DB, config common.Config) *appCtx {
	return &appCtx{db: db, config: config}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}
