package config

import (
	"fmt"
	"incidence_grade/models"
	"incidence_grade/utils"

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
		{Name: "resolved"},
		{Name: "rejected"},
		{Name: "draf"},
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

	adminRole := models.Role{}
	if err := db.Where("name = ?", "admin").First(&adminRole).Error; err != nil {
		fmt.Printf("Error al buscar rol admin: %v\n", err)
		return
	}

	hashedPassword, err := utils.HashPassword("123456")
	if err != nil {
		fmt.Printf("Error al hashear password para admin: %v\n", err)
		return
	}

	adminUser := models.User{
		Username:  "admin",
		Email:     "admin@gmail.com",
		Password:  hashedPassword,
		FirstName: "Admin",
		LastName:  "User",
		Cedula:    "0000000000",
		RoleID:    adminRole.ID,
	}

	var existingAdminUser models.User
	if err := db.Where("username = ?", adminUser.Username).First(&existingAdminUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			if err := db.Create(&adminUser).Error; err != nil {
				fmt.Printf("Error al sembrar usuario admin: %v\n", err)
			}
		} else {
			fmt.Printf("Error al consultar usuario admin: %v\n", err)
		}
	}
}
