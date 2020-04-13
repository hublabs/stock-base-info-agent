package factory

import (
	"fmt"
	"sync"
	"time"

	"github.com/hublabs/stock-base-info-agent/config"

	"github.com/go-xorm/xorm"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
)

var (
	defaultSqlServerEngine *xorm.Engine
	defaultMysqlEngine     *xorm.Engine
	once                   sync.Once
)

// Init 初始化 数据库引擎
func Init() {
	defaultSqlServerEngine = CreateSqlServerEngine(config.GetSqlServerConnString())
	SetDefaultSqlServerEngine(defaultSqlServerEngine)

	defaultMysqlEngine = CreateMySQLEngine(config.GetMysqlConnString())
	SetDefaultMysqlEngine(defaultMysqlEngine)
}

func SetDefaultSqlServerEngine(engine *xorm.Engine) {
	once.Do(func() {
		defaultSqlServerEngine = engine
	})
}

func SetDefaultMysqlEngine(engine *xorm.Engine) {
	once.Do(func() {
		defaultMysqlEngine = engine
	})
}

// 创建SqlServer数据库引擎
func CreateSqlServerEngine(connString string) *xorm.Engine {
	engine, err := xorm.NewEngine("mssql", connString)
	if err != nil {
		fmt.Println("createSqlServerEngine error")
		panic(err)
	}
	engine.TZLocation, _ = time.LoadLocation("UTC")

	return engine
}

func GetDefaultSqlServerEngine() *xorm.Engine {
	return defaultSqlServerEngine
}

// 创建MySQL数据库引擎
func CreateMySQLEngine(connString string) *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", connString)
	if err != nil {
		fmt.Println("createMysqlEngine error")
		panic(err)
	}

	return engine
}

func GetDefaultMysqlEngine() *xorm.Engine {
	return defaultMysqlEngine
}
