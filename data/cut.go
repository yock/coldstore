package data

import (
  "fmt"
  "time"
  "gorm.io/gorm"
  "image"
  "github.com/boombuler/barcode/code128"
)

type Cut struct {
  gorm.Model
  Name string
  Meat string
  Weight int64
  Unit string
  AddedAt time.Time
  ThawedAt time.Time
  PrintedAt time.Time
}

func (c *Cut) Barcode() image.Image {
  strid := fmt.Sprintf("%010d", c.ID)
  code, _ := code128.Encode(strid)
  return code
}

