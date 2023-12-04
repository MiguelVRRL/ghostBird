package models

import "gorm.io/gorm"

type Music struct {
  gorm.Model
  FileName string `gorm:"not null"`
  UUID     string `gorm:"unique;not null"`
}
