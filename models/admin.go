package models

import "gorm.io/gorm"



type Admin struct {
  gorm.Model

  IsAdmin bool
  IsStaff bool
}

