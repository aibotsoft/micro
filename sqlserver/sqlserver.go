package sqlserver

import (
	"database/sql"
	"github.com/aibotsoft/micro/config"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"

	"net"
	"net/url"
)

//sqlserver://sa:mypass@localhost:1234?database=master&connection+timeout=30

func BuildConnString(cfg *config.Config) string {
	query := url.Values{}
	query.Add("database", cfg.Mssql.Database)
	query.Add("app name", cfg.Service.Name)

	//in seconds (default is 15), set to 0 for no timeout
	query.Add("dial timeout", cfg.Mssql.DialTimeout)

	//in seconds; 0 to disable (default is 30) //my def 1440
	query.Add("keepAlive", cfg.Mssql.KeepAlive)
	//in bytes; 512 to 32767 (default is 4096)
	//Encrypted connections have a maximum packet size of 16383 bytes
	query.Add("packet size", cfg.Mssql.PacketSize)
	//logging flags (default 0/no logging, 63 for full logging)
	query.Add("log", cfg.Mssql.Log)

	u := &url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword(cfg.Mssql.User, cfg.Mssql.Password),
		Host:     net.JoinHostPort(cfg.Mssql.Host, cfg.Mssql.Port),
		RawQuery: query.Encode(),
	}
	return u.String()
}

func MustConnect(cfg *config.Config) *sql.DB {
	connString := BuildConnString(cfg)
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

func MustConnectX(cfg *config.Config) *sqlx.DB {
	connString := BuildConnString(cfg)
	db, err := sqlx.Open("sqlserver", connString)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	db.MapperFunc(func(s string) string { return s })
	return db
}
