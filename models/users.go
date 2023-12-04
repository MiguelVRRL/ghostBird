package models

import (
	"gorm.io/gorm"
)

type Status int

const ( 
  Online Status = iota
  Invisible
  Idle 
  busy
) 

type User struct {
  gorm.Model
  
  Name string
  Bio string
  Status Status 
  Data DataUser `gorm:"foreignkey:Email"`
  Admin Admin `gorm:"foreignkey:IsAdmin"`
}

type DataUser struct {
  gorm.Model

  Email string
  Password string
}


