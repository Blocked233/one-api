package model

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/microsoft/go-mssqldb/azuread"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var sqlDB *sql.DB

type AzureSqlConfig struct {
	Server   string
	Port     int
	Database string
}

func initAzureSql(config AzureSqlConfig) (*gorm.DB, error) {

	var err error

	// Build connection string
	connString := fmt.Sprintf("server=%s;port=%d;database=%s;fedauth=ActiveDirectoryDefault;", config.Server, config.Port, config.Database)

	// Create SQL connection
	sqlDB, err = sql.Open(azuread.DriverName, connString)
	if err != nil {
		log.Fatal("Error connecting to database: ", err.Error())
	}

	//Use the SQL connection to initialize *gorm.DB
	db, err := gorm.Open(sqlserver.New(sqlserver.Config{Conn: sqlDB}), &gorm.Config{
		PrepareStmt: true, // precompile SQL
	})
	if err != nil {
		log.Fatal("GORM failed to connect database: ", err.Error())
	}
	return db, err
}
