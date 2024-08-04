package data

import (
  "log"
  "gorm.io/gorm"
  "gorm.io/driver/sqlite"
)

var Conn *gorm.DB

func init() {
  connection, err := gorm.Open(sqlite.Open("coldstore.sqlite"))
  if err != nil {
    log.Panic("Could not open database")
  }

  Conn = connection

  Conn.AutoMigrate(&Cut{})
  log.Println("Data migrations complete")
}
