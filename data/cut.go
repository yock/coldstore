package data

import (
  "time"
  "gorm.io/gorm"
)

type Cut struct {
  gorm.Model
  Name string
  Meat string
  Weight int64
  Unit string
  AddedAt time.Time
  ThawedAt time.Time
}


