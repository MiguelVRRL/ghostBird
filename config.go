package main

import (
	"fmt"
	"log"
	"net/http"

	//"net/http"

	"github.com/MiguelVRRL/ghostbird/handlers/admin"
	"github.com/MiguelVRRL/ghostbird/handlers/users"
	"github.com/MiguelVRRL/ghostbird/models"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)



var DB *gorm.DB

func init() {
  viper.SetConfigFile(".env")
  viper.ReadInConfig()
}

func setupRouter() {
  // admin
  http.HandleFunc("/admin/pannel", admin.PannelHandler) 
  http.HandleFunc("/admin/manage_users", admin.ManageUsersHandler)
  http.HandleFunc("/admin/manage_music", admin.ManageMusicHandler) 

  // users
  http.HandleFunc("/home", users.HomeHandler)
  http.HandleFunc("/login", users.LoginHandler)
  http.HandleFunc("/logout", users.LogoutHandler)
  http.HandleFunc("/register", users.RegisterHandler)
  http.HandleFunc("/profile", users.ProfileHandler)
  http.HandleFunc("/music-player", users.MusicPlayerHandler)
}

func setupDBConfigs() ( host, user, password, dbname, port string ) {
  host = viper.GetString("DB_HOST")
  user = viper.GetString("DB_USER")
  password = viper.GetString("DB_PASSWORD")
  dbname =  viper.GetString("DB_NAME")
  port = viper.GetString("DB_PORT")

  return 
}

func setupDB() {
  host,user, password, dbname, port := setupDBConfigs()
  dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",  host,user, password, dbname, port )
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
   log.Fatal("Payment failed")
  }
  DB = db
  if err := DB.AutoMigrate(models.User{},models.DataUser{},models.Admin{},models.Music{}); err != nil {
    log.Fatal(err.Error())
  }

}
