package conn

import (
	"fmt"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	// "gorm.io/gorm/logger"

	"github.com/pascallin/gin-template/config"
)

type MysqlConn interface {
	GetMysqlDB() *gorm.DB
}

var (
	mysqlDB *gorm.DB
	mOnce   sync.Once
)

func GetMysqlDB() *gorm.DB {
	// newLogger := logger.New(
	// 	newLogrusWriter(), // io writer
	// 	logger.Config{
	// 		SlowThreshold:             time.Second,   // Slow SQL threshold
	// 		LogLevel:                  logger.Silent, // Log level
	// 		IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
	// 		Colorful:                  false,         // Disable color
	// 	},
	// )
	mOnce.Do(func() {
		db, err := openMysql()
		if err != nil {
			log.Error(err)
		}
		mysqlDB = db
	})

	return mysqlDB
}

func openMysql() (*gorm.DB, error) {
	c := config.GetMysqlConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.User, c.Password, c.Host, c.Port, c.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Info("Mysql database connected")

	return db, nil
}

type LogrusWriter struct {
	mlog *log.Logger
}

func (m *LogrusWriter) Printf(format string, v ...interface{}) {
	logstr := fmt.Sprintf(format, v...)
	m.mlog.Info(logstr)
}

func newLogrusWriter() *LogrusWriter {
	log := log.New()
	return &LogrusWriter{mlog: log}
}
