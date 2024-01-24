package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"sync"
	"time"
)

type DB struct {
	DB    *gorm.DB
	mu    sync.Mutex
	debug bool
}

var (
	_db  *DB
	once sync.Once
)

func nowFunc() time.Time {
	return time.Now().UTC()
}

func newDB(dsn string, debug bool) (*DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: &schema.NamingStrategy{
			TablePrefix: "sso_",
		},
		NowFunc:           nowFunc,
		PrepareStmt:       false,
		AllowGlobalUpdate: false,
		Plugins:           nil,
	})
	if err != nil {
		return nil, err
	}

	// check whether is debug model or not
	if debug {
		db = db.Debug()
	}

	return &DB{
		DB:    db,
		mu:    sync.Mutex{},
		debug: debug,
	}, nil
}

func GetDB(dsn string, debug bool) *DB {
	f := func() {
		db, err := newDB(dsn, debug)
		if err != nil {
			log.Fatalf("init database error, err = %v", err)
		}
		_db = db
	}
	once.Do(f)

	_db.mu.Lock()
	defer _db.mu.Unlock()
	return _db
}

func (ins *DB) Close() error {
	ins.mu.Lock()
	defer ins.mu.Unlock()

	d, err := ins.DB.DB()
	if err != nil {
		return err
	}
	return d.Close()
}

func (ins *DB) AutoMigrate(dst ...interface{}) error {
	return ins.DB.AutoMigrate(dst...)
}
