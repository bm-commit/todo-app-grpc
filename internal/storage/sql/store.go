package storage

import (
	"fmt"
	"sync"

	"github.com/todo-app/internal/logger"
	"github.com/todo-app/internal/settings"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var (
	store *SQLStorage

	mux sync.Mutex

	schema = "dbo"
)

// SQLStorage is the storage that contains the database connection.
type SQLStorage struct {
	db *gorm.DB
}

// NewSQLStorage creates a new storage instance.
func NewSQLStorage(conf settings.DBConn) (*SQLStorage, error) {
	mux.Lock()
	defer mux.Unlock()

	if store != nil {
		return store, nil
	}

	logger.Info("creating db connection...")

	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		conf.User,
		conf.Pass,
		conf.Host,
		conf.Port,
		conf.Database)

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	if err = sqlDB.Ping(); err != nil {
		logger.Error("checking db connection, %v", err)
		return nil, err
	}
	schema = conf.Schema

	logger.Info("connected to %s:%s", conf.Host, conf.Port)
	store = &SQLStorage{db: db}
	return store, nil
}
