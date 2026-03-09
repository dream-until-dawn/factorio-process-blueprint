package db

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "modernc.org/sqlite"
)

type SQLiteDB struct {
	DB *sql.DB
}

var instance *SQLiteDB
var once sync.Once

// 初始化
func Init(dbPath string) {
	once.Do(func() {
		db, err := sql.Open("sqlite", dbPath)
		if err != nil {
			log.Fatalf("打开sqlite数据库失败: %v", err)
		}

		// 设置WAL模式
		_, err = db.Exec("PRAGMA journal_mode = WAL;")
		if err != nil {
			log.Printf("设置WAL模式失败: %v", err)
		}

		instance = &SQLiteDB{DB: db}

		if err := instance.migrate(); err != nil {
			log.Fatalf("数据库初始化失败: %v", err)
		}
	})
}

// 取得数据库实例
func Get() *SQLiteDB {
	Init("db/factorio.db")
	if instance == nil {
		log.Fatal("SQLite还未初始化")
	}
	return instance
}

// 物品表
func (s *SQLiteDB) creationItemsTable() error {
	const tableTpl = `
		CREATE TABLE IF NOT EXISTS items (
			name TEXT PRIMARY KEY, 
			Icon TEXT NOT NULL,
			name_zh TEXT DEFAULT '',
			hidden INTEGER DEFAULT 0 CHECK (hidden IN (-1, 0, 1)),
			stack_size INTEGER DEFAULT 1 CHECK (stack_size > 0),
			weight INTEGER DEFAULT -1 CHECK (weight >= -1),
			subgroup TEXT DEFAULT '',
			place_result TEXT DEFAULT '',
			item_order TEXT DEFAULT '',
			fuel_category TEXT DEFAULT '',
			fuel_value TEXT DEFAULT '',
			burnt_result TEXT DEFAULT '',
			spoil_ticks INTEGER DEFAULT 0 CHECK (spoil_ticks >= 0),
			spoil_result TEXT DEFAULT ''
		);
	`
	stmt := fmt.Sprintf(tableTpl)
	_, err := s.DB.Exec(stmt)
	return err
}

func (s *SQLiteDB) migrate() error {
	if err := s.creationItemsTable(); err != nil {
		return err
	}
	return nil
}
