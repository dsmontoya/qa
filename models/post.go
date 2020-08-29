package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID         uint           `json:"id" gorm:"primarykey"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`
	Content    string         `json:"content"`
	QuestionID *uint          `json:"-"`
	Answers    []*Post        `json:"answers,omitempty" gorm:"foreignkey:QuestionID"`
}

func (p *Post) Answer(content string) (*Post, error) {
	answer := &Post{Content: content}
	if err := db.Model(p).Association("Answers").Append(answer); err != nil {
		return nil, err
	}
	return answer, nil
}

func GetQuestions() ([]*Post, error) {
	questions := []*Post{}
	d := db.Where("question_id is NULL").Find(&questions)
	if err := d.Error; err != nil {
		return nil, err
	}
	for _, question := range questions {
		if err := db.Model(question).Association("Answers").Find(&question.Answers); err != nil {
			return nil, err
		}
	}
	return questions, nil
}
