package db

import (
  "log"
  "yock.dev/coldstore/cuts"
  "gorm.io/gorm"
  "gorm.io/driver/sqlite"
)

func OpenDB() *gorm.DB {
  db, err := gorm.Open(sqlite.Open("coldstore.sqlite"))
  if err != nil {
    log.Fatal("Could not open database")
  }

  db.AutoMigrate(&cuts.Cut{})

  return db
}
