package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID    int
	Name  string
	Email string
}

// create a user - CreateUser
func CreateUser(db *gorm.DB, User *User) (err error) {
	err = db.Create(User).Error
	if err != nil {
		return err
	}

	return nil
}

// get users - GetUsers
func GetUsers(db *gorm.DB, User *User, id int) (err error) {
	err = db.Find(User).Error
	if err != nil {
		return err
	}

	return nil
}

// get user by id - GetUser
func GetUser(db *gorm.DB, User *User, id int) (err error) {
	err = db.Where("id = ?", id).First(User).Error
	if err != nil {
		return err
	}

	return nil
}

// update user - UpdateUser
func UpdateUser(db *gorm.DB, User *User) (err error) {
	db.Save(User)
	return nil
}

// delte user - DeleteUser
func DeleteUser(db *gorm.DB, User *User, id int) (err error) {
	db.Where("id = ?", id).Delete(User)
	return nil
}
