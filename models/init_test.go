package models_test

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"testing"

	"github.com/hublabs/stock-base-info-agent/models"

	"github.com/go-xorm/xorm"
	"github.com/pangpanglabs/goutils/echomiddleware"
	"github.com/pangpanglabs/goutils/jwtutil"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

var (
	dbContext context.Context
)

func TestMain(m *testing.M) {
	db := enterTest()
	code := m.Run()
	exitTest(db)
	os.Exit(code)
}

func enterTest() *xorm.Engine {
	runtime.GOMAXPROCS(1)
	xormEngine, err := xorm.NewEngine("mysql", "root:@tcp(localhost:3306)/hublabs_delivery?charset=utf8")
	if err != nil {
		fmt.Println("db error")
		panic(err)
	}
	dbContext = context.WithValue(context.Background(), echomiddleware.ContextDBName, xormEngine.NewSession())
	if err = models.DropTables(xormEngine); err != nil {
		panic(err)
	}
	if err = models.Init(xormEngine); err != nil {
		panic(err)
	}
	if err := CreateSeedData(xormEngine); err != nil {
		fmt.Println("create seed data err:", err)
	}

	jwtutil.SetJwtSecret(os.Getenv("JWT_SECRET"))
	return xormEngine
}

func exitTest(db *xorm.Engine) {
	// if err := models.DropTables(db); err != nil {
	// 	panic(err)
	// }
}
