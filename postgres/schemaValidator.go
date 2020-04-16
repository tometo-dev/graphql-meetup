package postgres

import (
	"github.com/tsuki42/graphql-meetup/logging"
	"github.com/tsuki42/graphql-meetup/models"
)

func ValidateSchema() {
	if !Connection.Table("MEETUP").HasTable(&models.Meetup{}) {
		if err := Connection.Table("MEETUP").AutoMigrate(&models.Meetup{}).Error; err != nil {
			logging.ERROR.Println("Failed to create table MEETUP in the database")
			panic(err)
		} else {
			logging.INFO.Println("MEETUP table successfully created")
		}
		Connection.Table("MEETUP").AddForeignKey("user_id", "USER(id)", "SET NULL", "CASCADE")
	}

	if !Connection.Table("USER").HasTable(&models.User{}) {
		if err := Connection.Table("USER").AutoMigrate(&models.User{}).Error; err != nil {
			logging.ERROR.Println("Failed to create table USER in the database")
			panic(err)
		} else {
			logging.INFO.Println("USER table successfully created")
		}
	}
}
