package manager

import (
	"api-payment/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type InfraManager interface {
	Conn() *sql.DB
}

type infraManager struct {
	db  *sql.DB
	cfg *config.Config
}

func (i *infraManager) initDb() error {
	var dbConf = i.cfg.DbConfig
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConf.Host,
		dbConf.Port,
		dbConf.User,
		dbConf.Password,
		dbConf.Name)
	db, err := sql.Open(dbConf.Driver, dataSourceName)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	i.db = db
	return nil
}

func (i *infraManager) Conn() *sql.DB {
	return i.db
}

func NewInfraManager(configParam *config.Config) (InfraManager, error) {
	conn := &infraManager{
		cfg: configParam,
	}
	err := conn.initDb()
	if err != nil {
		return nil, err
	}

	return conn, nil
}
