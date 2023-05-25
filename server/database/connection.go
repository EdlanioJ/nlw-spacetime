package database

import (
	"fmt"

	"github.com/EdlanioJ/nlwpacetime/server/configs"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func OpenConnection() (*sqlx.DB, error) {
	conf := configs.GetDB()
	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Pass, conf.Name)

	conn, err := sqlx.Open("postgres", sc)
	if err != nil {
		return nil, err
	}

	err = conn.Ping()
	return conn, err
}
