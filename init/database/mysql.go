package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"sync"
	"time"
)

var (
	_db  *gorm.DB
	once sync.Once
	mu   sync.Mutex
	//todo: for test
	debug bool
)

func nowFunc() time.Time {
	return time.Now().UTC()
}

func newDB() (*gorm.DB, error) {
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

	return db, nil
}

func GetDB() *gorm.DB {
	f := func() {
		db, err := newDB()
		if err != nil {
			log.Fatalf("init database error, err = %v", err)
		}
		_db = db
	}
	once.Do(f)

	mu.Lock()
	defer mu.Unlock()
	return _db
}

func Close() error {
	mu.Lock()
	defer mu.Unlock()

	db, err := _db.DB()
	if err != nil {
		return err
	}
	return db.Close()
}
