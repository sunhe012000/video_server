package dbops

import (
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

var (
  dbConn *sql.DB
  err error
)

func init() {
  dbConn, err = sql.Open("mysql", "root:123@/video_server")
  if err != nil {
    panic(err.Error())
  }
}
