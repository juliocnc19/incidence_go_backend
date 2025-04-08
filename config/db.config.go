package config

import (
	"fmt"
	"incidence_grade/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(enviroments *Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		enviroments.DBHost, enviroments.DBPort, enviroments.DBUser, enviroments.DBPassword, enviroments.DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	migrate(db)
	seed(db)
	return db
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(
		&models.User{},
		&models.Incident{},
		&models.Notification{},
		&models.Role{},
		&models.Status{},
		&models.UserToken{},
		&models.Attachment{},
	)
}

func seed(db *gorm.DB) {
	roles := []models.Role{
		{Name: "admin"},
		{Name: "student"},
		{Name: "teacher"},
		{Name: "analyst"},
	}

	for _, role := range roles {
		var existingRole models.Role
		if err := db.Where("name = ?", role.Name).First(&existingRole).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				if err := db.Create(&role).Error; err != nil {
					fmt.Printf("Error al sembrar rol %s: %v\n", role.Name, err)
				}
			} else {
				fmt.Printf("Error al consultar rol %s: %v\n", role.Name, err)
			}
		}
	}

	statuses := []models.Status{
		{Name: "in_progress"},
		{Name: "on_hold"},
		{Name: "resolved"},
	}

	for _, status := range statuses {
		var existingStatus models.Status
		if err := db.Where("name = ?", status.Name).First(&existingStatus).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				if err := db.Create(&status).Error; err != nil {
					fmt.Printf("Error al sembrar estado %s: %v\n", status.Name, err)
				}
			} else {
				fmt.Printf("Error al consultar estado %s: %v\n", status.Name, err)
			}
		}
	}
}
