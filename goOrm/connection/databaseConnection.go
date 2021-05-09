package connection

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"database_demo/session"
)

var GormDB *gorm.DB
var initDBOnce sync.Once

// dsn = "username:password@(ip:port)/database?timeout=5000ms&readTimeout=5000ms&writeTimeout=5000ms&charset=utf8mb4&parseTime=true&loc=Local"
var dsn = "root:root@tcp(localhost:3306)/gormDemo?timeout=5000ms&readTimeout=5000ms&writeTimeout=5000ms&charset=utf8mb4&parseTime=True&loc=Local"

func init() {
	initDBOnce.Do(initProcedure)
}

func initProcedure() {
	GormDB = initConnection()
	if GormDB == nil {
		log.Fatal("gorm db is nil")
		return
	}
}

func initConnection() *gorm.DB {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	// 创建 gorm.DB 实例时，默认情况下会向数据库服务端发起一次连接，以保证 dsn 的正确性.
	gormDB, err := gorm.Open(mysql.New(mysql.Config{DSN: dsn}), &gorm.Config{
		//默认事务
		SkipDefaultTransaction: false,
		//创建时间使用的函数
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
		//是否仅生成SQL
		DryRun: false,
		//执行任何 SQL 时都会创建一个 prepared statement 并将其缓存
		PrepareStmt: false,
		//GORM 会使用 SavePoint(savedPointName)，RollbackTo(savedPointName) 为你提供嵌套事务支持
		DisableNestedTransaction: false,
		//GORM 会自动 ping 数据库以检查数据库的可用性
		DisableAutomaticPing: false,
		//AllowGlobalUpdate 允许对全量数据进行更新
		AllowGlobalUpdate: false,
		//在 AutoMigrate 或 CreateTable 时，GORM 会自动创建外键约束
		DisableForeignKeyConstraintWhenMigrating: false,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: true,
		},
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Millisecond * 200,
				LogLevel:                  logger.Warn,
				IgnoreRecordNotFoundError: true,
				Colorful:                  true,
			},
		),
	})
	if err != nil {
		log.Fatal(err)
		return nil
	}
	sqlDB, err := gormDB.DB()
	if err != nil {
		log.Fatalf("SQL DB Fail :<%s>\n", err.Error())
	}
	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("Ping DB Fail :<%s>\n", err.Error())
	}
	sqlDB.SetMaxIdleConns(2)
	sqlDB.SetMaxOpenConns(1)
	sqlDB.SetConnMaxIdleTime(time.Minute)
	sqlDB.SetConnMaxLifetime(time.Minute)
	return gormDB
}

func GetOriginalDB() *gorm.DB {
	return GormDB
}
func GetOriginalDBWithCtx(ctx context.Context) *gorm.DB {
	return GormDB.WithContext(ctx)
}

func GetOriginalDBSession(ctx context.Context) *gorm.DB {
	return session.NewSession(ctx, GormDB)
}

func CloseDB() {
	sqlDB, err := GormDB.DB()
	if err != nil {
		log.Fatalf("SQL DB Fail :<%s>\n", err.Error())
	}
	err = sqlDB.Close()
	if err != nil {
		log.Fatalf("SQL DB CLOSE Fail :<%s>\n", err.Error())
	}
}
