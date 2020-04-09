package models

import (
	"github.com/go-xorm/xorm"
)

func Init(db *xorm.Engine) error {
	if err := db.Sync(
		new(MngBatchJob)); err != nil {
		return err
	}
	return nil
}

func DropTables(db *xorm.Engine) error {
	return db.DropTables(
		new(MngBatchJob))
}
