package models

import (
	"fmt"
	"strings"
)

type PasswordPolicy struct {
	Id            uint  `gorm:"primaryKey"`
	MinLength     *int  `json:"minLength"`
	RequireNumber *bool `json:"requireNumber"`
	RequireLower  *bool `json:"requireLower"`
	RequireUpper  *bool `json:"requireUpper"`
	RequireSymbol *bool `json:"requireSymbol"`
	HistorySize   *int  `json:"historySize"`
}

func (policy *PasswordPolicy) Validate(password string, pwdHistory [][]byte) error {
	if len(password) < *policy.MinLength {
		return fmt.Errorf("password must be at least %d characters long", *policy.MinLength)
	}
	if *policy.RequireUpper && !strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		return fmt.Errorf("password must contain at least one uppercase letter")
	}
	if *policy.RequireLower && !strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz") {
		return fmt.Errorf("password must contain at least one lowercase letter")
	}
	if *policy.RequireNumber && !strings.ContainsAny(password, "0123456789") {
		return fmt.Errorf("password must contain at least one digit")
	}
	if *policy.RequireSymbol && !strings.ContainsAny(password, "!@#$%^&*()_-+={}[]\\|/?.>,<~`") {
		return fmt.Errorf("password must contain at least one symbol")
	}
	// TODO password policy history
	// for _, oldHash := range pwdHistory {
	// 	if err := bcrypt.CompareHashAndPassword(oldHash, []byte(password)); err == nil {
	// 		return fmt.Errorf("password has been used before and cannot be reused")
	// 	}
	// }
	return nil
}
