package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type GormUser struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique;not null"`
	Age  int    `gorm:"not null"`
}

type AdvancedDatabase struct {
	Conn *gorm.DB
}

func ConnectAdvanced(psqlInfo string) (*AdvancedDatabase, error) {
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Get the underlying sql.DB object
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// Set connection pool settings
	sqlDB.SetMaxOpenConns(10)   // Maximum number of open connections
	sqlDB.SetMaxIdleConns(5)    // Maximum number of idle connections
	sqlDB.SetConnMaxLifetime(time.Hour) // Maximum lifetime of a connection

	return &AdvancedDatabase{Conn: db}, nil
}

func (db *AdvancedDatabase) Close() error {
	sqlDB, err := db.Conn.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (db *AdvancedDatabase) CreateTables() error {
	err := db.Conn.AutoMigrate(&GormUser{})
	if err != nil {
		return err
	}
	fmt.Println("Tables created")
	return nil
}

func (db *AdvancedDatabase) InsertGormUser(user GormUser) error {
	return db.Conn.Create(&user).Error
}

func (db *AdvancedDatabase) QueryGormUsers(ageFilter *int, sortBy string, page, pageSize int) ([]GormUser, error) {
	query := db.Conn

	if ageFilter != nil {
		query = query.Where("age = ?", *ageFilter)
	}

	if sortBy != "" {
		query = query.Order(sortBy)
	}

	var users []GormUser
	result := query.Limit(pageSize).Offset((page-1)*pageSize).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (db *AdvancedDatabase) UpdateGormUser(id uint, name string, age int) error {
	user := GormUser{ID: id, Name: name, Age: age}
	return db.Conn.Save(&user).Error
}

func (db *AdvancedDatabase) DeleteGormUser(id uint) error {
	return db.Conn.Delete(&GormUser{}, id).Error
}
