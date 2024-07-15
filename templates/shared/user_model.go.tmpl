package models

import (
	"gorm.io/gorm"
	"context"
)

type GoEasyCLITestUser struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateUser(ctx context.Context,db *gorm.DB, user *GoEasyCLITestUser) error {
	return db.WithContext(ctx).Create(user).Error
}

func GetAllUsers(ctx context.Context, db *gorm.DB) ([]GoEasyCLITestUser, error) {
	var users []GoEasyCLITestUser
	if err := db.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserByID(ctx context.Context, db *gorm.DB, id string) (GoEasyCLITestUser, error) {
	var user GoEasyCLITestUser
	if err := db.WithContext(ctx).First(&user, id).Error; err != nil {
		return user, err
	}
	return user, nil
}

func UpdateUser(ctx context.Context, db *gorm.DB, user *GoEasyCLITestUser) error {
	return db.WithContext(ctx).Save(user).Error
}

func DeleteUser(ctx context.Context, db *gorm.DB, id string) error {
	return db.WithContext(ctx).Delete(&GoEasyCLITestUser{}, id).Error
}
