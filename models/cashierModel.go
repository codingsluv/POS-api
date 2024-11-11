package models

import "time"

type Cashier struct {
	ID        uint      `json:"id" gorm:"type:INT(10) UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY"`
	Name      string    `json:"name" gorm:"type:VARCHAR(255) NOT NULL"`
	PassCode  string    `json:"pass_code" gorm:"type:VARCHAR(255) NOT NULL"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
