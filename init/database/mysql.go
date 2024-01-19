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
	once  sync.Once
	mu    sync.Mutex
	debug bool
}

var _db *DB

func nowFunc() time.Time {
	return time.Now().UTC()
}

func newDB(debug bool) (*DB, error) {
	dsn := ""
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

	err = db.AutoMigrate()
	if err != nil {
		return nil, err
	}

	return &DB{
		DB:    db,
		once:  sync.Once{},
		mu:    sync.Mutex{},
		debug: debug,
	}, nil
}

func GetDB(debug bool) *DB {
	f := func() {
		db, err := newDB(debug)
		if err != nil {
			log.Fatalf("init database error, err = %v", err)
		}
		_db = db
	}
	_db.once.Do(f)

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
