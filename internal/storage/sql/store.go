package storage

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/todo-app/internal/logger"
	"github.com/todo-app/internal/settings"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var (
	store             *SQLStorage
	mux               sync.Mutex
	schema            = "dbo"
	ErrNotFound       = errors.New("item not found")
	ErrInvalidRequest = errors.New("invalid input")
	ErrFKConstraint   = errors.New("fk constraint found")
	ErrDuplicateKey   = errors.New("duplicate key")
	ErrTruncatedValue = errors.New("truncated value")
)

// Predefined SQL errors: https://docs.microsoft.com/es-es/sql/relational-databases/errors-events/database-engine-events-and-errors
const (
	FKConstraintNumber   = 547
	DuplicateKeyNumber   = 2601
	TruncatedValueNumber = 2628
)

type SQLError interface {
	SQLErrorNumber() int32
	SQLErrorMessage() string
}

// Common define default columns
type Common struct {
	ID        UUID           `gorm:"column:Id;type:uuid;primary_key"`
	CreatedAt time.Time      `gorm:"column:CreatedAt"`
	UpdatedAt time.Time      `gorm:"column:UpdatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"column:DeletedAt;index"`
}

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

// CreateEntry creates a new entry.
func (s *SQLStorage) CreateEntry(entry interface{}) error {
	return s.checkSQLError(s.db.Create(entry).Error)
}

// UpdateEntry updates a single entry.
func (s *SQLStorage) UpdateEntryWithNulls(entry interface{}) error {
	return s.checkSQLError(s.db.Model(entry).
		Select("*").
		Omit("Id,CreatedAt").
		Updates(entry).Error)
}

// DeleteEntry delete a single entry
func (s *SQLStorage) DeleteEntry(entry interface{}) error {
	return s.checkSQLError(s.db.Delete(entry).Error)
}

func (s *SQLStorage) checkSQLError(err error) error {
	sqlError, ok := err.(SQLError)
	if ok {
		switch sqlError.SQLErrorNumber() {
		case FKConstraintNumber:
			return ErrFKConstraint
		case DuplicateKeyNumber:
			return ErrDuplicateKey
		case TruncatedValueNumber:
			return ErrTruncatedValue
		default:
			return err
		}
	}
	if !ok && err != nil {
		return ErrInvalidRequest
	}
	return nil
}

func first(db *gorm.DB, dst interface{}) error {
	err := db.First(dst).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrNotFound
	}
	return err
}
