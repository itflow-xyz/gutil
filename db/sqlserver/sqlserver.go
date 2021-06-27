package db

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

const (
	sqlserverDriver           = "sqlserver"
	sqlserverConnectionString = "server=%s;user id=%s;password=%s;port=%d;database=%s;"
)

var sqlserverDB *sql.DB

// ConnectSQLserver creates the connection to a Microsoft SQL Server
// via server address, user, password, database and port
func ConnectSQLServer(s, u, p, d string, port int) (*sql.DB, error) {

	connectionString := fmt.Sprintf(sqlserverConnectionString, s, u, p, port, d)

	var err error

	sqlserverDB, err = sql.Open(sqlserverDriver, connectionString)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	err = sqlserverDB.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return sqlserverDB, nil
}
