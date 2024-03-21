package models

type Post struct {
	ID      uint
	Title   string `json:"title" gorm:"text;not null;default:null`
	Content string `json:"content" gorm:"text;not null;default:null`
}
