package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Content    string
	QuestionID *uint
	Answers    []*Post `gorm:"foreignkey:QuestionID"`
}

func (p *Post) Answer(content string) error {
	return db.Model(p).Association("Answers").Append(&Post{Content: content})
}
