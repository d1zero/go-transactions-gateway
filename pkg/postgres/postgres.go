package postgres

import (
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"time"
)

type Config struct {
	Host            string        `mapstructure:"host" validate:"required"`
	Port            int           `mapstructure:"port" validate:"required"`
	User            string        `mapstructure:"user" validate:"required"`
	Password        string        `mapstructure:"password" validate:"required"`
	DBName          string        `mapstructure:"db_name" validate:"required"`
	SSLMode         string        `mapstructure:"ssl_mode" validate:"required"`
	ApplicationName string        `mapstructure:"application_name" validate:"required"`
	MaxOpenConns    int           `mapstructure:"max_open_conns" validate:"required"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime" validate:"required"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns" validate:"required"`
	ConnMaxIdleTime time.Duration `mapstructure:"conn_max_idle_time" validate:"required"`
}

func New(cfg Config) (db *sqlx.DB, err error) {
	connAddr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s application_name=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.DBName,
		cfg.SSLMode,
		cfg.ApplicationName,
	)
	db, err = sqlx.Connect("pgx", connAddr)
	if err != nil {
		return
	}

	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetConnMaxLifetime(cfg.ConnMaxLifetime * time.Second)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetConnMaxIdleTime(cfg.ConnMaxIdleTime * time.Second)

	return
}
