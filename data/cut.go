package data

import (
  "log"
  "time"
  "gorm.io/gorm"
  "image/jpeg"
  "encoding/base64"
  "bytes"
  "github.com/boombuler/barcode/code128"
  "github.com/google/uuid"
)

type Cut struct {
  UUID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
  Name string
  Meat string
  Weight int64
  Unit string
  AddedAt time.Time
  ThawedAt time.Time
  PrintedAt time.Time
}

func (c *Cut) Barcode() string {
  code, _ := code128.Encode(c.UUID.String())
  buffer := new(bytes.Buffer)
  if err := jpeg.Encode(buffer, code, nil); err != nil {
    log.Println("Unable to encode image")
  }
  return base64.StdEncoding.EncodeToString(buffer.Bytes())
}

