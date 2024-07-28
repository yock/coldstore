package cuts

import (
  "time"
  "gorm.io/gorm"
)

type Meat int64
type Unit int64

const (
  Beef Meat = 0
  Chicken Meat = 1
  Pork Meat = 2
)

const (
  Oz Unit = 0
  Gm Unit = 1
)

type Cut struct {
  gorm.Model
  Meat
  Weight int64
  Unit
  AddedAt time.Time
  ThawedAt time.Time
}

