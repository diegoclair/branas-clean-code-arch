package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	"github.com/GuiaBolso/darwin"
	"github.com/diegoclair/branas-clean-code-arch/infra/config"
	"github.com/diegoclair/branas-clean-code-arch/infra/data/migrations"
	"github.com/diegoclair/go_utils-lib/v2/logger"
	mysqlDriver "github.com/go-sql-driver/mysql"
)

var (
	conn    *sql.DB
	onceDB  sync.Once
	connErr error
)

//Connect returns a connection of MySQL
func Connect() (conn *sql.DB, err error) {
	onceDB.Do(func() {
		cfg := config.GetConfigEnvironment()

		dataSourceName := fmt.Sprintf("%s:root@tcp(%s:%d)/%s?charset=utf8",
			cfg.DB.User, cfg.DB.Host, cfg.DB.Port, cfg.DB.Name,
		)

		log.Println("Connecting to database...")
		db, connErr := sql.Open("mysql", dataSourceName)
		if connErr != nil {
			return
		}

		log.Println("Database Ping...")
		connErr = db.Ping()
		if connErr != nil {
			return
		}

		log.Println("Creating database...")
		if _, connErr = db.Exec("CREATE DATABASE IF NOT EXISTS sampamodas_db;"); connErr != nil {
			logger.Error("Create Database error: ", connErr)
			return
		}

		if _, connErr = db.Exec(fmt.Sprintf("USE %v;", cfg)); connErr != nil {
			logger.Error("Default Database error: ", connErr)
			return
		}

		connErr = mysqlDriver.SetLogger(logger.GetLogger())
		if connErr != nil {
			return
		}
		logger.Info("Database successfully configured")

		logger.Info("Running the migrations")
		driver := darwin.NewGenericDriver(db, darwin.MySQLDialect{})

		d := darwin.New(driver, migrations.Migrations, nil)

		connErr = d.Migrate()
		if connErr != nil {
			logger.Error("Migrate Error: ", connErr)
			return
		}

		logger.Info("Migrations executed")

		conn = db
	})

	return conn, connErr
}
