package model

import (
	"github.com/jinzhu/gorm"
)

// User based of gorm.Model
type User struct {
	gorm.Model
	Name string `json:"name"`
}

// Get user by ID
func (user *User) Get(tx *gorm.DB) *gorm.DB {
	return tx.Find(&user)
}

// FindByName get user by name
func (user *User) FindByName(tx *gorm.DB, name string) *gorm.DB {
	return tx.Where("name = ?", name).First(&user)
}

// Create new user
func (user *User) Create(tx *gorm.DB) *gorm.DB {
	return tx.Create(&user)
}

// Update user
func (user *User) Update(tx *gorm.DB) *gorm.DB {
	return tx.Save(&user)
}

// Delete user
func (user *User) Delete(tx *gorm.DB) *gorm.DB {
	return tx.Delete(&user)
}
