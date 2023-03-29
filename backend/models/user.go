package models

import (
	"time"
)

type User struct {
	Id                uint      `json:"id"`
	Name              string    `json:"name"`
	Email             string    `json:"email" gorm:"unique"`
	Password          []byte    `json:"-"`
	AdminRole         uint      `json:"adminRole"`
	BusinessRole      uint      `json:"businessRole"`
	ResidentialRole   uint      `json:"residentialRole"`
	LastModified      time.Time `json:"lastModified"`
	LastLoginAttempt  time.Time `json:"lastLoginAttempt"`
	LoginAttemptCount uint      `json:"loginAttemptCount"`
	Blocked           uint      `json:"blocked"`
	PasswordPolicyId  uint      `json:"-" gorm:"default:1;foreignKey:id"`
	LoginPolicyId     uint      `json:"-" gorm:"default:1;foreignKey:id"`

	// TODO password policy history. voir si on peut utiliser marshal et stocker un string car 2d array pas supporté par DB
	// History              [][]byte  `json:"-"`
}
