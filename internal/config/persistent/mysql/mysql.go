package mysql

import (
	"database/sql"
	"fmt"
	"net/url"
	"time"

	"github.com/spf13/viper"
	// initialize mysql driver
	_ "github.com/go-sql-driver/mysql"

	"github.com/subroll/cfgsvc/internal/config/persistent"
)

// Storage is the structure of mysql storage layer
type Storage struct {
	db *sql.DB

	ItemStore  persistent.ItemStore
	GroupStore persistent.GroupStore
}

// Close will close the mysql client connection pool
func (s *Storage) Close() error {
	return s.db.Close()
}

// NewStorage create storage layer for mysql
func NewStorage() (*Storage, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&charset=%s&loc=%s",
		viper.GetString("sql.username"),
		viper.GetString("sql.password"),
		viper.GetString("sql.host"),
		viper.GetString("sql.db_name"),
		viper.GetString("sql.charset"),
		url.QueryEscape("Asia/Jakarta")))
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(viper.GetDuration("sql.max_life_time") * time.Second)
	db.SetMaxIdleConns(viper.GetInt("sql.max_idle_conns"))
	db.SetMaxOpenConns(viper.GetInt("sql.max_open_conns"))

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Storage{
		db:         db,
		ItemStore:  &itemStore{db: db},
		GroupStore: &groupStore{db: db},
	}, nil
}
