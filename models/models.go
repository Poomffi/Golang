package models

import "gorm.io/gorm"

type Fact struct {
	gorm.Model
	Question string `json:"question" gorm:"text;not null;default:null`
	Answer   string `json:"answer" gorm:"text;not null;default:null`
	NickName string `json:"nickname" gorm:"text;not null;default:null`
}
