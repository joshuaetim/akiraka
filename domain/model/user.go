package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password,omitempty"`
	Matric    string `json:"matric" gorm:"unique"`
	// Staff     []Staff   `json:"staff,omitempty"`
	// Visitors  []Visitor `json:"visitors,omitempty"`
}

func (u *User) PublicUser() *User {
	user := &User{
		Firstname: u.Firstname,
		Lastname:  u.Lastname,
		Email:     u.Email,
		Matric:    u.Matric,
	}
	user.ID = u.ID
	user.CreatedAt = u.CreatedAt
	return user
}

// table name
func (User) TableName() string {
	return "users"
}
