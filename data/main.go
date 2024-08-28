package data

import (
  "log"
  "fmt"
  "os"
  "gorm.io/gorm"
  "gorm.io/driver/postgres"
)

var Conn *gorm.DB

func Connect() {
  dsn := fmt.Sprintf("host=%s user=%s dbname=%s TimeZone=UTC", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"))
  log.Println("Connecting to database")
  connection, err := gorm.Open(postgres.Open(dsn))
  if err != nil {
    log.Panic("Could not open database")
  }

  Conn = connection

  Conn.AutoMigrate(&Cut{})
  log.Println("Data migrations complete")
}
