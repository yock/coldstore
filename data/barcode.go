package data

import (
  "time"
  "gorm.io/gorm"
  "image"
  "github.com/boombuler/barcode/code128"
  "github.com/google/uuid"
)

type Barcode struct {
  UUID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
  CutID uint
  Cut Cut
  PrintedAt time.Time
}

func (b *Barcode) AsImage() image.Image {
  code, _ := code128.Encode(b.UUID.String())
  return code
}
