package db

import (
	"database/sql"
	"fmt"
	"time"

	"vayer-electric-backend/env"

	"github.com/DavidHuie/gomigrate"
	_ "github.com/lib/pq"
)

type DbSource struct {
	conn *sql.DB
}

func CreateDbSource(dsn string) (DbSource, error) {
	d, err := sql.Open("postgres", dsn)

	if err != nil {
		return DbSource{}, err
	}

	go func() {
		for {
			d.Ping()
			time.Sleep(time.Second * 30)
		}
	}()

	d.SetMaxOpenConns(6)
	d.SetMaxIdleConns(2)

	return DbSource{
		conn: d,
	}, nil
}

func GetDbSource() DbSource {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		env.DB_HOST,
		env.DB_PORT,
		env.DB_USER,
		env.DB_PASSWORD,
		env.DB_NAME,
	)
	src, err := CreateDbSource(dsn)

	if err != nil {
		panic(err)
	}

	return src
}

func (s DbSource) ValidateConnection() bool {
	return s.conn.Ping() == nil
}

func (s DbSource) Migrate(path string) error {
	migrator, _ := gomigrate.NewMigrator(s.conn, gomigrate.Mysql{}, path)
	defer s.conn.Close()
	return migrator.Migrate()
}